package tcp

import (
	"bufio"
	"context"
	"io"
	"log"
	"net"
	"strconv"
	"studygolang/wangdis/db"
	"sync"
)

type Handler struct {
	// 记录活跃的客户端连接
	activeConn sync.Map

	// 数据库引擎，执行指令并返回结果
	db db.DB

	// 关闭状态标志位，处于关闭过程中时拒绝建立新连接和接收新请求
	closing bool
	//closing atomic.AtomicBool
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing {
		// 关闭过程中不再接收新连接
		_ = conn.Close()
	}

	// 初始化客户端状态
	client := &Client{
		conn: conn,
	}
	h.activeConn.Store(client, 1)

	reader := bufio.NewReader(conn)
	var fixedLen int64  // 将要读取的 BulkString的正文长度
	var err error
	var msg []byte
	for {
		// 读取下一行数据
		if fixedLen == 0 {
			// 正常模式下使用 \r\n 区分数据行
			msg, err = reader.ReadBytes('\n')
			// 判断读到的数据是否以 \r\n 结尾
			if len(msg) == 0 || msg[len(msg) - 2] != '\r' {
				errReply := "invalid multibulk length"
				_, _ = client.conn.Write([]byte(errReply))
			}
		} else {
			// 当读取到BulkString第二行时，根据给出的长度进行读取
			msg = make([]byte, fixedLen + 2)
			_, err = io.ReadFull(reader, msg)
			// 判断读到的数据是否以 \r\n 结尾
			if len(msg) == 0 || msg[len(msg) - 2] != '\r' || msg[len(msg) - 1] != '\n' {
				errReply := "invalid multibulk length"
				client.conn.Write([]byte(errReply))
			}

			// Bulk String读取完毕，重新使用正常模式
			fixedLen = 0
		}

		// 处理IO异常
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				log.Println("connection close")
			} else {
				log.Println(err)
			}

			_ = client.Close()
			h.activeConn.Delete(client)
			return
		}

		// 解析收到的数据
		if !client.sending {
			// sending == false 表明收到了一条新指令
			if msg[0] == '*' {
				// 读取第一行获取参数个数
				expectedLine, err := strconv.ParseUint(string(msg[1:len(msg)-2]), 10, 32)
				if err != nil {
					_, _ = client.conn.Write(UnknownErrReplyBytes)
					continue
				}

				// 初始化客户端状态
				client.waitingReply.Add(1)    // 有指令未处理完成，阻止服务器关闭
				client.sending = true   // 正在接收指令中
				// 初始化计数器和缓冲区
				client.expectedArgsCount = uint32(expectedLine)
				client.receivedCount = 0
				client.args = make([][]byte, expectedLine)
			} else {
				// TODO: text protocol
			}
		} else {
			// 接收指令的剩余部分 (非首行)
			line := msg[0 : len(msg)-2]    // 移除换行符 \r\n
			if line[0] == '$' {
				// BulkString的首行，读取String的长度
				fixedLen, err = strconv.ParseInt(string(line[1:]), 10, 64)
				if err != nil {
					errReply := "invalid multibulk length"
					_, _ = client.conn.Write([]byte(errReply))
				}
				if fixedLen <= 0 {
					errReply := "invalid multibulk length"
					_, _ = client.conn.Write([]byte(errReply))
				}
			} else {
				// 收到参数
				client.args[client.receivedCount] = line
				client.receivedCount++
			}

			// 一条命令发送完毕
			if client.receivedCount == client.expectedArgsCount {
				client.sending = false

				// 执行命令并响应
				result := h.db.Exec(client.args)
				if result != nil {
					_, _ = conn.Write(result.ToBytes())
				} else {
					_, _ = conn.Write(UnknownErrReplyBytes)
				}

				// 重置客户端状态，等待下一个指令
				client.expectedArgsCount = 0
				client.receivedCount = 0
				client.args = nil
				client.waitingReply.Done()
			}
		}
	}
}

func (h *Handler) Close() error {
	// 遍历所有活跃的连接，并逐个关闭
	h.activeConn.Range(func(conn, value interface{}) bool {
		client, ok := conn.(*Client)
		if ok {
			_ = client.conn.Close()
		}
		return true
	})
	return nil
}
package server

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"studygolang/wangdis/db"
	"studygolang/wangdis/interface/database"
	"studygolang/wangdis/lib/sync/atomic"
	"studygolang/wangdis/redis/connection"
	"studygolang/wangdis/redis/parser"
	"studygolang/wangdis/redis/protocol"
	"sync"
)

var (
	unknownErrReplyBytes = []byte("-ERR unknown\r\n")
)

type Handler struct {
	// 记录活跃的客户端连接
	activeConn sync.Map

	// 数据库引擎，执行指令并返回结果
	db database.DB

	// 关闭状态标志位，处于关闭过程中时拒绝建立新连接和接收新请求
	closing atomic.Boolean
}

func MakeHandler() *Handler {
	return &Handler{
		db: &db.TestDB{},
	}
}

func (h *Handler) closeClient(client *connection.Connection) {
	_ = client.Close()
	h.activeConn.Delete(client)
}

func (h *Handler) Handle(ctx context.Context, conn net.Conn) {
	if h.closing.Get() {
		// 关闭过程中不再接收新连接
		_ = conn.Close()
	}

	// 初始化客户端状态
	client := connection.NewConn(conn)
	h.activeConn.Store(client, 1)

	ch := parser.ParseStream(conn)

	for payload := range ch {
		if payload.Err != nil {
			if payload.Err == io.EOF || payload.Err == io.ErrUnexpectedEOF || strings.Contains(payload.Err.Error(), "use of closed network connection") {
				h.closeClient(client)
				log.Printf("connection closed: %s", client.RemoteAddr().String())
				return
			}

			errReply := protocol.MakeErrReply(payload.Err.Error())
			err := client.Write(errReply.ToBytes())
			if err != nil {
				h.closeClient(client)
				log.Printf("connection closed: %s", client.RemoteAddr().String())
				return
			}
			continue
		}

		if payload.Data == nil {
			log.Println("empty payload")
			continue
		}

		r, ok := payload.Data.(*protocol.MultiBulkReply)
		if !ok {
			log.Println("require multi bulk protocol")
			continue
		}

		var cmdLine string
		for i := range r.Args {
			cmdLine = fmt.Sprintf("%s%s ", cmdLine, string(r.Args[i]))
		}
		log.Printf("got cmd from client: %s\n", cmdLine)

		// 执行命令并响应
		result := h.db.Exec(client, r.Args)
		if result != nil {
			_ = client.Write(result.ToBytes())
		} else {
			_ = client.Write(unknownErrReplyBytes)
		}
	}
}

func (h *Handler) Close() error {
	h.closing.Set(true)

	// 遍历所有活跃的连接，并逐个关闭
	h.activeConn.Range(func(conn, value interface{}) bool {
		client, ok := conn.(*connection.Connection)
		if ok {
			_ = client.Close()
		}
		return true
	})

	h.db.Close()
	return nil
}
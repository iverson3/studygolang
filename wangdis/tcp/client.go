package tcp

import (
	"net"
	"studygolang/wangdis/wait"
)

var (
	UnknownErrReplyBytes = []byte("UnknownErr ReplyBytes")
)

type Client struct {
	// 与客户端的Tcp连接
	conn net.Conn

	// 带有timeout功能的WaitGroup，用于优雅的关闭连接
	// 当响应被完整发送前保持waiting状态，阻止连接被关闭
	waitingReply wait.Wait
	// 标记客户端是否正在发送指令
	sending bool
	//sending atomic.AtomicBool

	// 客户端正在发送的参数数量，即Array第一行指定的数组长度
	expectedArgsCount uint32
	// 已经接收的参数数量，即 len(args)
	receivedCount uint32

	// 已经接收到的命令参数，每个参数由一个[]byte表示
	args [][]byte
}

func (c *Client) Close() error {
	return nil
}


package model

import (
	"net"
	"studygolang/chatroom/server/model"
)

type CurUser struct {
	Conn net.Conn
	model.User
}

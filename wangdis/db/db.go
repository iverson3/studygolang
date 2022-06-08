package db

import (
	"studygolang/wangdis/interface/redis"
	"studygolang/wangdis/redis/protocol"
)

type TestDB struct {

}

func (db *TestDB) Exec(c redis.Connection, args [][]byte) redis.Reply {
	return &protocol.StatusReply{Status: "OK"}
}
func (db *TestDB) Close() {

}
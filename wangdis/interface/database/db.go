package database

import "studygolang/wangdis/interface/redis"

type DB interface {
	Exec(client redis.Connection, cmdLine [][]byte) redis.Reply
	Close()
}

// DataEntity stores data bound to a key, including a string, list, hash, set and so on
type DataEntity struct {
	Data interface{}
}

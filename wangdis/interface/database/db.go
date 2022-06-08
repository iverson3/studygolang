package database

import "studygolang/wangdis/interface/redis"

type DB interface {
	Exec(client redis.Connection, cmdLine [][]byte) redis.Reply
	Close()
}

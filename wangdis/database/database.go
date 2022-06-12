package database

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"studygolang/wangdis/aof"
	"studygolang/wangdis/config"
	"studygolang/wangdis/interface/redis"
	"studygolang/wangdis/redis/protocol"
)

type MultiDB struct {
	dbSet []*DB

	aofHandler *aof.Handler
}

func NewStandaloneServer() *MultiDB {
	mdb := &MultiDB{}
	if config.Properties.Databases == 0 {
		config.Properties.Databases = 16
	}

	mdb.dbSet = make([]*DB, config.Properties.Databases)
	for i := range mdb.dbSet {
		singleDB := makeDB()
		singleDB.index = i
		mdb.dbSet[i] = singleDB
	}
	return mdb
}

func (mdb *MultiDB) Exec(c redis.Connection, cmdLine [][]byte) (result redis.Reply) {
	defer func() {
		if err := recover(); err != nil {
			log.Println(fmt.Sprintf("error occurs: %v\n%s", err, string(debug.Stack())))
			result = &protocol.UnknownErrReply{}
		}
	}()

	cmdName := strings.ToLower(string(cmdLine[0]))
	// 对于cmdName是特殊命令时的判断和处理
	_ = cmdName
	// todo: xxx

	dbIndex := c.GetDBIndex()
	if dbIndex >= len(mdb.dbSet) {
		return protocol.MakeErrReply("ERR DB index is out of range")
	}
	selectdDB := mdb.dbSet[dbIndex]
	return selectdDB.Exec(c, cmdLine)
}

// AfterClientClose does some clean after client close connection
func (mdb *MultiDB) AfterClientClose(c redis.Connection) {
	// todo: xxx
}

func (mdb *MultiDB) Close() {
	if mdb.aofHandler != nil {
		mdb.aofHandler.Close()
	}
}

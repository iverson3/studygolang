package database

import (
	"fmt"
	"log"
	"runtime/debug"
	"strings"
	"studygolang/wangdis/aof"
	"studygolang/wangdis/config"
	"studygolang/wangdis/interface/database"
	"studygolang/wangdis/interface/redis"
	"studygolang/wangdis/lib/utils"
	"studygolang/wangdis/redis/protocol"
	"time"
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

	var validAof bool
	if config.Properties.AppendOnly {
		aofHandler, err := aof.NewAOFHandler(mdb, func() database.EmbedDB {
			return MakeBasicMultiDB()
		})
		if err != nil {
			panic(err)
		}

		mdb.aofHandler = aofHandler
		for _, db := range mdb.dbSet {
			singleDB := db
			singleDB.addAof = func(cmdLine CmdLine) {
				mdb.aofHandler.AddAof(singleDB.index, cmdLine)
			}
		}
		validAof = true
	}

	if config.Properties.RDBFilename != "" && !validAof {
		// todo: load rdb
		//loadRdb(mdb)
	}
	return mdb
}

func MakeBasicMultiDB() *MultiDB {
	mdb := &MultiDB{}
	mdb.dbSet = make([]*DB, config.Properties.Databases)
	for i := range mdb.dbSet {
		mdb.dbSet[i] = makeBasicDB()
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
	// 1.检验权限
	// 2.特殊命令的单独处理 (不能在事务中执行的特殊命令，subscribe publish flushall等)
	if cmdName == "bgrewriteaof" {
		return BGRewriteAOF(mdb, cmdLine[1:])
	} else if cmdName == "rewriteaof" {
		return RewriteAOF(mdb, cmdLine[1:])
	} else if cmdName == "flushall" {
		return mdb.flushAll()
	}

	// todo: support multi database transaction

	// 之后则是普通命令的执行
	dbIndex := c.GetDBIndex()
	if dbIndex >= len(mdb.dbSet) {
		return protocol.MakeErrReply("ERR DB index is out of range")
	}
	selectedDB := mdb.dbSet[dbIndex]
	return selectedDB.Exec(c, cmdLine)
}

// AfterClientClose does some clean after client close connection
func (mdb *MultiDB) AfterClientClose(c redis.Connection) {
	// todo: AfterClientClose
}

func (mdb *MultiDB) Close() {
	if mdb.aofHandler != nil {
		mdb.aofHandler.Close()
	}
}

func (mdb *MultiDB) flushAll() redis.Reply {
	for _, db := range mdb.dbSet {
		db.Flush()
	}

	if mdb.aofHandler != nil {
		mdb.aofHandler.AddAof(0, utils.ToCmdLine("FlushAll"))
	}
	return &protocol.OkReply{}
}

func (mdb *MultiDB) ExecWithLock(conn redis.Connection, cmdLine [][]byte) redis.Reply {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) ExecMulti(conn redis.Connection, watching map[string]uint32, cmdLines []database.CmdLine) redis.Reply {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) GetUndoLogs(dbIndex int, cmdLine [][]byte) []database.CmdLine {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) ForEach(dbIndex int, cb func(key string, data *database.DataEntity, expiration *time.Time) bool) {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) RWLocks(dbIndex int, writeKeys []string, readKeys []string) {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) RWUnLocks(dbIndex int, writeKeys []string, readKeys []string) {
	//TODO implement me
	panic("implement me")
}

func (mdb *MultiDB) GetDBSize(dbIndex int) (int, int) {
	//TODO implement me
	panic("implement me")
}

// BGRewriteAOF 在后台异步的执行aof重写
func BGRewriteAOF(db *MultiDB, args [][]byte) redis.Reply {
	go db.aofHandler.Rewrite()
	return protocol.MakeStatusReply("Background append only file rewriting started")
}

// RewriteAOF 同步的执行aof重写
func RewriteAOF(db *MultiDB, args [][]byte) redis.Reply {
	err := db.aofHandler.Rewrite()
	if err != nil {
		return protocol.MakeErrReply(err.Error())
	}
	return protocol.MakeOkReply()
}

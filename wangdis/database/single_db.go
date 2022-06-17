package database

import (
	"strings"
	"studygolang/wangdis/datastruct/dict"
	"studygolang/wangdis/datastruct/lock"
	"studygolang/wangdis/interface/database"
	"studygolang/wangdis/interface/redis"
	"studygolang/wangdis/redis/protocol"
	"sync"
	"time"
)

const (
	dataDictSize = 1 << 16
	ttlDictSize  = 1 << 10
	lockerSize   = 1024
)

type DB struct {
	index int

	// key -> DataEntity
	data dict.Dict
	// key -> expireTime (time.Time)
	ttlMap dict.Dict
	// key -> version(uint32)
	versionMap dict.Dict

	// dict.Dict will ensure concurrent-safety of its method
	// use this mutex for complicated command only, eg. rpush, incr ...
	// 保证redis某些命令执行时的并发安全
	locker *lock.Locks
	// stop all data access for execFlushDB
	stopWorld sync.WaitGroup
	addAof    func(CmdLine)
}

// ExecFunc is interface for command executor
// args don't include cmd line
type ExecFunc func(db *DB, args [][]byte) redis.Reply

// PreFunc analyses command line when queued command to `multi`
// returns related write keys and read keys
type PreFunc func(args [][]byte) ([]string, []string)

// CmdLine is alias for [][]byte, represents a command line
type CmdLine = [][]byte

// UndoFunc returns undo logs for the given command line
// execute from head to tail when undo
type UndoFunc func(db *DB, args [][]byte) []CmdLine

func makeDB() *DB {
	db := &DB{
		data:       dict.MakeConcurrent(dataDictSize),
		ttlMap:     dict.MakeConcurrent(ttlDictSize),
		versionMap: dict.MakeConcurrent(dataDictSize),
		locker:     lock.Make(lockerSize),
		addAof:     func(line CmdLine) {},
	}
	return db
}

func makeBasicDB() *DB {
	return &DB{
		data:       dict.MakeSimple(),
		ttlMap:     dict.MakeSimple(),
		versionMap: dict.MakeSimple(),
		locker:     lock.Make(1),
		addAof:     func(line CmdLine) {},
	}
}

func (db *DB) Exec(c redis.Connection, cmdLine [][]byte) redis.Reply {
	cmdName := strings.ToLower(string(cmdLine[0]))
	// 对于cmdName的一些判断和检验
	_ = cmdName
	// todo: xxx

	return db.execNormalCommand(cmdLine)
}

func (db *DB) execNormalCommand(cmdLine [][]byte) redis.Reply {
	cmdName := strings.ToLower(string(cmdLine[0]))
	cmd, ok := cmdTable[cmdName]
	if !ok {
		return protocol.MakeErrReply("ERR unknown command '" + cmdName + "'")
	}
	if !validateArity(cmd.arity, cmdLine) {
		return protocol.MakeArgNumErrReply(cmdName)
	}

	prepare := cmd.prepare
	write, read := prepare(cmdLine[1:])
	db.addVersion(write...)
	db.RWLocks(write, read)
	defer db.RWUnLocks(write, read)
	executorFunc := cmd.executor
	return executorFunc(db, cmdLine[1:])
}

func validateArity(arity int, cmdArgs [][]byte) bool {
	argNum := len(cmdArgs)
	if arity >= 0 {
		return argNum == arity
	}
	return argNum >= -arity
}

func (db *DB) addVersion(keys ...string) {
	for _, key := range keys {
		versionCode := db.GetVersion(key)
		db.versionMap.Put(key, versionCode+1)
	}
}

func (db *DB) GetVersion(key string) uint32 {
	entity, ok := db.versionMap.Get(key)
	if !ok {
		return 0
	}
	return entity.(uint32)
}

func (db *DB) RWLocks(writeKeys []string, readKeys []string) {
	db.locker.RWLocks(writeKeys, readKeys)
}

func (db *DB) RWUnLocks(writeKeys []string, readKeys []string) {
	db.locker.RWUnLocks(writeKeys, readKeys)
}

// 数据访问

func (db *DB) PutEntity(key string, entity *database.DataEntity) int {
	db.stopWorld.Wait()
	return db.data.Put(key, entity)
}

func (db *DB) PutIfExists(key string, entity database.DataEntity) int {
	db.stopWorld.Wait()
	return db.data.PutIfExists(key, entity)
}

func (db *DB) PutIfAbsent(key string, entity database.DataEntity) int {
	db.stopWorld.Wait()
	return db.data.PutIfAbsent(key, entity)
}

// GetEntity returns DataEntity bind to given key
func (db *DB) GetEntity(key string) (*database.DataEntity, bool) {
	db.stopWorld.Wait()

	raw, ok := db.data.Get(key)
	if !ok {
		return nil, false
	}

	if db.IsExpired(key) {
		return nil, false
	}

	entity, _ := raw.(*database.DataEntity)
	return entity, true
}

// Remove the given key from db
func (db *DB) Remove(key string) {
	db.stopWorld.Wait()
	db.data.Remove(key)
	db.ttlMap.Remove(key)
	//taskKey := genExpireTask(key)
	//timewheel.Cancel(taskKey)
}

func (db *DB) Removes(keys ...string) (deleted int) {
	db.stopWorld.Wait()
	deleted = 0
	for _, key := range keys {
		_, exists := db.data.Get(key)
		if exists {
			db.Remove(key)
			deleted++
		}
	}
	return
}

func (db *DB) Flush() {
	db.stopWorld.Add(1)
	defer db.stopWorld.Done()

	db.data.Clear()
	db.ttlMap.Clear()
	db.locker = lock.Make(lockerSize)
}

func (db *DB) IsExpired(key string) bool {
	rawExpireTime, ok := db.ttlMap.Get(key)
	if !ok {
		return false
	}

	expireTime, _ := rawExpireTime.(time.Time)
	expired := time.Now().After(expireTime)
	if expired {
		db.Remove(key)
	}
	return expired
}

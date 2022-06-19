package database

import (
	"studygolang/wangdis/aof"
	"studygolang/wangdis/lib/utils"
)

func readFirstKey(args [][]byte) ([]string, []string) {
	if args == nil || len(args) == 0 {
		return nil, nil
	}

	key := string(args[0])
	return nil, []string{key}
}

func writeFirstKey(args [][]byte) ([]string, []string) {
	if args == nil || len(args) == 0 {
		return nil, nil
	}

	key := string(args[0])
	return []string{key}, nil
}

func rollbackFirstKey(db *DB, args [][]byte) []CmdLine {
	key := string(args[0])
	return rollbackGivenKeys(db, key)
}

func rollbackGivenKeys(db *DB, keys ...string) []CmdLine {
	var undoCmdLines [][][]byte
	for _, key := range keys {
		entity, ok := db.GetEntity(key)
		if !ok {
			undoCmdLines = append(undoCmdLines, utils.ToCmdLine("DEL", key))
		} else {
			undoCmdLines = append(undoCmdLines,
				utils.ToCmdLine("DEL", key), // clean existed first
				aof.EntityToCmd(key, entity).Args,
				//toTTLCmd(db, key).Args,
			)
		}
	}
	return undoCmdLines
}

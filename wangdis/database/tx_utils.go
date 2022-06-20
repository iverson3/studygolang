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

func rollbackSetMembers(db *DB, key string, members ...string) []CmdLine {
	var undoCmdLines [][][]byte
	set, errReply := db.getAsSet(key)
	if errReply != nil {
		return nil
	}

	if set == nil {
		undoCmdLines = append(undoCmdLines, utils.ToCmdLine("DEL", key))
		return undoCmdLines
	}

	for _, member := range members {
		has := set.Has(member)
		if has {
			undoCmdLines = append(undoCmdLines, utils.ToCmdLine("SADD", key, member))
		} else {
			undoCmdLines = append(undoCmdLines, utils.ToCmdLine("SREM", key, member))
		}
	}
	return undoCmdLines
}

// rollback for command: SADD or SREM
func undoSetChange(db *DB, args [][]byte) []CmdLine {
	key := string(args[0])
	memberArgs := args[1:]
	members := make([]string, 0, len(memberArgs))
	for _, arg := range memberArgs {
		members = append(members, string(arg))
	}
	return rollbackSetMembers(db, key, members...)
}
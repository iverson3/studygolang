package dict

import (
	"math"
	"sync"
)

type ConcurrentDict struct {
	table []*shard
	count int32     // key的总数
	shardCount int  // shard的数量
}

type shard struct {
	m map[string]interface{}
	mu sync.RWMutex
}

func computeCapacity(param int) int {
	if param <= 16 {
		return 16
	}
	n := param - 1
	n |= n >> 1
	n |= n >> 2
	n |= n >> 4
	n |= n >> 8
	n |= n >> 16
	if n < 0 {
		return math.MaxInt32
	}
	return n + 1
}

func MakeConcurrent(shardCount int) *ConcurrentDict {
	shardCount = computeCapacity(shardCount)
	table := make([]*shard, shardCount)
	for i := 0; i < shardCount; i++ {
		table[i] = &shard{
			m: make(map[string]interface{}),
		}
	}
	return &ConcurrentDict{
		table:      table,
		count:      0,
		shardCount: shardCount,
	}
}
package dict

import "sync"

type ConcurrentDict struct {
	table []*shard
	count int32     // key的总数
	shardCount int  // shard的数量
}

type shard struct {
	m map[string]interface{}
	mu sync.RWMutex
}


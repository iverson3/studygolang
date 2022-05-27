package wait

import "sync"

type Wait struct {
	sync.WaitGroup
}
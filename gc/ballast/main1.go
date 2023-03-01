package main

import (
	"flag"
	"fmt"
	"sync"
	"time"
)

type node struct {
	next *next
}
type next struct {
	left, right node
}
func create(d int) node {
	if d == 1 {
		return node{&next{node{}, node{}}}
	}
	return node{&next{create(d - 1), create(d - 1)}}
}
func (p node) check() int {
	sum := 1
	current := p.next
	for current != nil {
		sum += current.right.check() + 1
		current = current.left.next
	}
	return sum
}

var (
	ballastObj []byte
	depth = flag.Int("depth", 5, "depth")
	ballast = flag.Bool("ballast", false, "run program with Ballast")
)

func main() {
	flag.Parse()
	withBallast := *ballast

	fmt.Println(withBallast)
	if withBallast {
		ballastObj = make([]byte, 2*1024*1024*1024)
		//runtime.KeepAlive(ballastObj)
	}

	start := time.Now()
	const MinDepth = 4
	const NoTasks = 4
	maxDepth := *depth
	longLivedTree := create(maxDepth)
	stretchTreeCheck := ""
	wg := new(sync.WaitGroup)
	wg.Add(1)
	go func() {
		stretchDepth := maxDepth + 1
		stretchTreeCheck = fmt.Sprintf("stretch tree of depth %d\t check: %d",
			stretchDepth, create(stretchDepth).check())
		wg.Done()
	}()
	results := make([]string, (maxDepth-MinDepth)/2+1)
	for i := range results {
		depth := 2*i + MinDepth
		n := (1 << (maxDepth - depth + MinDepth)) / NoTasks
		tasks := make([]int, NoTasks)
		wg.Add(NoTasks)
		// 执行NoTasks个goroutine, 每个goroutine执行n个深度为depth的tree的check
		// 一共是n*NoTasks个tree,每个tree的深度是depth
		for t := range tasks {
			go func(t int) {
				check := 0
				for i := n; i > 0; i-- {
					check += create(depth).check()
				}
				tasks[t] = check
				wg.Done()
			}(t)
		}
		wg.Wait()
		check := 0 // 总检查次数
		for _, v := range tasks {
			check += v
		}
		results[i] = fmt.Sprintf("%d\t trees of depth %d\t check: %d",
			n*NoTasks, depth, check)
	}
	fmt.Println(stretchTreeCheck)
	for _, s := range results {
		fmt.Println(s)
	}
	fmt.Printf("long lived tree of depth %d\t check: %d\n",
		maxDepth, longLivedTree.check())
	fmt.Printf("took %.02f s\n", float64(time.Since(start).Milliseconds())/1000)
}
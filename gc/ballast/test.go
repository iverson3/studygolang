package main
//
//import (
//	"fmt"
//	"math/rand"
//	"sync"
//	"time"
//)
//
//var result [][]int
//
//func main() {
//	//var upCounter int
//	// 输入
//	var summonRate float64 = 1
//
//	result = make([][]int, 7)
//
//	fmt.Println("Official Rate (%): ")
//	fmt.Print("Summon times        :  ")
//	start := time.Now()
//
//	for i := 50; i < 600; i += 50 {
//		fmt.Printf("  %d   |", i)
//	}
//	fmt.Println()
//
//	var wg sync.WaitGroup
//	for k := 0; k < 7; k++ {
//		wg.Add(1)
//		go task(&wg, &result, summonRate, k)
//	}
//	wg.Wait()
//
//	for i, list := range result {
//		fmt.Printf("more than %d star(s) : ", i)
//		for _, item := range list {
//			fmt.Printf(" %.02f%s |", float64(item)/1000, "%")
//		}
//		fmt.Println()
//	}
//
//	fmt.Printf("took: %.02f s\n", time.Since(start).Seconds())
//}
//
//func task(wg *sync.WaitGroup, result *[][]int, summonRate float64, k int) {
//	var upCounter int
//	for j := 50; j < 600; j += 50 {
//		upCounter = 0
//		for i :=0; i < 100000; i++ {
//			if upSummon(summonRate, j) > k {
//				upCounter++
//			}
//		}
//		(*result)[k] = append((*result)[k], upCounter)
//	}
//	wg.Done()
//}
//
//func upSummon(summonRate float64, summonTimes int) int {
//	var upCounter int
//	var lastSummon int
//	rand.Seed(time.Now().UnixNano())
//
//	for i := 0; i < summonTimes; i++ {
//		lastSummon++
//		if lastSummon == 100 && rand.Float64() * 15 < summonRate {
//			upCounter++
//			lastSummon = 0
//		} else if rand.Float64() * 1000 < summonRate * 10 {
//			upCounter++
//		}
//	}
//	return upCounter
//}

package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"syscall"
	"time"
	"net/http"
	_ "net/http/pprof"
	file_box "git.myscrm.cn/component-center/cc-file-box-sdk"
)

// 切片 - 预分配内存 vs append

type Test struct {
	A string
	B string
	C string
	D string
	E string
	F string
	G string
	H string
	I string
	J string
	K int64
	L int64
	M string
}

var client *file_box.Client

func init() {
	client = file_box.NewClient(file_box.Config{Scene: "forsdk"})
}

func GetObject() {
	ctx := context.Background()
	path := "cc-file-box/forsdk/2022-06-01 17:42:59-export_template.xlsx" // 上传文件返回的path
	// 默认过期时间为30min
	url, err := client.GetObjectUrl(ctx, path)
	// url = http://bucketName.oss-cn-shenzhen.aliyuncs.com/appName/scene/filename?sign=xxx

	// 指定过期时间1小时
	url, err = client.GetObjectUrl(ctx, path, file_box.WithExpiredInSec(60*60))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(url)
	file := os.File{}
	file.Sync()
}

func main() {
	//GetObject()

	f1, _ := os.Open("test_main.go")
	fd1 := f1.Fd()
	f1.Seek(64, 0)

	f2, _ := os.Open("test_main.go")
	fd2 := f2.Fd()


	fmt.Println(fd1)
	fmt.Println(fd2)


	file := "test_main.go"
	fileinfo, _ := os.Stat(file)
	fmt.Printf("fileinfo.Sys() = %#v\n", fileinfo.Sys())
	fmt.Printf("fileinfo = %#v\n", fileinfo)
	stat, ok := fileinfo.Sys().(*syscall.Stat_t)
	if !ok {
		fmt.Printf("Not a syscall.Stat_t")
		return
	}
	fmt.Printf("stat = %#v\n", stat)
	fmt.Printf("stat.Ino = %#v\n", stat.Ino)

	return

	go func() {
		log.Println(http.ListenAndServe(":8080", nil))
	}()

	list := createSlice(100)

	repeat := 1000

	for {
		//DoPreAllocMem(repeat, list)
		DoWithAppend(repeat, list)
	}
}

func DoPreAllocMem(repeat int, list []Test) {
	// 预分配内存
	start := time.Now()
	for i := 0; i < repeat; i++ {
		_ = PreAllocMem(list)
	}
	duration := time.Since(start).String()
	fmt.Printf("prealloc memory duration: %s \n", duration)
}
func DoWithAppend(repeat int, list []Test) {
	// with append
	start2 := time.Now()
	for i := 0; i < repeat; i++ {
		_ = WithAppend(list)
	}
	duration2 := time.Since(start2).String()
	fmt.Printf("with append duration: %s \n", duration2)
}

// PreAllocMem 预分配内存
func PreAllocMem(list []Test) []*Test {
	tests := make([]*Test, len(list))
	for i := range list {
		tests[i] = &Test{
			A: list[i].A,
			B: list[i].B,
			C: list[i].C,
			D: list[i].D,
			E: list[i].E,
			F: list[i].F,
			G: list[i].G,
			H: list[i].H,
			I: list[i].I,
			J: list[i].J,
			K: list[i].K,
			L: list[i].L,
			M: list[i].M,
		}
	}
	return tests
}

// WithAppend 初始切片长度为0，然后append
func WithAppend(list []Test) []*Test {
	tests := make([]*Test, 0)
	for i := range list {
		tests = append(tests, &Test{
			A: list[i].A,
			B: list[i].B,
			C: list[i].C,
			D: list[i].D,
			E: list[i].E,
			F: list[i].F,
			G: list[i].G,
			H: list[i].H,
			I: list[i].I,
			J: list[i].J,
			K: list[i].K,
			L: list[i].L,
			M: list[i].M,
		})
	}
	return tests
}

func createSlice(lens int) []Test {
	list := make([]Test, lens)
	for i := 0; i < lens; i++ {
		list[i] = Test{
			A: "a",
			B: "b",
			C: "c",
			D: "d",
			E: "e",
			F: "f",
			G: "g",
			H: "h",
			I: "i",
			J: "j",
			K: 1,
			L: 2,
			M: "",
		}
	}
	return list
}

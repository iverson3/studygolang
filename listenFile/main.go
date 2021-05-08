package main

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"studygolang/listenFile/model"
	"studygolang/listenFile/process"
	"time"
)

// goroutine监听文件变化
// 可动态管理goroutine
// 1. 新建一个goroutine对指定的文件进行监听
// 2. 停止监听文件并退出goroutine
func main() {
	go WaitExitSignal()

	process.Init()
	process.Start()


	//WatchDirTest()
}

func WaitExitSignal() {
	ch := make(chan os.Signal)

	signal.Notify(ch, os.Interrupt, os.Kill)
	<-ch

	fmt.Println("got exit signal, and main process will exit")

	err := model.TaskMgr.ExitAllTaskProcess()
	if err != nil {
		fmt.Printf("exit other task process failed! error: %v", err)
	}

	// 主进程睡眠，等待其他工作协程退出
	time.Sleep(1 * time.Second)
	fmt.Println("main process exit!")
	os.Exit(1)
}

func WatchDirTest() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("New watcher failed! error: ", err)
		return
	}
	defer watcher.Close()

	rootPath := "./watchDir"
	done := make(chan bool)
	err = WalkAllDirs(done, watcher, rootPath)
	if err != nil {
		log.Fatal("Walk Dirs failed! error: ", err)
	}

	<-done
	fmt.Println("bye~")
}

func WalkAllDirs(done chan bool,watcher *fsnotify.Watcher, rootPath string) (err error) {
	// 通过Walk来遍历目录下的所有子目录
	err = filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		// 这里判断是否为目录，只需监控目录即可
		// 目录下的文件也在监控范围内，不需要我们一个一个加

		if info.IsDir() {
			path, err = filepath.Abs(path)
			if err != nil {
				return err
			}
			err = watcher.Add(path)
			if err != nil {
				return err
			}
			log.Println("添加监控目录：", path)
		}
		return nil
	})
	if err != nil {
		return 
	}

	go WatchProcess(done, watcher)
	return 
}

func WatchProcess(done chan bool, watcher *fsnotify.Watcher) {
	defer close(done)

	eventChan := make(chan *fsnotify.Event)
	go DealEvent(watcher, eventChan)

	log.Println("start to watch...")
	for {
		select {
		case event, ok := <-watcher.Events:
			if !ok {
				break
			}
			eventChan<- &event
		case err, ok := <-watcher.Errors:
			if !ok {
				break
			}
			log.Println("watcher got error: ", err)
		}
	}
}

func DealEvent(watcher *fsnotify.Watcher, eventChan chan *fsnotify.Event) {
	for event := range eventChan {
		switch event.Op {
		case fsnotify.Create:
			log.Println("create event")
			go DealCreateEvent(watcher, event)
		case fsnotify.Remove:
			log.Println("remove event")
			go DealRemoveEvent(watcher, event)
		case fsnotify.Write:
			log.Println("write event")
			go DealWriteEvent(event)
		case fsnotify.Rename:
			log.Println("rename event")
			go DealRenameEvent(watcher, event)
		case fsnotify.Chmod:
			log.Println("chmod event")
		default:
			log.Println("other event of watcher")
			fmt.Printf("event: %v\n", event)
			fmt.Println(event.Op)
			fmt.Println(event.Name)
		}
	}
}

func DealCreateEvent(watcher *fsnotify.Watcher, event *fsnotify.Event)  {
	log.Println("====")
	log.Println(event)

	info, err := os.Stat(event.Name)
	if err != nil {
		log.Printf("get path[%s] info failed! error: %v", event.Name, err)
		return
	}
	if info.IsDir() {
		err = watcher.Add(event.Name)
		if err != nil {
			log.Printf("add path[%s] to watch failed! error: %v", event.Name, err)
			return
		}
		log.Println("添加监控目录：", event.Name)
	}
}
func DealRemoveEvent(watcher *fsnotify.Watcher, event *fsnotify.Event)  {
	info, err := os.Stat(event.Name)
	if err != nil {
		log.Printf("get path[%s] info failed! error: %v", event.Name, err)
		return
	}
	if info.IsDir() {
		err = watcher.Remove(event.Name)
		if err != nil {
			log.Printf("remove path[%s] from watch failed! error: %v", event.Name, err)
			return
		}
		log.Println("移除监控目录：", event.Name)
	}
}
func DealWriteEvent(event *fsnotify.Event)  {

}
func DealRenameEvent(watcher *fsnotify.Watcher, event *fsnotify.Event)  {
	// 如果重命名文件是目录，则移除监控
	// 注意这里无法使用os.Stat来判断是否是目录了
	// 因为重命名后，go已经无法找到原文件来获取信息了
	// 所以这里就简单粗爆的直接remove好了
	err := watcher.Remove(event.Name)
	if err != nil {
		log.Printf("remove path[%s] from watch failed! error: %v", event.Name, err)
	}
}
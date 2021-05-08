package main

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/gpmgo/gopm/modules/log"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
)

const (
	serverName = "myserver"
	serverPath = "../server/myserver"  // myserver服务程序文件相对于当前文件的相对路径
	confFilePath = "../config"
)

func main() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Error("new watcher failed! error: ", err)
		return
	}
	defer watcher.Close()

	err = watcher.Add(confFilePath)
	if err != nil {
		log.Error("add watcher path failed! error: ", err)
		return
	}

	go watchConfig(watcher)

	select {}
}

func watchConfig(watcher *fsnotify.Watcher) {
	fmt.Println("start to watch config file")
	for {
		select {
		case event := <-watcher.Events: {
			if event.Op&fsnotify.Write == fsnotify.Write {
				fmt.Println("config be changed!")

				pid, err := getPid(serverName)
				exePath, _ := filepath.Abs(serverPath)

				fmt.Println(exePath)

				if err != nil {
					fmt.Println("cant get pid")
					go StartProcess(exePath, []string{})
				} else {
					process, err := os.FindProcess(pid)
					if err == nil {
						fmt.Println("old process will be killed")
						//_ = process.Kill()  // 不能使用Kill()或者发送os.Kill信号
						err = process.Signal(os.Interrupt)
						if err != nil {
							fmt.Printf("send Interrupt-signal to old process failed! error: %v", err)
						} else {
							fmt.Println("old process exit")
						}
					}

					go StartProcess(exePath, []string{})
				}
			}
		}
		case err := <-watcher.Errors:
			log.Error("watcher got error: ", err)
			return
		}
	}
}

func StartProcess(exePath string, args []string) {
	attr := &os.ProcAttr{
		// files指定新进程继承的活动文件对象
		// 前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 新进程的环境变量
		Env:   os.Environ(),
	}

	process, err := os.StartProcess(exePath, args, attr)
	if err != nil {
		log.Error("start new process failed! error: ", err)
		return
	}

	fmt.Println("new process start successfully!")
	_, _ = process.Wait()
}

// 获取进程Id
func getPid(processName string) (int, error) {
	// 通过命令：wmic process get name,processid | findstr myserver 获取进程ID (windows平台)
	// 通过命令： ps -eo pid,command | grep myserver 获取进程ID (linux平台)
	buf := bytes.Buffer{}
	//cmd := exec.Command("wmic", "process", "get", "name,processid")
	cmd := exec.Command("ps", "-eo", "pid,command")
	cmd.Stdout = &buf
	err := cmd.Run()
	if err != nil {
		return -1, err
	}

	//cmd2 := exec.Command("findstr", processName)
	cmd2 := exec.Command("grep", processName)
	cmd2.Stdin = &buf
	data, err := cmd2.CombinedOutput()
	if err != nil {
		return -1, err
	}

	//fmt.Println("data")
	//fmt.Println(string(data))
	if len(data) == 0 {
		return -1, errors.New("not found")
	}

	// 正则匹配进程pid
	reg := regexp.MustCompile(`[0-9]+`)
	pid := reg.FindString(string(data))

	return strconv.Atoi(pid)
}
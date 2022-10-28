package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path"
	"strconv"
	"syscall"
)

const cgroupMemoryHierarchyCount = "/sys/fs/cgroup/memory"

func main() {
	//fmt.Println("len: ", len(os.Args))
	//s := os.Args[0]
	//fmt.Println("args: ", s)

	if os.Args[0] == "/proc/self/exe" {
		// 代码运行进来的时候，这里可以看作是一个简易的容器了
		// 通过打印当前进程的PID可发现pid=1，即当前容器已经对进程进行了隔离 (CLONE_NEWPID)
		fmt.Printf("current pid: %d\n", syscall.Getpid())

		cmd := exec.Command("sh", "-c", `stress --vm-bytes 200m --vm-keep -m 1`)
		//cmd := exec.Command("sh", "-c", `ls -al`)
		cmd.SysProcAttr = &syscall.SysProcAttr{}
		// 标准输入 输出 错误输出
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr


		if err := cmd.Run(); err != nil {
			fmt.Printf("ERROR: %v\n", err)
			os.Exit(1)
		}
	}

	cmd := exec.Command("/proc/self/exe")
	cmd.SysProcAttr = &syscall.SysProcAttr{
		Cloneflags: syscall.CLONE_NEWUTS | syscall.CLONE_NEWPID | syscall.CLONE_NEWNS,
	}
	// 标准输入 输出 错误输出
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Start(); err != nil {
		fmt.Errorf("ERROR: %v\n", err)
		os.Exit(1)
	} else {
		fmt.Printf("pid: %v\n", cmd.Process.Pid)

		// 接下来三段对 cgroup 操作
		// the hierarchy has been already created by linux on the memory subsystem
		// create a sub cgroup
		os.Mkdir(path.Join(
			cgroupMemoryHierarchyCount,
			"testmemorylimit",
		), 0755)

		// place container process in this cgroup
		ioutil.WriteFile(path.Join(
			cgroupMemoryHierarchyCount,
			"testmemorylimit",
			"tasks",
		), []byte(strconv.Itoa(cmd.Process.Pid)), 0644)

		// restrict the stress process on this cgroup
		ioutil.WriteFile(path.Join(
			cgroupMemoryHierarchyCount,
			"testmemorylimit",
			"memory.limit_int_bytes",
		), []byte("100m"), 0644)

		cmd.Process.Wait()
	}
}

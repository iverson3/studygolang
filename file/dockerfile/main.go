package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

// 支持的命令
var supportCmd = map[string]struct{}{
	"FROM": {},
	"RUN": {},
	"COPY": {},
	"WORKDIR": {},
	"CMD": {},
	"ARG": {},
	"ENV": {},
	"ENTRYPOINT": {},
	//"ADD": {},
	//"EXPOSE": {},
	//"LABEL": {},
	//"VOLUME": {},
	//"USER": {},
}

type DockerfileCmd interface {
	// FormatCheck 命令的格式检查
	FormatCheck([]string) bool
	// Exec 执行命令 (不同命令执行完成后的返回值不同，不能抽象成接口)
	//Exec([]string) error
}

type DockerfileFromCmd struct {
}
func (dc *DockerfileFromCmd) FormatCheck(cmdLine []string) bool {
	return true
}
func (dc *DockerfileFromCmd) Exec(cmdLine []string) (containerId string, err error) {
	return "", nil
}

type DockerfileCopyCmd struct {
}
func (dc *DockerfileCopyCmd) FormatCheck(cmdLine []string) bool {
	return true
}
func (dc *DockerfileCopyCmd) Exec(cmdLine []string) (err error) {
	return nil
}

func main() {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	dockerFilePath := filepath.Join(pwd, "Dockerfile")

	// 提取dockerfile文件中所有的命令行
	cmdLines, err := scanDockerfile(dockerFilePath)
	if err != nil {
		panic(err)
	}

	// 解析dockerfile命令
	err = parseDockerfileCommand(cmdLines)
	if err != nil {
		return
	}
}

func parseDockerfileCommand(cmdLines [][]string) error {
	for _, cmdLine := range cmdLines {
		fmt.Println(cmdLine)
	}
	return nil
}

// 扫描并获取dockerfile文件中的每一行命令
func scanDockerfile(Dockerfile string) ([][]string, error) {
	f, err := os.Open(Dockerfile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	cmdLines := make([][]string, 0)
	reader := bufio.NewReader(f)
	for {
		// 逐行读取文件内容
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}

		// 忽略空行和注释行
		if len(line) != 0 && line[0] != '#' {
			cmdArgs := strings.Split(string(line), " ")
			if len(cmdArgs) == 1 {
				return nil, fmt.Errorf("command format wrong")
			}

			// 检查命令是否支持
			if _, ok := supportCmd[cmdArgs[0]]; !ok {
				return nil, fmt.Errorf("not supported command: %s", cmdArgs[0])
			}

			cmdLines = append(cmdLines, cmdArgs)
		}
	}

	return cmdLines, nil
}

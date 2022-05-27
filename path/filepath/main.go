package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

func main() {
	p, err := GetCurrentPath()
	if err != nil {
		panic(err)
	}

	fmt.Println("current path: ", p)
}

func GetCurrentPath() (string, error) {
	if runtime.GOOS == "windows" {
		return os.Getwd()
	} else {
		return filepath.Abs(filepath.Dir(os.Args[0]))
	}
}
package main

import (
	"fmt"
	"github.com/go-ini/ini"
)

func main() {
	cfg, err := ini.Load("env.ini")
	if err != nil {
		fmt.Println("load failed! error: ", err)
		return
	}

	section, err := cfg.GetSection("proxy")
	if err != nil {
		fmt.Println("no this section.")
		return
	}

	path, err := section.GetKey("path")
	if err != nil {
		fmt.Println("no this key: path")
		return
	}

	fmt.Println("path = ", path)
}

package main

import (
	"fmt"
	"studygolang/interface/best_practice/service"
)

func main() {
	srv := service.NewService("conn")
	posts, err := srv.ListPosts()
	if err != nil {
		panic(err)
	}

	fmt.Println(posts)
}

package main

import (
	"log"
	"net"
	"net/rpc/jsonrpc"
	rpcdemo "studygolang/rpc"
)

func main() {
	conn, err := net.Dial("tcp", ":1234")
	if err != nil {
		panic(err)
	}
	client := jsonrpc.NewClient(conn)

	var result float64
	err = client.Call("DemoService.Div", rpcdemo.Args{10, 3}, &result)
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Printf("rpc call result is: %v", result)
	}

	err = client.Call("DemoService.Div", rpcdemo.Args{10, 0}, &result)
	if err != nil {
		log.Printf("error: %v", err)
	} else {
		log.Printf("rpc call result is: %v", result)
	}

	defer conn.Close()
	defer client.Close()
}























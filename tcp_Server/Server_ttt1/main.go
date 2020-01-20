package main

import (
	"fmt"
	"net"
)

//tcp 启动一个server端
func f(coon net.Conn) {
	var b [128]byte
	for {
		n, err := coon.Read(b[:])
		if err != nil {
			fmt.Println("read failed err :", err)
			return
		}
		fmt.Println(string(b[:n]))
	}

}

func main() {
	//1。本地启动一个端口。并监听
	Listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listener failed err :", err)
		return
	}
	for {
		coon, err := Listener.Accept()
		if err != nil {
			fmt.Println("Accept failed err :", err)
			return
		}
		go f(coon)
	}

}

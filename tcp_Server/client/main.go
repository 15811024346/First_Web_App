package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	coon, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("dial 127.0.0.1:20000 failed err : ", err)
		return
	}
	reders := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("请说话:")
		msg, _ := reders.ReadString('\n')
		msg = strings.TrimSpace(msg)
		if msg == "exit" {
			break
		}
		coon.Write([]byte(msg))
	}
	coon.Close()
}

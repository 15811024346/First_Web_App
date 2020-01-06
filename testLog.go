package main

import (
	log "./utils/log"
	"time"
)

func main() {
	for {

		log.Debug("11111111")
		log.Release("11111")
		time.Sleep(time.Second * 2)

		//log.Fatal("11111111")
	}
}

package main

import (
	"fmt"
	"time"
)

var conf *config

var ctr *Token

var bn int64 = 0

func main() {
	conf = initConfig()

	ctr = initGeth()

	for {
		if bn != blockNumber() {
			bn = blockNumber()
			fmt.Println(bn)
		}
		time.Sleep(1 * time.Second)
	}
}

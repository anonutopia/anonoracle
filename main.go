package main

import (
	"time"
)

var conf *config

var ctr *Token

var bn int64 = 0

var s *Storage

func main() {
	conf = initConfig()

	ctr = initGeth()

	s = initStorage()

	for {
		if bn != blockNumber() {
			bn = blockNumber()
			updateEurPrice()
		}
		time.Sleep(1 * time.Second)
	}
}

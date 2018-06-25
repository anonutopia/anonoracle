package main

import (
	"time"
)

var conf *config

var ctr *Token

var s *Storage

func main() {
	conf = initConfig()

	ctr = initGeth()

	s = initStorage()

	for {
		if s.BlockNumber != blockNumber() {
			s.BlockNumber = blockNumber()
			updateEurPrice()
		}
		time.Sleep(1 * time.Second)
	}
}

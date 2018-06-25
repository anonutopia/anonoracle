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

	counter := 0

	for {
		if s.BlockNumber != blockNumber() {
			counter++
			s.BlockNumber = blockNumber()
			updateEurPrice(&counter)
		}
		time.Sleep(1 * time.Second)
	}
}

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
		bn, err := blockNumber()
		if err == nil && s.BlockNumber != bn {
			counter++
			s.BlockNumber = bn
			updateEurPrice(&counter)
		}
		// log.Printf("Tick: %d", bn)
		time.Sleep(1 * time.Second)
	}
}

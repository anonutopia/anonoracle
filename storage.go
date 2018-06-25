package main

type Storage struct {
	EurPrice    int64
	BlockNumber int64
}

func initStorage() *Storage {
	s := &Storage{EurPrice: 0}
	return s
}

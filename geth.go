package main

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var conn *ethclient.Client

func initGeth() *Token {
	var err error

	conn, err = ethclient.Dial(conf.Geth)
	if err != nil {
		log.Printf("Failed to connect to the Ethereum client: %v", err)
		return nil
	}

	token, err := NewToken(common.HexToAddress(conf.ContractAddress), conn)
	if err != nil {
		log.Printf("Failed to instantiate a Token contract: %v", err)
		return nil
	}

	return token
}

func blockNumber() int64 {
	header, err := conn.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return header.Number.Int64()
}

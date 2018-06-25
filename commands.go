package main

import (
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func updateEurPrice(counter *int) {
	ep, _ := priceCheckRequest()
	add := common.HexToAddress("0x712c8029CBdFa2E45E6d13f0ae753071BE59daa2")

	key, err := ioutil.ReadFile("keyfile.json")
	if err != nil {
		log.Print(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(key)), conf.Password)
	if err != nil {
		log.Fatalf("could not create auth: %v", err)
	}

	eurPrice := int64(math.Pow(10, 18) / ep.EUR)

	if eurPrice != s.EurPrice {
		if *counter == 100 || float64(eurPrice)/float64(s.EurPrice) >= 1.01 || float64(eurPrice)/float64(s.EurPrice) <= 0.99 {
			ctr.UpdateCurrencyPrice(auth, add, big.NewInt(eurPrice))
			log.Printf("Updated: %d %f", eurPrice, float64(eurPrice)/float64(s.EurPrice))
			s.EurPrice = eurPrice
			*counter = 0
		}
	}
}

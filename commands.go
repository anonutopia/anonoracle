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
	add := common.HexToAddress(conf.EurAddress)

	key, err := ioutil.ReadFile("keyfile.json")
	if err != nil {
		log.Print(err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(string(key)), conf.Password)
	if err != nil {
		log.Fatalf("could not create auth: %v", err)
	}

	auth.GasLimit = 40000

	eurPrice := int64(math.Pow(10, 18) / ep.EUR)

	if eurPrice != s.EurPrice {
		priceRatio := float64(s.EurPrice) / float64(eurPrice)
		if *counter == 100 || priceRatio >= 1.01 || priceRatio <= 0.99 {
			_, err := ctr.UpdateCurrencyPrice(auth, add, big.NewInt(eurPrice))
			if err == nil {
				log.Printf("Updated: %d %f", eurPrice, float64(eurPrice)/float64(s.EurPrice))
			} else {
				log.Printf("Error: %s", err)
			}
			s.EurPrice = eurPrice
			*counter = 0
		}
	}
}

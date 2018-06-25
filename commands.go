package main

import (
	"log"
	"math"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

const key = "{\"address\":\"c910f4f3b67799ef20774b05691ebffe23da5ee5\",\"crypto\":{\"cipher\":\"aes-128-ctr\",\"ciphertext\":\"7ed7d3c1e8be58fc54ddb71f5664e62bc342684817e5afbe24f3d43e629ffc67\",\"cipherparams\":{\"iv\":\"6370a979083a69917c42d3cf484f3934\"},\"kdf\":\"scrypt\",\"kdfparams\":{\"dklen\":32,\"n\":262144,\"p\":1,\"r\":8,\"salt\":\"bfc56c046ca58246f00672de6033a5f6ce341225ba32e8df093e01ed80c5bb1a\"},\"mac\":\"3509ac30ac433788a15fea2be9f121de50a121cc4a0ce7a84817283776d2404d\"},\"id\":\"518263e7-8427-42f4-80e5-2d993ee06c08\",\"version\":3}"

func updateEurPrice() {
	ep, _ := priceCheckRequest()
	add := common.HexToAddress("0x712c8029CBdFa2E45E6d13f0ae753071BE59daa2")

	auth, err := bind.NewTransactor(strings.NewReader(key), conf.Password)
	if err != nil {
		log.Fatalf("could not create auth: %v", err)
	}

	eurPrice := int64(math.Pow(10, 18) / ep.EUR)

	if eurPrice != s.EurPrice {
		ctr.UpdateCurrencyPrice(auth, add, big.NewInt(eurPrice))
		s.EurPrice = eurPrice
		log.Printf("Updated: %d", eurPrice)
	}
}

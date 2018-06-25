package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func doRequest(url string, method string, body interface{}, resp interface{}) error {
	cl := http.Client{}

	var req *http.Request
	var err error

	if body != nil {
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(body)
		req, err = http.NewRequest(method, url, b)
	} else {
		req, err = http.NewRequest(method, url, nil)
	}

	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	res, err := cl.Do(req)

	if err == nil {
		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return err
		}
		if res.StatusCode != 200 {
			log.Printf("[doRequest] Error, body: %s", string(body))
		}
		json.Unmarshal(body, resp)
	} else {
		return err
	}

	return nil
}

func priceCheckRequest() (*EthPrices, error) {
	ep := &EthPrices{}
	err := doRequest("https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=BTC,USD,EUR,HRK", http.MethodGet, nil, ep)
	return ep, err
}

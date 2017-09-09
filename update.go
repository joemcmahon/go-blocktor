package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// UpdateFromTOR get all addresses from the past 16 hours.
func UpdateFromTOR(address string) []string {
	url := "https://check.torproject.org/cgi-bin/TorBulkExitList.py?ip="
	res, errGet := http.Get(url + address)

	if errGet != nil {
		panic(errGet)
	}

	defer res.Body.Close()
	body, errBody := ioutil.ReadAll(res.Body)

	if errBody != nil {
		panic(errBody)
	}

	result := strings.Split(string(body), "\n")
	return result
}

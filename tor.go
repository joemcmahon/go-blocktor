package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// UpdateAddresses get all addresses from the past 16 hours.
func UpdateAddresses(address string) []string {
	res, _ := http.Get("https://check.torproject.org/cgi-bin/TorBulkExitList.py?ip=" + address)
	defer res.Body.Close()

	body, _ := ioutil.ReadAll(res.Body)
	result := strings.Split(string(body), "\n")

	return result
}

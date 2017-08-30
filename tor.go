package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

// UpdateAddresses get all addresses from the past 16 hours.
func UpdateAddresses(address string) []string {
	response, _ := http.Get("https://check.torproject.org/cgi-bin/TorBulkExitList.py?ip=" + address)
	defer response.Body.Close()

	body, _ := ioutil.ReadAll(response.Body)
	result := strings.Split(string(body), "\n")

	return result
}

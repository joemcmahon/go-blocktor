package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os/exec"
	"runtime"
	"strings"
)

// Exec execute the commands.
func Exec(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	err := c.Run()

	if err != nil {
		panic(err)
	}
}

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

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Operating system not supported.")
		return
	}

	flagAddress := flag.String("address", "1.1.1.1", "IP Address to block TOR network")
	flagRules := flag.String("rules", "", "Custom script with initial rules")

	flag.Parse()

	fmt.Println("Updating...")
	TORaddresses := UpdateFromTOR(*flagAddress)

	fmt.Println("Cleaning...")

	Exec("iptables", "-F")
	Exec("iptables", "-t", "nat", "-F")

	fmt.Println("Initial rules...")
	if *flagRules != "" {
		Exec(*flagRules)
	}

	fmt.Println("Blocking...")

	for _, address := range TORaddresses {
		if len(address) < 16 && len(address) > 7 {
			Exec("iptables", "-A", "INPUT", "-s", address, "-j", "DROP")
		}
	}
}

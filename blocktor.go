package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Updating...")

	// parameters to use.
	addrDestiny := flag.String("address", "1.1.1.1", "IP Address")
	initialRules := flag.String("rules", "./start.sh", "Script initial rules")
	clearRules := flag.String("clear", "false", "Clear rules")

	flag.Parse()

	allAddresses := UpdateAddresses(*addrDestiny)

	if *clearRules != "true" {
		ClearAllNAT()
		ClearAllRules()

		InitialRules(*initialRules)

		fmt.Println("Blocking...")
		for _, addr := range allAddresses {
			if len(addr) < 16 && len(addr) > 7 {
				BlockRequestFrom(addr)
			}
		}
	} else {
		ClearAllRules()
	}
}

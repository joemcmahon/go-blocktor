package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Updating...")

	// parameters to use.
	addrDestiny := flag.String("address", "1.1.1.1", "IP Address")
	initialRules := flag.String("initial-rules", "./start.sh", "Script with initial rules")

	flag.Parse()

	allAddresses := UpdateAddresses(*addrDestiny)

	// clear old rules.
	ClearAllNAT()
	ClearAllRules()

	// load new rules.
	InitialRules(*initialRules)

	// block addresses from list.
	fmt.Println("Blocking...")
	for _, addr := range allAddresses {
		if len(addr) < 16 && len(addr) > 7 {
			BlockRequestFrom(addr)
		}
	}
}

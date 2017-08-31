package main

import (
	"flag"
	"fmt"
)

func main() {
	fmt.Println("Updating TOR nodes...")
	
	// parameters to use.
	addrDestiny := flag.String("address", "1.1.1.1", "IP Address")
	initialRules := flag.String("initial-rules", "./start.sh", "Script with initial rules")

	flag.Parse()

	allAddresses := UpdateAddresses(*addrDestiny)
	ClearAllRules()

	InitialRules(*initialRules)

	fmt.Println("Blocking addresses...")
	
	// block addresses from list.
	for _, addr := range allAddresses {
		if len(addr) < 16 && len(addr) > 7 {
			BlockRequestFrom(addr)
		}
	}
}

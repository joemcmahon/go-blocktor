package main

import (
	"fmt"
)

func main() {
	fmt.Println("Updating TOR list addresses...")

	allAddresses := UpdateAddresses()

	for _, addr := range allAddresses {
		if len(addr) < 16 && len(addr) > 7 {
			fmt.Println("Blocking ", addr, "...")
			BlockRequestFrom(addr)
		}
	}
}

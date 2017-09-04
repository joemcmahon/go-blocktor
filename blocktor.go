package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS == "linux" {
		fmt.Println("Operating system not supported.")
		return
	}

	fmt.Println("Updating...")

	// parameters to use.
	flagAddress := flag.String("address", "1.1.1.1", "IP Address")
	flagRules := flag.String("rules", "./start.sh", "Script initial rules")
	flagClear := flag.String("clear", "false", "Clear rules")

	flag.Parse()

	// get TOR addresses.
	TORaddresses := UpdateAddresses(*flagAddress)

	if *flagClear != "true" {
		ClearAll()
		InitialRules(*flagRules)

		fmt.Println("Blocking...")

		for _, addr := range TORaddresses {
			if len(addr) < 16 && len(addr) > 7 {
				BlockRequestFrom(addr)
			}
		}
	} else {
		ClearAllRules()
	}
}

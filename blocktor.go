package main

import (
	"flag"
	"fmt"
	"runtime"
)

func main() {
	if runtime.GOOS != "linux" {
		fmt.Println("Operating system not supported.")
		return
	}

	flagA := flag.String("address", "1.1.1.1", "IP Address")
	flagR := flag.String("rules", "./start.sh", "Script initial rules")
	flagC := flag.String("clear", "false", "Clear rules")

	flag.Parse()

	fmt.Println("Updating...")
	TORaddresses := UpdateFromTOR(*flagA)

	if *flagC == "true" {
		fmt.Println("Cleaning...")
		ClearBlockTOR()
	} else {
		ClearAll()
		InitialRules(*flagR)

		fmt.Println("Blocking...")

		for _, addr := range TORaddresses {
			if len(addr) < 16 && len(addr) > 7 {
				BlockRequestFrom(addr)
			}
		}
	}
}

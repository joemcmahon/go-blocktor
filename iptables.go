package main

import "os/exec"
import "fmt"

// BlockRequestFrom represent the command iptables to block address.
func BlockRequestFrom(address string) {
	ports := []string{"80", "443"}

	for _, port := range ports {
		line := "iptables -A INPUT -s", address, " -p tcp --dport ", port, " -j DROP"
		fmt.Println(line)

		exec.Command(line).Output()
	}
}

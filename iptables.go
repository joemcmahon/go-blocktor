package main

import (
	"os"
	"os/exec"
)

// BlockRequestFrom represent the command iptables to block address.
func BlockRequestFrom(address string) {
	ports := []string{"80", "443"}

	os.Setenv("PATH", "/sbin:/bin:/usr/bin:/usr/sbin")
	os.Getenv("PATH")

	for _, port := range ports {
		cmd := exec.Command("iptables", "-A", "INPUT", "-s", address, "-p", "tcp", "--dport", port, "-j", "DROP")
		cmd.Run()
	}
}

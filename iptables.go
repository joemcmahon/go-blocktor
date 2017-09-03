package main

import (
	"os/exec"
)

// InitRules load initials rules.
func InitialRules(scriptPath string) {
	c := exec.Command(scriptPath)
	c.Run()
}

// ClearAllNAT clear all rules of nat iptables.
func ClearAllNAT() {
	c := exec.Command("iptables", "-t", "nat", "-F")
	c.Run()
}

// ClearAllRules clear all rules of iptables.
func ClearAllRules() {
	c := exec.Command("iptables", "-F")
	c.Run()
}

//ClearAll clear all rules.
func ClearAll() {
	ClearAllNAT()
	ClearAllRules()
}

// BlockRequestFrom block requests from address in a port 80 and 443.
func BlockRequestFrom(address string) {
	ports := []string{"80", "443"}

	for _, port := range ports {
		c := exec.Command("iptables", "-A", "INPUT", "-s", address, "-p", "tcp", "--dport", port, "-j", "DROP")
		c.Run()
	}
}

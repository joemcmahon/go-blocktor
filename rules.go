package main

import (
	"os/exec"
)

// Exec execute the commands.
func Exec(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	c.Run()
}

// InitRules load initials rules.
func InitialRules(scriptPath string) {
	Exec(scriptPath)
}

// ClearAllNAT clear all rules of nat iptables.
func ClearAllNAT() {
	Exec("iptables", "-t", "nat", "-F")
}

// ClearAllRules clear all rules of iptables.
func ClearAllRules() {
	Exec("iptables", "-F")
}

//ClearAll clear all rules.
func ClearAll() {
	ClearAllNAT()
	ClearAllRules()
}

// BlockRequestFrom block requests from address in a port 80 and 443.
func BlockRequestFrom(address string) {
	Exec("iptables", "-A", "INPUT", "-s", address, "-j", "DROP")
}

package main

import (
	"os/exec"
)

// Exec execute the commands.
func Exec(cmd string, args ...string) {
	c := exec.Command(cmd, args...)
	err := c.Run()

	if err != nil {
		panic(err)
	}
}

// InitialRules load initials rules.
func InitialRules(pathScript string) {
	Exec(pathScript)
}

// ClearNAT clear all rules of nat iptables.
func ClearNAT() {
	Exec("iptables", "-t", "nat", "-F")
}

// ClearBlockTOR clear all rules of iptables.
func ClearBlockTOR() {
	Exec("iptables", "-F")
}

//ClearAll clear all rules.
func ClearAll() {
	ClearNAT()
	ClearBlockTOR()
}

// BlockRequestFrom block requests from address in a port 80 and 443.
func BlockRequestFrom(address string) {
	Exec("iptables", "-A", "INPUT", "-s", address, "-j", "DROP")
}

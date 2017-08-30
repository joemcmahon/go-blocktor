package main

import (
	"os"
	"os/exec"
	"syscall"
)

// BlockRequestFrom represent the command iptables to block address.
func BlockRequestFrom(address string) {
	ports := []string{"80", "443"}

	binary, lookErr := exec.LookPath("iptables")
	if lookErr != nil {
		panic(lookErr)
	}

	for _, port := range ports {
		args := []string{"-A INPUT -s", address, "-p tcp", "--destination-port 80", "-j DROP"}
		env := os.Environ()

		execErr := syscall.Exec(binary, args, env)
		if execErr != nil {
			panic(execErr)
		}
	}

}

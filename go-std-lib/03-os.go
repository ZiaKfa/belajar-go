package main

import (
	"fmt"
	"os"
)

func main() {
	//os.Args is a slice of strings of the command line arguments
	args := os.Args
	for _, arg := range args {
		fmt.Println(arg)
	}

	//os.Hostname() returns the host name reported by the kernel
	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(hostname)
	}
}

package main

import (
	"fmt"
	"net"

	"github.com/jonathanruiz/iptracker-app/cmd"
)

func main() {

	// Get all of the IP addresses of the machine.
	addrs := cmd.GetAllPrivateIP()

	// Get the IP addresses of the machine.
	ipAddrs := cmd.GetPrivateIP(addrs)

	// Print out the IP addresses of the machine.
	for _, ip := range ipAddrs {
		fmt.Println(ip)
	}

	var ips = ipAddrs
	for _, ip := range ips {
		IP := net.ParseIP(ip)
		fmt.Println(IP, ":", cmd.IsPublicIP(IP))
	}
}

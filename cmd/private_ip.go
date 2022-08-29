package cmd

import (
	"fmt"
	"net"
)

// getAllPrivateIP returns all of the IP addresses of the machine.
func getAllPrivateIP() []net.Addr {
	// Get all of the IP addresses associated with the machine.
	addrs, err := net.InterfaceAddrs()

	// If there is an error, print it out and return nil.
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// return the IP addresses of the machine.
	return addrs
}

// Get single Private IP address of the machine.
func getPrivateIP(ips []net.Addr) []string {
	// Create an array to hold the IP addresses.
	var ipAddrs []string

	// Loop through all of the addresses returned by the machine.
	for _, a := range ips {
		// Check to see if the address is a valid IPv4 address.
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				// If it is a valid IPv4 address, add it to the array.
				ipAddrs = append(ipAddrs, ipnet.IP.String())
			}
		}
	}

	// Return the array of IP addresses.
	return ipAddrs
}

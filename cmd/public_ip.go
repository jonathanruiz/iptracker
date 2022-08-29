package cmd

import (
	"net"
)

// Class       Starting IP Address     Ending IP Address    # of Hosts
// A           10.0.0.0                10.255.255.255       16,777,216
// B           172.16.0.0              172.31.255.255       1,048,576
// C           192.168.0.0             192.168.255.255      65,536
// Link-local  169.254.0.0             169.254.255.255      65,536
// Local       127.0.0.0               127.255.255.255      16777216

func IsPublicIP(IP net.IP) bool {
	if IP.IsLoopback() || IP.IsLinkLocalMulticast() || IP.IsLinkLocalUnicast() {
		return false
	}
	if ip4 := IP.To4(); ip4 != nil {
		switch {
		case ip4[0] == 10:
			return false
		case ip4[0] == 172 && ip4[1] >= 16 && ip4[1] <= 31:
			return false
		case ip4[0] == 192 && ip4[1] == 168:
			return false
		default:
			return true
		}
	}
	return false
}

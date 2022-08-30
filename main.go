package main

import (
	"fmt"

	"github.com/jonathanruiz/iptracker-app/cmd"
)

func main() {
	fmt.Println("Private IP : ", cmd.GetPrivateIP())
	fmt.Println("Public IP: ", cmd.GetPublicIP())
}

package main

import (
	"fmt"
	"many-motd/cmd"
	"many-motd/config"
)

func main() {

	config.GetConfig()

	msg := cmd.Generate()
	fmt.Println(msg)
}

package main

import "github.com/diegobernardes/ndn"

func main() {
	// start the application
	c, _ := ndn.NewClient()
	c.Start()
}

package main

import (
	"ndn/application"
)

func main() {
	// start the application
	c, _ := application.NewClient(application.ClientConfig{})
	c.Start()
}

/*
	fazer um servidor com cobra e viper, carregar a config do hcl e enviar para o client ligar.
*/

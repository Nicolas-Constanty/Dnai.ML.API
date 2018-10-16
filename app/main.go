package main

import (
	"flag"
	"fmt"
	"github.com/Nicolas-Constanty/Dnai.ML.API/api"
)

func main() {
	// Arguments handing
	password := flag.String("password", "", "Defining server password")
	routes := flag.String("routes", "", "Defining routes conf file")
	port := flag.String("port", "8000", "Defining server port")
	flag.Parse()
	if *password == "" {
		fmt.Println("You need to define a password with -password!")
		return
	}
	if *routes == "" {
		fmt.Println("You need to define a routes.conf with -routes!")
		return
	}

	// Create new API
	mlApi := api.NewMlApi()
	// Start API with routes.conf, a password and a port
	mlApi.Start(*routes, *password, *port)
}



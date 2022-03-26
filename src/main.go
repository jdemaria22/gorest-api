package main

import (
	routes "demoproject/src/api"
	config "demoproject/src/config"

	"fmt"
)

func main() {
	fmt.Println("Initializing Config...")
	config.InitConfig()
	fmt.Println("Initializing Routes...")
	routes.InitialiseRoutes()
}

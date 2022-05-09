package main

import (
	routes "demoproject/src/api"
	config "demoproject/src/config"
	logging "demoproject/src/log"
)

func main() {
	logging.Info("Initializing config")
	config.InitConfig()
	logging.Info("Initializing routes")
	routes.InitialiseRoutes()
}

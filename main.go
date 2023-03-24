package main

import "ch6/routes"

func main() {
	var PORT = ":8080"

	routes.StartServer().Run(PORT)
}

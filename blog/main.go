package main

import "blog/routes"

func main() {
	router := routes.SetupRouter()

	router.Run()
}

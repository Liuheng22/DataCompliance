package main

import "DataCompliance/router"

func main() {
	router := router.StartRouter()
	router.Run(":8001")
}

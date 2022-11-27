package main

import (
	"member-site-go/controller"
)

func main() {
	router := controller.GetRouter()
	router.Run(":8080")
}

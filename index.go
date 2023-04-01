package main

import (
	config "library/src/configs"
	"library/src/routers"
)

func main() {
	config.Connect("mongodb://localhost:27017", "test")
	routers.InitServer()
}

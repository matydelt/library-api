package main

import (
	"library/src/configs"
	"library/src/routers"
)

func main() {
	configs.Connect("mongodb://localhost:27017", "test")

	routers.InitServer()
}

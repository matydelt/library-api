package main

import (
	"fmt"
	"library/src/configs"
	"library/src/routers"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error cargando archivo .env")
	}
	user := os.Getenv("MONGO_USER")
	pass := os.Getenv("MONGO_PASS")
	cluster := os.Getenv("MONGO_CLUSTER")
	db := os.Getenv("MONGO_DB")
	conecctionUrl := fmt.Sprintf("mongodb+srv://%s:%s@%s/", user, pass, cluster)
	configs.MongoConnect(conecctionUrl, db)
	routers.InitServer()
}

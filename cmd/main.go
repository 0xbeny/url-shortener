package main

import (
	"os"
	"url-shortener/router"
	"url-shortener/shared/database"
)

func main() {

	 database.Connect(os.Getenv("REDIS_URL"))

	app:=router.Wrapper()

	app.Run(":8080")

}

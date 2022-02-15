package main

import (
	"url-shortener/router"
	"url-shortener/shared/database"
)

func main() {

	 database.Connect("localhost:6379")

	app:=router.Wrapper()


	app.Run(":8080")

}

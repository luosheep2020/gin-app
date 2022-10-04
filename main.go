package main

import (
	"gin-app/db"
	"gin-app/routers"
)

func main() {
	app := routers.InitRouter()
	db.InitDb()
	app.Run()
}

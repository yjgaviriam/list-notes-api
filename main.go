package main

import (
	app "list-notes/bootstrap"
)

func main() {
	var a app.App
	a.CreateConnection()
	//a.Migrate()
	a.CreateRoutes()
	a.Run()
}

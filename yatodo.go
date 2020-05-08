package main

import (
	"fmt"
	"github.com/fredericorecsky/yatodo/app"
	"github.com/fredericorecsky/yatodo/web"
)

func main() {

	fmt.Println("Migrate")

	app := new(app.Todo)

	app.Migrate()



	//// if cmd line argument migrate
	//
	web.StartGin()
}

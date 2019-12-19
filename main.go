package main

import (
	"cafekalaa/api/app"

	"cafekalaa/api/config"
)

func main() {
	Config := config.GetConfig()
	App := &app.App{}
	App.Initialize(Config)
	App.Run(":3000")

}

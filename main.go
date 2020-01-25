package main

import (
	"github.com/jindrichskupa/supervisord-statuspage/app"
	"github.com/jindrichskupa/supervisord-statuspage/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(config.ListenAddress())
}

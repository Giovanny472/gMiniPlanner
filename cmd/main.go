package main

import (
	"github.com/Giovanny472/gMiniPlanner/app"
	"github.com/Giovanny472/gMiniPlanner/interfaces"
)

func main() {
	app := app.NewFactoryApp().MakeApp(interfaces.AppTypeDesktop)
	app.Exec()
}

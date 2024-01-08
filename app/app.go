package app

import (
	"github.com/Giovanny472/gMiniPlanner/interfaces"
	"github.com/Giovanny472/gMiniPlanner/ui/widgets"
)

type application struct {
	gui interfaces.IUI
}

var apl *application

func NewApp() interfaces.IApp {
	if apl == nil {
		apl = &application{}
	}
	return apl
}

func (ap *application) Exec() {

	ap.gui = widgets.NewFyneMain()
	ap.gui.Show()
}

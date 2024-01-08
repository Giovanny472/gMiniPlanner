package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"

	"github.com/Giovanny472/gMiniPlanner/interfaces"
	"github.com/Giovanny472/gMiniPlanner/ui"
)

type fynemain struct{}

var fm *fynemain

func NewFyneMain() interfaces.IUI {

	if fm == nil {
		fm = &fynemain{}
	}
	return fm
}

func (f *fynemain) Show() {

	app := app.NewWithID("io.ve.planner")
	win := app.NewWindow("Мини-План")

	win.Resize(fyne.NewSize(ui.Width, ui.Height))
	win.Show()

	app.Run()
}

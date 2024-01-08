package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"

	"github.com/Giovanny472/gMiniPlanner/interfaces"
	"github.com/Giovanny472/gMiniPlanner/ui"
)

type fynemain struct {
	menu      *fyne.Container
	detail    *fyne.Container
	listPlans *widget.List
}

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

	// layouts
	f.menu = container.NewVBox()
	f.detail = container.NewStack()

	//split
	split := container.NewHSplit(f.menu, f.detail)
	split.Offset = 0.2
	win.SetContent(split)

	win.Show()
	app.Run()
}

package widgets

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"

	"github.com/Giovanny472/gMiniPlanner/interfaces"
	"github.com/Giovanny472/gMiniPlanner/ui"
)

type fynemain struct {
	menu   *fyne.Container
	detail *fyne.Container
	//listPlans *widget.List

	win fyne.Window
}

var fm *fynemain

func NewFyneMain() interfaces.IUI {

	if fm == nil {
		fm = &fynemain{}
	}
	return fm
}

func (f *fynemain) Show() {

	app := app.NewWithID("com.viem.planner")
	f.win = app.NewWindow("Мини-План")
	f.win.Resize(fyne.NewSize(ui.Width, ui.Height))

	// layouts
	f.menu = container.NewVBox()
	f.detail = container.NewStack()

	// bottom
	options := container.NewGridWithColumns(2,
		widget.NewButton("+", f.onClickAddTask),
		widget.NewButton("-", f.onClickDelTask),
	)
	ctLeft := container.NewBorder(nil, options, nil, nil, f.menu)

	//split
	split := container.NewHSplit(ctLeft, f.detail)
	split.Offset = 0.2
	f.win.SetContent(split)

	f.win.Show()
	app.Run()
}

func (f *fynemain) onClickAddTask() {

	txtTask := widget.NewEntry()
	txtTask.SetPlaceHolder("название задачи")
	ct1 := container.NewBorder(nil, nil, nil, nil, txtTask)

	FormTask := widget.NewFormItem("", ct1)
	listItems := make([]*widget.FormItem, 1)
	listItems[0] = FormTask

	dialog.ShowForm("Tasks", "да", "Отмена", listItems,
		func(a bool) {

			if a && len(txtTask.Text) > 0 {
				fmt.Println("OK")
			}
		},
		f.win)
}

func (f *fynemain) onClickDelTask() {

}

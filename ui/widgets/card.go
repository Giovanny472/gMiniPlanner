package widgets

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
	"github.com/Giovanny472/gMiniPlanner/interfaces"
)

func NewCard(data interfaces.ICard) fyne.CanvasObject {

	pb := widget.NewProgressBar()
	pb.Min, pb.Max = 0, 100
	pb.SetValue(50)

	return widget.NewCard(data.Title(), "", pb)
}

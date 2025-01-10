package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type DialogStruct struct {
	parent  fyne.Window
	content fyne.CanvasObject
	dialog  Dialog
}

func NewCustomDialog(parent fyne.Window, message string) *DialogStruct {
	content := widget.NewLabel(message)
	d := &DialogStruct{
		parent:  parent,
		content: content,
	}
	d.dialog = dialog.NewCustom("Custom Dialog", "Dismiss", content, parent)
	return d
}
func (d *DialogStruct) Show() {
	d.dialog.Show()
}

func (d *DialogStruct) Hide() {
	d.dialog.Hide()
}

func (d *DialogStruct) SetDismissText(label string) {
	d.dialog.SetDismissText(label)
}

func (d *DialogStruct) SetOnClosed(closed func()) {
	d.dialog.SetOnClosed(closed)
}

func (d *DialogStruct) Refresh() {
	d.dialog.Refresh()
}

func (d *DialogStruct) Resize(size fyne.Size) {
	d.dialog.Resize(size)
}

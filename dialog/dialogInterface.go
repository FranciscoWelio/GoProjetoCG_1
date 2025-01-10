package dialog

import "fyne.io/fyne/v2"

type Dialog interface {
	Show()
	Hide()
	SetDismissText(label string)
	SetOnClosed(closed func())
	Refresh()
	Resize(size fyne.Size)

	MinSize() fyne.Size
}

package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

type DialogStruct struct {
	parent  fyne.Window
	content fyne.CanvasObject
	dialog  dialog.Dialog
}

var _ Dialog = (*DialogStruct)(nil)

func NewCustomDialog(parent fyne.Window, message string) *DialogStruct {
	content := widget.NewLabel(message)
	d := &DialogStruct{
		parent:  parent,
		content: content,
		dialog:  dialog.NewCustom("Atenção", "Fechar", content, parent),
	}
	return d
}

// Exibe o diálogo
func (d *DialogStruct) Show() {
	if d.dialog != nil {
		d.dialog.Show()
	}
}

// Oculta o diálogo
func (d *DialogStruct) Hide() {
	if d.dialog != nil {
		d.dialog.Hide()
	}
}

// Altera o texto do botão de fechamento
func (d *DialogStruct) SetDismissText(label string) {
	if d.dialog != nil {
		d.dialog.SetDismissText(label)
	}
}

// Define o callback ao fechar o diálogo
func (d *DialogStruct) SetOnClosed(closed func()) {
	if d.dialog != nil {
		d.dialog.SetOnClosed(closed)
	}
}

// Atualiza o conteúdo do diálogo
func (d *DialogStruct) Refresh() {
	if d.dialog != nil {
		d.dialog.Refresh()
	}
}

// Redimensiona o diálogo
func (d *DialogStruct) Resize(size fyne.Size) {
	if d.dialog != nil {
		d.dialog.Resize(size)
	}
}

// Retorna o tamanho mínimo do diálogo
func (d *DialogStruct) MinSize() fyne.Size {
	if d.dialog != nil {
		return d.dialog.MinSize()
	}
	return fyne.NewSize(0, 0)
}

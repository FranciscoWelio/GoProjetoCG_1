package util

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateButtons cria e retorna o frame com os botões
func CreateButtons(app fyne.App, canvasContent *fyne.Container, width, height int) fyne.CanvasObject {
	x0 := widget.NewEntry()
	x0.SetPlaceHolder("X")
	y0 := widget.NewEntry()
	y0.SetPlaceHolder("Y")
	xEnd := widget.NewEntry()
	xEnd.SetPlaceHolder("X Final")
	yEnd := widget.NewEntry()
	yEnd.SetPlaceHolder("Y Final")
	label := widget.NewLabel("Controles:")
	gridInput := container.New(layout.NewGridLayout(2), x0, y0, xEnd, yEnd)
	var ddaButon *widget.Button
	ddaButon = widget.NewButton("Reta DDA", func() {
		x, _ := strconv.ParseFloat(x0.Text, 64)
		y, _ := strconv.ParseFloat(y0.Text, 64)
		XE, _ := strconv.ParseFloat(xEnd.Text, 64)
		YE, _ := strconv.ParseFloat(yEnd.Text, 64)

		DDA(x, y, XE, YE, canvasContent, width, height)
	})

	aux1 := false
	var PM *widget.Button
	PM = widget.NewButton("Reta Ponto Médio", func() {
		if aux1 {
			PM.SetText("Reta Ponto Médio")
		} else {
			PM.SetText("Em desenvolvimento")
		}
		aux1 = !aux1
	})

	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	// Frame da esquerda
	return container.NewVBox(
		label,
		ddaButon,
		gridInput,
		PM,
		exitButton,
		layout.NewSpacer(), // Adiciona espaço flexível
	)
}

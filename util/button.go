package util

import (
	"ProjetoCG/err"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

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
	gridInputdda := container.New(layout.NewGridLayout(2), x0, y0, xEnd, yEnd)

	ddaButton := widget.NewButton("Reta DDA", func() {
		x, _ := err.ValidateInput(x0.Text, "X", app.Driver().AllWindows()[0])
		y, _ := err.ValidateInput(y0.Text, "Y", app.Driver().AllWindows()[0])
		XE, _ := err.ValidateInput(xEnd.Text, "XEnd", app.Driver().AllWindows()[0])
		YE, _ := err.ValidateInput(yEnd.Text, "YEnd", app.Driver().AllWindows()[0])

		// Chama o algoritmo DDA
		DDA(x, y, XE, YE, canvasContent, width, height)
	})

	x0pm := widget.NewEntry()
	x0pm.SetPlaceHolder("X")
	y0pm := widget.NewEntry()
	y0pm.SetPlaceHolder("Y")
	xEndpm := widget.NewEntry()
	xEndpm.SetPlaceHolder("X Final")
	yEndpm := widget.NewEntry()
	yEndpm.SetPlaceHolder("Y Final")

	gridInputpm := container.New(layout.NewGridLayout(2), x0pm, y0pm, xEndpm, yEndpm)
	pmButton := widget.NewButton("Reta Ponto Médio", func() {
		x, _ := err.ValidateInput(x0pm.Text, "X", app.Driver().AllWindows()[0])
		y, _ := err.ValidateInput(y0pm.Text, "Y", app.Driver().AllWindows()[0])
		XE, _ := err.ValidateInput(xEndpm.Text, "XEnd", app.Driver().AllWindows()[0])
		YE, _ := err.ValidateInput(yEndpm.Text, "YEnd", app.Driver().AllWindows()[0])

		// Chama o algoritmo do Ponto Médio
		PontoMedio(x, y, XE, YE, canvasContent, width, height)
	})

	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	// Frame da esquerda
	return container.NewVBox(
		label,
		ddaButton,
		gridInputdda,
		pmButton,
		gridInputpm,
		exitButton,
		layout.NewSpacer(),
	)
}

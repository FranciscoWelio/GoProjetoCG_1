package util

import (
	"ProjetoCG/dialog"
	"ProjetoCG/err"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	// "fyne.io/fyne/v2/dialog"
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
	myDialog := dialog.NewCustomDialog(app.NewWindow("Erro"), "Erro nos valores")
	gridInputdda := container.New(layout.NewGridLayout(2), x0, y0, xEnd, yEnd)
	var ddaButon *widget.Button
	ddaButon = widget.NewButton("Reta DDA", func() {
		x, erro := err.ValidateInput(x0.Text, "x")
		fmt.Println(erro)
		if erro != nil {
			myDialog.Show()
			return
		}
		y, erro := err.ValidateInput(y0.Text, "Y")
		if erro != nil {
			// dialog.ShowError(erro, app.NewWindow("Erro"))
			return
		}
		XE, erro := err.ValidateInput(xEnd.Text, "x")
		if erro != nil {
			// dialog.ShowError(erro, app.NewWindow("Erro"))
			return
		}
		YE, erro := err.ValidateInput(yEnd.Text, "x")
		if erro != nil {
			// dialog.ShowError(erro, app.NewWindow("Erro"))
			return
		}

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
	var PM *widget.Button
	PM = widget.NewButton("Reta Ponto Médio", func() {
		x, err := strconv.ParseFloat(x0pm.Text, 64)
		y, _ := strconv.ParseFloat(y0pm.Text, 64)
		XE, _ := strconv.ParseFloat(xEndpm.Text, 64)
		YE, _ := strconv.ParseFloat(yEndpm.Text, 64)
		if err != nil {
			// dialog.ShowError(err, app.NewWindow("Erro"))
			return
		}
		PontoMedio(x, y, XE, YE, canvasContent, width, height)
	})

	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	// Frame da esquerda
	return container.NewVBox(
		label,
		ddaButon,
		gridInputdda,
		PM,
		gridInputpm,
		exitButton,
		layout.NewSpacer(), // Adiciona espaço flexível
	)
}

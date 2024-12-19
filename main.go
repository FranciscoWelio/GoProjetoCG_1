package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Computação Gráfica")

	windowWidth := 1920
	windowHeight := 1080
	canvaWidth := 1500
	canvaHeight := 1000
	myWindow.Resize(fyne.Size{Width: float32(windowWidth), Height: float32(windowHeight)})

	//Botões
	label := widget.NewLabel("Controles:")
	aux := false
	var DDA *widget.Button
	DDA = widget.NewButton("Reta DDA", func() {
		if aux {
			DDA.SetText("Reta Ponto Médio")

		} else {
			DDA.SetText("Em desenvolvimento")

		}
		aux = !aux
	})

	var exitButton *widget.Button
	exitButton = widget.NewButton("Sair", func() {
		myApp.Quit()
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

	//Frame a esquerda
	leftFrame := container.NewVBox(
		label,
		DDA,
		PM,
		exitButton,
		layout.NewSpacer(), // Adiciona espaço flexível
	)

	//Config do Canvas
	backgroundColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	canvasRect := canvas.NewRectangle(backgroundColor)
	canvasRect.SetMinSize(fyne.Size{Width: float32(canvaWidth), Height: float32(canvaHeight)})

	rightFrame := container.NewWithoutLayout(canvasRect)

	canvasRect.Resize(fyne.NewSize(float32(canvaWidth), float32(canvaHeight)))
	canvasRect.Move(fyne.NewPos(0, 0))

	//Definindo os objetos principais do frame
	mainContent := container.NewHSplit(
		leftFrame,
		rightFrame,
	)
	mainContent.SetOffset(0.25)

	//Setando os objetos na janela e mostrando
	myWindow.SetContent(mainContent)
	myWindow.ShowAndRun()
}

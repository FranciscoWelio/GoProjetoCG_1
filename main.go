package main

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("Computação Gráfica")

	windowWidth := 1920
	windowHeight := 1080
	canvaWidth := 1500
	canvaHeight := 1000
	myWindow.Resize(fyne.Size{Width: float32(windowWidth), Height: float32(windowHeight)})

	// Frame esquerdo (botões)
	leftFrame := CreateButtons(myApp)

	// Config do Canvas
	backgroundColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	canvasRect := canvas.NewRectangle(backgroundColor)
	canvasRect.SetMinSize(fyne.Size{Width: float32(canvaWidth), Height: float32(canvaHeight)})

	rightFrame := container.NewWithoutLayout(canvasRect)
	canvasRect.Resize(fyne.NewSize(float32(canvaWidth), float32(canvaHeight)))
	canvasRect.Move(fyne.NewPos(0, 0))

	// Definindo os objetos principais do frame
	mainContent := container.NewHSplit(
		leftFrame,
		rightFrame,
	)
	mainContent.SetOffset(0.25)

	// Setando os objetos na janela e mostrando
	myWindow.SetContent(mainContent)
	myWindow.ShowAndRun()
}

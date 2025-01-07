package main

import (
	"ProjetoCG/util"
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
	// centroX := canvaWidth / 2
	// centroY := canvaHeight / 2

	myWindow.Resize(fyne.Size{Width: float32(windowWidth), Height: float32(windowHeight)})

	// Frame esquerdo (botões)
	canvaConteiner := container.NewWithoutLayout()
	leftFrame := util.CreateButtons(myApp, canvaConteiner, canvaWidth, canvaHeight)

	// Config do Canvas
	backgroundColor := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	canvasRect := canvas.NewRectangle(backgroundColor)
	canvasRect.SetMinSize(fyne.Size{Width: float32(canvaWidth), Height: float32(canvaHeight)})

	retaVertical := canvas.NewLine(color.Black)
	retaVertical.StrokeWidth = 1
	retaVertical.Position1 = fyne.NewPos(float32(canvaWidth/2)+0.5, 0)
	retaVertical.Position2 = fyne.NewPos(float32(canvaWidth/2)+0.5, float32(canvaHeight))

	retaHorizontal := canvas.NewLine(color.Black)
	retaHorizontal.StrokeWidth = 1
	retaHorizontal.Position1 = fyne.NewPos(0, float32(canvaHeight/2)+0.5)
	retaHorizontal.Position2 = fyne.NewPos(float32(canvaWidth), float32(canvaHeight/2)+0.5)

	// reta := util.DDA(float64(centroX+2), float64(centroY+2), 1500, 500)
	// retaFrame := container.NewWithoutLayout()
	// for _, obj := range reta {
	// 	retaFrame.Add(obj)
	// }

	rightFrame := container.NewWithoutLayout(
		canvasRect,
		retaHorizontal,
		retaVertical,
		canvaConteiner,
	)

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

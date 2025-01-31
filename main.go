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

	myWindow.Resize(fyne.Size{Width: float32(windowWidth), Height: float32(windowHeight)})

	// Frame esquerdo (botões)
	canvaConteiner := container.NewWithoutLayout()
	buttonContainer := container.NewVBox()
	// leftFrame := util.CreateButtons(myApp, canvaConteiner, canvaWidth, canvaHeight)

	updateButton := func(tipo string) {
		buttonContainer.Objects = nil
		if tipo == "2D" {
			buttonContainer.Add(util.CreateButtons(myApp, canvaConteiner, canvaWidth, canvaHeight))
		} else if tipo == "Matriz" {
			buttonContainer.Add(util.FormasMatriz(myApp, canvaConteiner, canvaWidth, canvaHeight))
		} else if tipo == "Matriz 3D" {
			buttonContainer.Add(util.FormasMatriz3D(myApp, canvaConteiner, canvaWidth, canvaHeight))
		}
		buttonContainer.Refresh()
	}

	updateButton("2D")

	menuBar := util.MakeMenu(myApp, myWindow, func(tipo string) {
		updateButton(tipo)
	})
	myWindow.SetMainMenu(menuBar)

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
		buttonContainer,
		rightFrame,
	)
	mainContent.SetOffset(0.25)

	// Setando os objetos na janela e mostrando
	myWindow.SetContent(mainContent)
	myWindow.ShowAndRun()
}

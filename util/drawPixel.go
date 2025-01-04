package util

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func DrawPixel(x int, y int) fyne.CanvasObject {
	ponto := canvas.NewCircle(color.RGBA{R: 255, G: 0, B: 0, A: 255})
	ponto.StrokeWidth = 0
	ponto.Resize(fyne.NewSize(5, 5))                        // Define o tamanho do "ponto"
	ponto.Move(fyne.NewPos(float32(x)-2.5, float32(y)-2.5)) // Centraliza o círculo na posição
	return ponto
}

package util

import (
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func DDA(x0, y0, xEnd, yEnd float64, canvasContent *fyne.Container, width, heigth int) {
	dx := xEnd - x0
	dy := yEnd - y0
	passo := math.Max(math.Abs(dx), math.Abs(dy))
	xIncremento := dx / passo
	yIncremento := dy / passo
	centerX := float64(width / 2)
	centerY := float64(heigth / 2)
	x := x0
	y := y0
	var ponto []fyne.CanvasObject
	pixel := canvas.NewRectangle(color.Black)
	pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
	pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
	ponto = append(ponto, pixel)
	for i := 0; i < int(passo); i++ {
		x += xIncremento
		y += yIncremento
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
	}

	canvasContent.Refresh()
}

package util

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func DDA(x0, y0, xEnd, yEnd float64) []fyne.CanvasObject {
	dx := xEnd - x0
	dy := yEnd - y0
	passo := math.Max(math.Abs(dx), math.Abs(dy))
	xIncremento := dx / passo
	yIncremento := dy / passo
	x := x0
	y := y0
	var ponto []fyne.CanvasObject
	pixel := canvas.NewRectangle(color.Black)
	pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
	pixel.Move(fyne.NewPos(float32(x), float32(y)))
	ponto = append(ponto, pixel)
	for i := 0; i < int(passo); i++ {
		x += xIncremento
		y += yIncremento
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(x), float32(y)))
		ponto = append(ponto, pixel)
	}
	fmt.Println(ponto[1])
	return ponto
}

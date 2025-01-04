package util

import (
	"fmt"
	"math"

	"fyne.io/fyne/v2"
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
	ponto = append(ponto, DrawPixel(int(x), int(-y)))
	fmt.Println(int(x), int(-y))
	for i := 0; i < int(passo); i++ {
		x += xIncremento
		y += yIncremento
		fmt.Println(int(x), int(-y))
		ponto = append(ponto, DrawPixel(int(x), int(-y)))
	}
	return ponto
}

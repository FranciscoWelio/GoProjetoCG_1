package util

import (
	"fmt"
	"image/color"
	"math"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func PontoMedio(x0, y0, xEnd, yEnd float64, canvasContent *fyne.Container, width, heigth int) {
	dx := math.Abs(xEnd - x0)
	dy := math.Abs(yEnd - y0)
	centerX := float64(width / 2)
	centerY := float64(heigth / 2)
	if dx >= dy && x0 <= xEnd && y0 <= yEnd {
		fmt.Println("1º Oitante")
		ds := 2*dy - dx
		incE := 2 * dy
		incNE := 2 * (dy - dx)
		x, y := float64(0), float64(0)
		if x0 > xEnd {
			x, y = xEnd, yEnd
			xEnd = x0
		} else {
			x, y = x0, y0
		}
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for x < xEnd {
			x += 1
			if ds < 0 {
				ds += incE
			} else {
				y += 1
				ds += incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
		canvasContent.Refresh()
	} else if dy >= dx && x0 <= xEnd && y0 <= yEnd {
		fmt.Println("2º Oitante")
		ds := 2*dx - dy
		incE := 2 * dx
		incNE := 2 * (dx - dy)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for y < yEnd {
			y += 1
			if ds < 0 {
				ds += incE
			} else {
				x += 1
				ds += incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dy >= dx && x0 >= xEnd && y0 <= yEnd {
		fmt.Println("3º Oitante")
		ds := 2*dx - dy
		incE := 2 * dx
		incNE := 2 * (dx - dy)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for x > xEnd {
			y += 1
			if ds < 0 {
				ds += incE
			} else {
				x -= 1
				ds += incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dx >= dy && x0 >= xEnd && y0 <= yEnd {
		fmt.Println("4º Oitante")
		ds := 2*dy - dx
		incE := 2 * dy
		incNE := 2 * (dy - dx)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for x > xEnd {
			x -= 1
			if ds < 0 {
				ds += incE
			} else {
				y += 1
				ds += incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dx >= dy && x0 >= xEnd && y0 >= yEnd {
		fmt.Println("5º Oitante")
		ds := 2*-dy + dx
		incE := 2 * -dy
		incNE := 2 * (-dy + dx)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for x > xEnd {
			x -= 1
			if ds < 0 {
				ds += -incE
			} else {
				y -= 1
				ds += -incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dy >= dx && x0 >= xEnd && y0 >= yEnd {
		fmt.Println("6º Oitante")
		dx, dy := dy, dx
		ds := 2*-dy + dx
		incE := 2 * -dy
		incNE := 2 * (-dy + dx)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for y > yEnd {
			y -= 1
			if ds < 0 {
				ds += -incE
			} else {
				x -= 1
				ds += -incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dy >= dx && xEnd >= x0 && y0 >= yEnd {
		fmt.Println("7º Oitante")
		dx, dy := dy, dx
		ds := 2*-dy + dx
		incE := 2 * -dy
		incNE := 2 * (-dy + dx)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for y > yEnd {
			y -= 1
			if ds < 0 {
				ds += -incE
			} else {
				x += 1
				ds += -incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	} else if dx >= dy && x0 <= xEnd && y0 >= yEnd {
		fmt.Println("8º Oitante")
		ds := 2*-dy + dx
		incE := 2 * -dy
		incNE := 2 * (-dy + dx)
		x, y := x0, y0
		pixel := canvas.NewRectangle(color.Black)
		pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
		pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
		canvasContent.Add(pixel)
		for x < xEnd {
			x += 1
			if ds < 0 {
				ds += -incE
			} else {
				y -= 1
				ds += -incNE
			}
			pixel := canvas.NewRectangle(color.Black)
			pixel.Resize(fyne.NewSize(1, 1)) // Tamanho do pixel
			pixel.Move(fyne.NewPos(float32(centerX+x), float32(centerY-y)))
			canvasContent.Add(pixel)
		}
	}

}

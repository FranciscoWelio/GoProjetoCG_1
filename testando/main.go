package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	// Criar um slice de objetos
	objects := []fyne.CanvasObject{
		canvas.NewRectangle(nil),
		canvas.NewText("Texto", nil),
		canvas.NewLine(nil),
	}

	// Imprimir o tipo de cada objeto
	for i, obj := range objects {
		fmt.Printf("Índice %d: Tipo %T\n", i, obj)
	}
}

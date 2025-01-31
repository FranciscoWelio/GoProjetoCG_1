package util

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func MakeMenu(a fyne.App, w fyne.Window, updateButton func(string)) *fyne.MainMenu {
	primitive := fyne.NewMenu("Primitivas",
		fyne.NewMenuItem("Primitivas 2D", func() {
			fmt.Println("Primitivas 2D Aqui")
			updateButton("2D")
		}),
		fyne.NewMenuItem("Primitivas 3D", func() {
			fmt.Println("Primitivas 3D Aqui")
			updateButton("3D")
		}),
	)

	matriz := fyne.NewMenu("Matriz",
		fyne.NewMenuItem("matriz 2D", func() {
			fmt.Println("matriz 2D Aqui")
			updateButton("Matriz")
		}),
		fyne.NewMenuItem("matriz 3D", func() {
			fmt.Println("matriz 3D Aqui")
			updateButton("Matriz 3D")
		}),
	)

	return fyne.NewMainMenu(primitive, matriz)
}

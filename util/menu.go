package util

import (
	"fmt"

	"fyne.io/fyne/v2"
)

func MakeMenu(a fyne.App, w fyne.Window) *fyne.MainMenu {
	primitive := fyne.NewMenu("Primitivas",
		fyne.NewMenuItem("Primitivas 2D", func() {
			fmt.Println("Primitivas 2D Aqui")

		}),
		fyne.NewMenuItem("Primitivas 3D", func() {
			fmt.Println("Primitivas 3D Aqui")
		}),
		fyne.NewMenuItem("Primitivas 4D", func() {
			fmt.Println("Primitivas 4D Aqui")
		}),
	)

	matriz := fyne.NewMenu("Matriz",
		fyne.NewMenuItem("matriz 2D", func() {
			fmt.Println("matriz 2D Aqui")
		}),
		fyne.NewMenuItem("matriz 3D", func() {
			fmt.Println("matriz 3D Aqui")
		}),
		fyne.NewMenuItem("matriz 4D", func() {
			fmt.Println("matriz 4D Aqui")
		}),
	)

	return fyne.NewMainMenu(primitive, matriz)
}

package util

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// CreateButtons cria e retorna o frame com os botões
func CreateButtons(app fyne.App) fyne.CanvasObject {
	label := widget.NewLabel("Controles:")
	aux := false
	var DDA *widget.Button
	DDA = widget.NewButton("Reta DDA", func() {
		if aux {
			DDA.SetText("Reta Ponto Médio")
		} else {
			DDA.SetText("Em desenvolvimento")
		}
		aux = !aux
	})

	aux1 := false
	var PM *widget.Button
	PM = widget.NewButton("Reta Ponto Médio", func() {
		if aux1 {
			PM.SetText("Reta Ponto Médio")
		} else {
			PM.SetText("Em desenvolvimento")
		}
		aux1 = !aux1
	})

	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	// Frame da esquerda
	return container.NewVBox(
		label,
		DDA,
		PM,
		exitButton,
		layout.NewSpacer(), // Adiciona espaço flexível
	)
}

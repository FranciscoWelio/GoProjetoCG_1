package util

import (
	"ProjetoCG/err"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"

	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func CreateButtons(app fyne.App, canvasContent *fyne.Container, width, height int) fyne.CanvasObject {
	x0 := widget.NewEntry()
	x0.SetPlaceHolder("X")
	y0 := widget.NewEntry()
	y0.SetPlaceHolder("Y")
	xEnd := widget.NewEntry()
	xEnd.SetPlaceHolder("X Final")
	yEnd := widget.NewEntry()
	yEnd.SetPlaceHolder("Y Final")
	label := widget.NewLabel("Controles:")
	gridInputdda := container.New(layout.NewGridLayout(2), x0, y0, xEnd, yEnd)

	ddaButton := widget.NewButton("Reta DDA", func() {
		x, _ := err.ValidateInput(&x0.Text, "X", app.Driver().AllWindows()[0])
		y, _ := err.ValidateInput(&y0.Text, "Y", app.Driver().AllWindows()[0])
		XE, _ := err.ValidateInput(&xEnd.Text, "XEnd", app.Driver().AllWindows()[0])
		YE, _ := err.ValidateInput(&yEnd.Text, "YEnd", app.Driver().AllWindows()[0])

		// Chama o algoritmo DDA
		DDA(x, y, XE, YE, canvasContent, width, height)
	})

	x0pm := widget.NewEntry()
	x0pm.SetPlaceHolder("X")
	y0pm := widget.NewEntry()
	y0pm.SetPlaceHolder("Y")
	xEndpm := widget.NewEntry()
	xEndpm.SetPlaceHolder("X Final")
	yEndpm := widget.NewEntry()
	yEndpm.SetPlaceHolder("Y Final")

	gridInputpm := container.New(layout.NewGridLayout(2), x0pm, y0pm, xEndpm, yEndpm)
	pmButton := widget.NewButton("Reta Ponto Médio", func() {
		x, _ := err.ValidateInput(&x0pm.Text, "X", app.Driver().AllWindows()[0])
		y, _ := err.ValidateInput(&y0pm.Text, "Y", app.Driver().AllWindows()[0])
		XE, _ := err.ValidateInput(&xEndpm.Text, "XEnd", app.Driver().AllWindows()[0])
		YE, _ := err.ValidateInput(&yEndpm.Text, "YEnd", app.Driver().AllWindows()[0])

		// Chama o algoritmo do Ponto Médio
		PontoMedio(x, y, XE, YE, canvasContent, width, height)
	})

	cleanButton := widget.NewButton("Limpar Tela", func() {
		canvasContent.Objects = nil
		canvasContent.Refresh()
	})

	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	// Frame da esquerda
	return container.NewVBox(
		label,
		ddaButton,
		gridInputdda,
		pmButton,
		gridInputpm,
		cleanButton,
		exitButton,
		layout.NewSpacer(),
	)
}

func Buttons2D(app fyne.App, canvasContent *fyne.Container, width, height int) fyne.CanvasObject {

	x0 := widget.NewEntry()
	x0.SetPlaceHolder("X")
	y0 := widget.NewEntry()
	y0.SetPlaceHolder("Y")
	xEnd := widget.NewEntry()
	xEnd.SetPlaceHolder("X Final")
	yEnd := widget.NewEntry()
	yEnd.SetPlaceHolder("Y Final")
	label := widget.NewLabel("Controles 2D:")
	gridInputdda := container.New(layout.NewGridLayout(2), x0, y0, xEnd, yEnd)
	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})

	return container.NewVBox(
		label,
		gridInputdda,
		exitButton,
	)
}

func FormasMatriz(app fyne.App, canvasContent *fyne.Container, width, height int) fyne.CanvasObject {
	tx := widget.NewEntry()
	tx.SetPlaceHolder("Translado em X")
	ty := widget.NewEntry()
	ty.SetPlaceHolder("Translado em Y")
	sx := widget.NewEntry()
	sx.SetPlaceHolder("Escala em X")
	sy := widget.NewEntry()
	sy.SetPlaceHolder("Escala em Y")
	rx := widget.NewEntry()
	rx.SetPlaceHolder("Rotação em X")
	ry := widget.NewEntry()
	ry.SetPlaceHolder("Rotação em Y")
	scx := widget.NewEntry()
	scx.SetPlaceHolder("Cisalhar em X")
	scy := widget.NewEntry()
	scy.SetPlaceHolder("Cisalhar em Y")

	label := widget.NewLabel("Composição 2D:")
	gridInputada := container.New(layout.NewGridLayout(2), tx, ty, sx, sy, rx, ry, scx, scy)
	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})
	return container.NewVBox(
		label,
		gridInputada,
		exitButton,
	)
}

func FormasMatriz3D(app fyne.App, canvasContent *fyne.Container, width, height int) fyne.CanvasObject {
	tx := widget.NewEntry()
	tx.SetPlaceHolder("Translado em X")
	ty := widget.NewEntry()
	ty.SetPlaceHolder("Translado em Y")
	tz := widget.NewEntry()
	tz.SetPlaceHolder("Translado em Z")
	sx := widget.NewEntry()
	sx.SetPlaceHolder("Escala em X")
	sy := widget.NewEntry()
	sy.SetPlaceHolder("Escala em Y")
	sz := widget.NewEntry()
	sz.SetPlaceHolder("Escala em Z")
	rxy := widget.NewEntry()
	rxy.SetPlaceHolder("Rotação em XY")
	ryz := widget.NewEntry()
	ryz.SetPlaceHolder("Rotação em YZ")
	rzx := widget.NewEntry()
	rzx.SetPlaceHolder("Rotação em ZX")
	scx := widget.NewEntry()
	scx.SetPlaceHolder("Cisalhar em X")
	scy := widget.NewEntry()
	scy.SetPlaceHolder("Cisalhar em Y")
	scz := widget.NewEntry()
	scz.SetPlaceHolder("Cisalhar em Z")

	label := widget.NewLabel("Composição 2D:")

	aplicarT := widget.NewButton("Aplicar Translação", func() {
		fmt.Println("Transladrou")
	})
	aplicarS := widget.NewButton("Aplicar Escala", func() {
		fmt.Println("Escalou")
	})
	aplicarR := widget.NewButton("Aplicar Rotação", func() {
		fmt.Println("Rotacionou")
	})
	aplicarRefXY := widget.NewButton("Reflexão XY", func() {
		fmt.Println("Refletiu")
	})
	aplicarRefYZ := widget.NewButton("Reflexão YZ", func() {
		fmt.Println("Refletiu")
	})
	aplicarRefZX := widget.NewButton("Reflexão ZX", func() {
		fmt.Println("Refletiu")
	})
	aplicarSci := widget.NewButton("Aplicar Cisalhamento", func() {
		fmt.Println("Cisalhou")
	})
	exitButton := widget.NewButton("Sair", func() {
		app.Quit()
	})
	gridInputadaT := container.New(layout.NewGridLayout(3), tx, ty, tz)
	gridInputadaS := container.New(layout.NewGridLayout(3), sx, sy, sz)
	gridInputadaR := container.New(layout.NewGridLayout(3), rxy, ryz, rzx)
	gridInputadaSc := container.New(layout.NewGridLayout(3), scx, scy, scz)
	gridButtonsRef := container.New(layout.NewGridLayout(3), aplicarRefXY, aplicarRefYZ, aplicarRefZX)
	return container.NewVBox(
		label,
		gridInputadaT,
		aplicarT,
		gridInputadaS,
		aplicarS,
		gridInputadaR,
		aplicarR,
		gridButtonsRef,
		gridInputadaSc,
		aplicarSci,
		exitButton,
	)
}

package err

import (
	"ProjetoCG/dialog"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
)

func ValidateInput(variavel *string, nome string, myJanela fyne.Window) (float64, error) {
	if *variavel == "" {
		myDialog := dialog.NewCustomDialog(myJanela, "Erro nos valores "+nome+"\nO valores serão transformados em \"0\"")
		myDialog.Show()
		*variavel = "0"
		return 0, fmt.Errorf("o campo %s está vazio", nome)
	}
	vari, err := strconv.ParseFloat(*variavel, 64)
	if err != nil {
		return 0, fmt.Errorf("O campo %s é inválido: %v", nome, err)
	}
	return vari, nil
}

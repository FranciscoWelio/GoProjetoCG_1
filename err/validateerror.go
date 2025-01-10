package err

import (
	"fmt"
	"strconv"
)

func ValidateInput(variavel string, nome string) (float64, error) {
	if variavel == "" {
		return 0, fmt.Errorf("O campo %s não esta preenchido", nome)
	}
	vari, err := strconv.ParseFloat(variavel, 64)
	if err != nil {
		return 0, fmt.Errorf("O campo %s é inválido: %v", nome, err)
	}
	return vari, nil
}

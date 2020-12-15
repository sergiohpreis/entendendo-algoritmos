package main

import (
	"fmt"
	"log"
)

func pesquisaBinaria(lista [5]int, item int) (found int, err error) {
	var baixo = 0
	var alto = len(lista) - 1
	fmt.Printf("Buscando item: %d \n", item)
	for baixo <= alto {
		meio := (baixo + alto) / 2
		chute := lista[meio]
		fmt.Printf("Baixo: %d, Alto: %d, Meio: %d, Chute: %d \n", baixo, alto, meio, chute)
		if chute == item {
			return meio, nil
		}

		if chute > item {
			alto = meio - 1
		} else {
			baixo = meio + 1
		}
	}
	return 0, fmt.Errorf("Item %d não encontrado", item)
}

func main() {
	var lista = [5]int{1, 3, 5, 7, 9}
	p1, err := pesquisaBinaria(lista, 3)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("O item %d foi encontrado \n", p1)
	p2, err := pesquisaBinaria(lista, -1)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("O item %d foi encontrado \n", p2)
}

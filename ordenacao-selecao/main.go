package main

import "fmt"

func buscaMenor(arr []int) int {
	fmt.Printf("...Executando função buscaMenor(%v)\n", arr)
	menor := arr[0]
	fmt.Printf("Inicializando \"menor\" = %d\n", menor)
	menorIndice := 0
	fmt.Printf("Inicializando \"menorIndice\" = %d\n", menorIndice)
	for i := range arr {
		fmt.Printf("Elemento atual = %d\n", arr[i])
		if arr[i] < menor {
			fmt.Println("Elemento atual < menor")
			menor = arr[i]
			fmt.Printf("O menor elemento atualmente é o %d\n", menor)
			menorIndice = i
			fmt.Printf("O indice do menor elemento atualmente é o %d\n", menorIndice)
		}
	}
	fmt.Printf("- O indice do menor elemento é o %d\n", menorIndice)
	return menorIndice
}

func ordernacaoPorSelecao(arr []int) []int {
	var novoArr []int

	for range arr {
		menor := buscaMenor(arr)
		novoArr = append(novoArr, arr[menor])
		arr = append(arr[:menor], arr[menor+1:]...)
	}

	return novoArr
}

func main() {
	lista := []int{5, 3, 6, 2, 10}
	ordernado := ordernacaoPorSelecao(lista)
	fmt.Println("RESULTADO: Valores ordenados =", ordernado)
}

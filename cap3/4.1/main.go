package main

import "fmt"

func soma(arr []int) int {
	total := 0
	for _, numero := range arr {
		fmt.Printf("Total: %d, Número: %d\n", total, numero)
		total += numero
	}
	return total
}

func somaRecursiva(arr []int) int {
	tamanhoArray := len(arr)

	if tamanhoArray == 0 {
		return 0
	}

	primeiro := arr[0]
	resto := arr[1:]
	soma := somaRecursiva(resto)

	fmt.Printf("Primeiro Elemento: %d, Resto do Array: %d, Soma: %d\n", primeiro, resto, soma)

	return primeiro + soma
}

func main() {
	fmt.Println("Soma:", soma([]int{1, 2, 3, 4, 5, 6}))
	fmt.Println("Soma recursiva:", somaRecursiva([]int{1, 2, 3, 4, 5, 6}))
}

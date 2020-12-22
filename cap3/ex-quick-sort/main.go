package main

import (
	"fmt"
	"time"
)

func quickSort1(arr []int) []int {
	fmt.Println("1 - Lista:", arr)
	if len(arr) < 2 {
		return arr
	}

	pivo := arr[0]
	var menores []int
	var maiores []int
	for _, element := range arr[1:] {
		if element < pivo {
			menores = append(menores, element)
		} else {
			maiores = append(maiores, element)
		}
	}

	fmt.Printf("1 - Pivo: %d, Menores: %d, Maiores: %d\n", pivo, menores, maiores)

	return append(append(quickSort1(menores), pivo), quickSort1(maiores)...)
}

func quickSort2(arr []int) []int {
	fmt.Println("2 - Lista:", arr)
	if len(arr) < 2 {
		return arr
	}

	pivo := arr[len(arr)/2]
	resto := append(arr[:len(arr)/2], arr[1+len(arr)/2:]...)
	var menores []int
	var maiores []int
	for _, element := range resto {
		if element < pivo {
			menores = append(menores, element)
		} else {
			maiores = append(maiores, element)
		}
	}

	fmt.Printf("2- Pivo: %d, Menores: %d, Maiores: %d\n", pivo, menores, maiores)

	return append(append(quickSort2(menores), pivo), quickSort2(maiores)...)
}

type quicksort func([]int) []int

func executeQuickSortWithInfo(f quicksort, list []int) ([]int, time.Duration) {
	start := time.Now()
	elapsed := time.Since(start)
	sorted := f(list)

	return sorted, elapsed
}

func main() {
	lista := []int{10, 5, 2, 3, 15, 21, 69, 13}

	sorted1, time1 := executeQuickSortWithInfo(quickSort1, lista)
	fmt.Printf("O algoritmo retornou %d e levou %s para ser executado\n", sorted1, time1)

	sorted2, time2 := executeQuickSortWithInfo(quickSort2, lista)
	fmt.Printf("O algoritmo retornou %d e levou %s para ser executado\n", sorted2, time2)
}

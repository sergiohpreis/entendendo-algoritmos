package main

import "fmt"

func contador(arr []int) int {
	if len(arr) == 0 {
		return 0
	}

	contagem := contador(arr[1:])
	fmt.Println("Contagem", contagem)

	return 1 + contagem
}

func main() {
	fmt.Println(contador([]int{1, 2, 3, 4, 5, 6, 6, 6, 6, 6, 6, 6, 6}))
}

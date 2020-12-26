package main

import "fmt"

var count = 0

// Minha solução
func encontraMaisAlto(arr []int) int {
	count++
	if len(arr) == 1 {
		return arr[0]
	}

	if arr[0] > arr[1] {
		return encontraMaisAlto(append([]int{arr[0]}, arr[2:]...))
	}

	return encontraMaisAlto(append([]int{arr[1]}, arr[2:]...))
}

var countLivro = 0

// Solução do livro
func encontraMaisAltoLivro(arr []int) int {
	countLivro++
	if len(arr) == 2 {
		if arr[0] > arr[1] {
			return arr[0]
		}

		return arr[1]
	}

	subMax := encontraMaisAltoLivro(arr[1:])

	if arr[0] > subMax {
		return arr[0]
	}

	return subMax
}

func main() {
	fmt.Printf("Minha solução retornou %d e fez %d operações\n", encontraMaisAlto([]int{1, 1000, 3}), count)
	fmt.Printf("A solução do livro retornou %d e fez %d operações\n", encontraMaisAltoLivro([]int{1, 1000, 3}), countLivro)
}

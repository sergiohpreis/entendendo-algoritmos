package main

import "fmt"

func vendaManga(nome string) bool {
	// verifica a ultima letra do nome
	return string(nome[len(nome)-1]) == "m"
}

func estaNaLista(lista []string, nome string) bool {
	verificada := false
	for i := 0; i < len(lista); i++ {
		if lista[i] == nome {
			verificada = true
		}
	}

	return verificada
}

func pesquisa(grafo map[string][]string, nome string) bool {
	var fila []string
	var verificadas []string
	fila = append(fila, grafo[nome]...)

	for len(fila) > 0 {
		// pega o primeiro da fila
		pessoa := fila[0]
		// fila sem a primeiro
		fila = fila[1:]

		if estaNaLista(fila, pessoa) == false {
			if vendaManga(pessoa) {
				fmt.Printf("%s é um vendedor de manga!\n", pessoa)
				return true
			}
		}

		fila = append(fila, grafo[pessoa]...)
		verificadas = append(verificadas, pessoa)
	}

	return false
}

func main() {
	var grafo = make(map[string][]string)

	grafo["voce"] = []string{"alice", "bob", "claire"}
	grafo["bob"] = []string{"anuj", "peggy"}
	grafo["alice"] = []string{"peggy"}
	grafo["claire"] = []string{"thom", "jonny"}
	grafo["anuj"] = []string{}
	grafo["peggy"] = []string{}
	grafo["thom"] = []string{}
	grafo["jonny"] = []string{}

	pesquisa(grafo, "voce")
}

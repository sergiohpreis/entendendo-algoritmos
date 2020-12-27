package main

import (
	"fmt"
	"math"
)

func getKeysInMap(m map[string]int) []string {
	keys := []string{}

	for key := range m {
		keys = append(keys, key)
	}

	return keys
}

func isNodeProcessed(node string, processed []string) bool {
	isProcessed := false

	for i := 0; i < len(processed); i++ {
		if node == processed[i] {
			isProcessed = true
		}
	}

	return isProcessed
}

func findCheaperNode(costs map[string]float64, processed []string) string {
	cheaper := math.Inf(1)
	cheaperNode := ""

	for node := range costs {
		cost := costs[node]
		if cost < cheaper && !isNodeProcessed(node, processed) {
			cheaper = cost
			cheaperNode = node
		}
	}

	return cheaperNode
}

func dijkstra(graph map[string]map[string]int, costs map[string]float64, parents map[string]string) (map[string]float64, map[string]string) {
	var processed []string

	node := findCheaperNode(costs, processed)

	for node != "" {
		fmt.Println("---- node:", node)
		cost := costs[node]
		fmt.Println("------ cost:", cost)
		siblings := graph[node]
		fmt.Println("------ siblings:", siblings)

		for _, n := range getKeysInMap(siblings) {
			newCost := cost + float64(siblings[n])
			fmt.Printf("-------- from start to <%s> costs: %f \n", n, costs[n])
			fmt.Printf("-------- from start to here costs: %f, and from here to <%s> costs: %f \n", cost, n, float64(siblings[n]))
			if costs[n] > newCost {
				costs[n] = newCost
				parents[n] = node
			}
		}
		processed = append(processed, node)
		node = findCheaperNode(costs, processed)
	}

	return costs, parents
}

func graph1() {
	fmt.Println("////////////////////////////////////////////////////////////////////////")
	fmt.Println("- Grafo 1")
	graph := map[string]map[string]int{}
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 6
	graph["start"]["b"] = 2

	graph["a"] = map[string]int{}
	graph["a"]["end"] = 1

	graph["b"] = map[string]int{}
	graph["b"]["a"] = 3
	graph["b"]["end"] = 5

	graph["end"] = map[string]int{}

	costs := map[string]float64{}
	costs["a"] = 6.0
	costs["b"] = 2.0
	costs["end"] = math.Inf(1)

	parents := map[string]string{}
	parents["a"] = "start"
	parents["b"] = "start"
	parents["end"] = ""

	fmt.Println("-- Costs:", costs)
	fmt.Println("-- Parents:", parents)

	newCosts, newParents := dijkstra(graph, costs, parents)

	fmt.Println("-- after dijkstra")
	fmt.Println("-- Costs:", newCosts)
	fmt.Println("-- Parents:", newParents)
	fmt.Println("////////////////////////////////////////////////////////////////////////")
}

func graph2() {
	fmt.Println("////////////////////////////////////////////////////////////////////////")
	fmt.Println("- Grafo 2")
	graph := map[string]map[string]int{}
	graph["start"] = map[string]int{}
	graph["start"]["a"] = 2
	graph["start"]["b"] = 5

	graph["a"] = map[string]int{}
	graph["a"]["b"] = 8
	graph["a"]["d"] = 7

	graph["b"] = map[string]int{}
	graph["b"]["c"] = 4
	graph["b"]["d"] = 2

	graph["c"] = map[string]int{}
	graph["c"]["d"] = 6
	graph["c"]["end"] = 3

	graph["d"] = map[string]int{}
	graph["d"]["end"] = 1

	graph["end"] = map[string]int{}

	costs := map[string]float64{}
	costs["a"] = 2.0
	costs["b"] = 5.0
	costs["c"] = math.Inf(1)
	costs["d"] = math.Inf(1)
	costs["end"] = math.Inf(1)

	parents := map[string]string{}
	parents["a"] = "start"
	parents["b"] = "start"
	parents["c"] = ""
	parents["d"] = ""
	parents["end"] = ""

	fmt.Println("-- Costs:", costs)
	fmt.Println("-- Parents:", parents)

	newCosts, newParents := dijkstra(graph, costs, parents)
	fmt.Println("-- after dijkstra")
	fmt.Println("-- Costs:", newCosts)
	fmt.Println("-- Parents:", newParents)
	fmt.Println("////////////////////////////////////////////////////////////////////////")
}

func main() {
	graph1()
	graph2()
}

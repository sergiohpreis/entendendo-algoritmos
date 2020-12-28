package main

import "fmt"

func getLowAndHighSlice(a []string, b []string) (low []string, high []string) {
	low, high = a, b
	if len(a) > len(b) {
		low = b
		high = a
	}

	return
}

func intersection(a []string, b []string) (inter []string) {
	low, high := getLowAndHighSlice(a, b)

	for i := range low {
		for j := range high {
			if low[i] == high[j] {
				inter = append(inter, low[i])
			}
		}
	}

	return inter
}

func diference(a []string, b []string) (diff []string) {
	low, high := getLowAndHighSlice(a, b)
	valueInLower := map[string]bool{}

	for i := range low {
		valueInLower[low[i]] = true
	}

	for i := range high {
		if !valueInLower[high[i]] {
			diff = append(diff, high[i])
		}
	}

	return diff
}

func main() {
	statesToProcess := []string{"mt", "wa", "or", "id", "nv", "ut", "ca", "az"}
	stations := map[string][]string{}
	stations["k1"] = []string{"id", "nv", "ut"}
	stations["k2"] = []string{"wa", "id", "mt"}
	stations["k3"] = []string{"or", "nv", "ca"}
	stations["k4"] = []string{"nv", "ut"}
	stations["k5"] = []string{"ca", "az"}

	/*
		In go, range in a map is random, so this is to
		keep the range always the same stations order
	*/
	stationKeys := []string{"k1", "k2", "k3", "k4", "k5"}

	var finalStations []string

	for len(statesToProcess) > 0 {
		fmt.Println("- statesToProcess:", statesToProcess)
		var bestStation string
		var statesCovered []string

		for _, station := range stationKeys {
			states := stations[station]
			fmt.Printf("-- station: %s, cover the states: %s \n", station, states)
			covered := intersection(statesToProcess, states)
			fmt.Println("-- covered:", covered)

			if len(covered) > len(statesCovered) {
				bestStation = station
				statesCovered = covered
				fmt.Println("---- bestStation:", bestStation)
				fmt.Println("---- statesCovered:", statesCovered)
			}
		}

		statesToProcess = diference(statesToProcess, statesCovered)
		finalStations = append(finalStations, bestStation)
	}

	fmt.Println("------ finalStations:", finalStations)
}

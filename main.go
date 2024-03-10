package main

import (
	"fmt"
	"sort"
)

func maxChickensProtected(n, k int, positions []int) int {
	maxProtected := 0

	for i := 0; i < n; i++ {
		roofStart := positions[i]
		roofEnd := roofStart + k

		//find maximum chickens covered by the roof
		coveredChikens := sort.Search(n, func(j int) bool {
			return positions[j] > roofEnd
		}) - i

		if coveredChikens > maxProtected {
			maxProtected = coveredChikens
		}
	}

	return maxProtected
}

func main() {
	// input values
	n := 6
	k := 10
	positions := []int{1, 11, 30, 34, 35, 37}

	//sort positions
	sort.Ints(positions)

	//calculate
	result := maxChickensProtected(n, k, positions)
	fmt.Println(result)
}

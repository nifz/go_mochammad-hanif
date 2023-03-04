package main

import (
	"fmt"
	"sort"
)

type pair struct {
	name  string
	count int
}

func MostAppearItem(items []string) []pair {
	counter := make(map[string]int)
	for _, item := range items {
		counter[item]++
	}

	var pairs []pair
	for name, count := range counter {
		pairs = append(pairs, pair{name, count})
	}

	// Sorting pairs by count in ascending order
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count < pairs[j].count
	})

	return pairs
}

func main() {
	pairs := MostAppearItem([]string{"js", "js", "golang", "ruby", "ruby", "js", "js"})
	for _, p := range pairs {
		fmt.Printf("%s->%d ", p.name, p.count)
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"A", "B", "B", "C", "A", "A", "B", "A", "D", "D"})
	for _, p := range pairs {
		fmt.Printf("%s->%d ", p.name, p.count)
	}
	fmt.Println()

	pairs = MostAppearItem([]string{"football", "basketball", "tenis"})
	for _, p := range pairs {
		fmt.Printf("%s->%d ", p.name, p.count)
	}
	fmt.Println()
}

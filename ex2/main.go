package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {

	tags := os.Args[1:]

	if len(tags) == 0 {
		fmt.Fprintln(os.Stderr, "Erreur aucun tag fournie")
		os.Exit(1)
	}

	type tagCount struct {
		tag   string
		count int
	}

	counts := make(map[string]int)

	for _, tag := range tags {
		counts[tag]++
	}

	list := make([]tagCount, 0, len(counts))

	for tag, n := range counts {
		list = append(list, tagCount{tag: tag, count: n})
	}

	sort.Slice(list, func(i, j int) bool {
		return list[i].count < list[j].count
	})

	for _, tc := range list {
		if tc.count > 1 {
			fmt.Printf("%s : %d\n", tc.tag, tc.count)
		}

	}

}

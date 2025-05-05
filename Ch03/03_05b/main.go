package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["Al"] = "Montgomery"
	states["Alsk"] = "Juneau"
	states["Ar"] = "Little Rock"
	states["Ca"] = "California"
	fmt.Println(states)


	california := states["Ca"]
	fmt.Println(california)
	delete(states, "Ca")
	fmt.Println(states)

	states["NY"] = "NEW YORK"
	fmt.Println(states)

	for k, v := range states {
		fmt.Printf("%v : %v\n", k, v)
	}

	keys := make([] string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	fmt.Println("\nSorted keys:")

	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}

package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	b := *NewBoard(4)
	p := BoardProblem{b, Tile{0, 0}, Tile{3, 3}}
	f := Queue{}
	solution := UninformedSearch(p, &f)
	fmt.Println(solution)

	fmt.Println("---------------------")
	for solution != nil {
		tile := solution.State.(Tile)

		fmt.Println(tile)
		solution = solution.Parent
	}
	fmt.Println("--------------------")
}

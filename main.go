package main

import "fmt"

func main() {
	// fmt.Println("Hello World")
	b := *NewBoard(10)
	p := BoardProblem{b, Tile{0, 0}, Tile{9, 9}}
	pq := make(PriorityQueue, 0)

	solution := AstarSearch(p, pq)
	// f := &Queue{}
	// UninformedSearch(p, f)
	// fmt.Println(solution)

	fmt.Println("---------------------")
	fmt.Println(solution.PathCost)
	for solution != nil {
		tile := solution.State.(Tile)

		fmt.Println(tile)
		solution = solution.Parent
	}
	fmt.Println("--------------------")
}

/*

_____________
|  |  |  |  |
+--+--+--+--+
|  |  |  |  |
+--+--+--+--+
|  |  |  |  |
+--+--+--+--+
|  |  |  |  |
+--+--+--+--+






*/

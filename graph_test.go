package main

import (
	"reflect"
	"testing"
)

func TestBfs(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{0, 0},
		GoalTiles: []Tile{Tile{2, 2}},
	}
	f := Queue{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTiles[0] {
		t.Error(
			"expected to found solution",
		)
	}

	expectedCost := 4
	actualCost := solution.PathCost
	equal := reflect.DeepEqual(actualCost, expectedCost)
	if !equal {
		t.Error(
			"expected path cost", expectedCost,
			"got path cost", actualCost,
		)
	}

	for _, n := range b.Neighbors(Tile{2, 2}) {
		b.RemoveNeighbor(n, Tile{2, 2})
	}
	f = Queue{}
	solution = UninformedSearch(p, &f)

	if solution != nil {
		t.Error(
			"expected no solution",
		)
	}
}

func TestDfs(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{0, 0},
		GoalTiles: []Tile{Tile{2, 2}},
	}
	f := Stack{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTiles[0] {
		t.Error(
			"expected to found solution",
		)
	}

	minCost := 4
	maxCost := 24
	actualCost := solution.PathCost
	if !(actualCost >= minCost && actualCost <= maxCost) {
		t.Error(
			"expected path cost between", minCost, maxCost,
			"got path cost", actualCost,
		)
	}

	for _, n := range b.Neighbors(Tile{2, 2}) {
		b.RemoveNeighbor(n, Tile{2, 2})
	}
	f = Stack{}
	solution = UninformedSearch(p, &f)

	if solution != nil {
		t.Error(
			"expected no solution",
		)
	}
}

func TestUninformedTrivialSolution(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{2, 2},
		GoalTiles: []Tile{Tile{2, 2}},
	}
	f := Stack{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTiles[0] {
		t.Error(
			"expected to found solution",
		)
	}

	expectedCost := 0
	actualCost := solution.PathCost
	equal := reflect.DeepEqual(actualCost, expectedCost)
	if !equal {
		t.Error(
			"expected path cost", expectedCost,
			"got path cost", actualCost,
		)
	}
}

func TestAstar(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{0, 0},
		GoalTiles: []Tile{Tile{4, 4}},
	}
	pq := make(PriorityQueue, 0)
	solution := AstarSearch(p, pq)

	if solution.State != p.GoalTiles[0] {
		t.Error(
			"expected to found solution",
		)
	}

	expectedCost := 8
	actualCost := solution.PathCost
	equal := reflect.DeepEqual(actualCost, expectedCost)
	if !equal {
		t.Error(
			"expected path cost", expectedCost,
			"got path cost", actualCost,
		)
	}

	for _, n := range b.Neighbors(Tile{4, 4}) {
		b.RemoveNeighbor(n, Tile{4, 4})
	}
	pq = make(PriorityQueue, 0)
	solution = AstarSearch(p, pq)

	if solution != nil {
		t.Error(
			"expected no solution",
		)
	}
}

func TestAstarTrivialSolution(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{2, 2},
		GoalTiles: []Tile{Tile{2, 2}},
	}
	pq := make(PriorityQueue, 0)
	solution := AstarSearch(p, pq)

	if solution.State != p.GoalTiles[0] {
		t.Error(
			"expected to found solution",
		)
	}

	expectedCost := 0
	actualCost := solution.PathCost
	equal := reflect.DeepEqual(actualCost, expectedCost)
	if !equal {
		t.Error(
			"expected path cost", expectedCost,
			"got path cost", actualCost,
		)
	}
}

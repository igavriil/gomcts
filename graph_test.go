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
		GoalTile:  Tile{2, 2},
	}
	f := Queue{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTile {
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
		GoalTile:  Tile{2, 2},
	}
	f := Stack{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTile {
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

func TestTrivialSolution(t *testing.T) {
	b := *NewBoard(5)
	p := BoardProblem{
		Board:     b,
		StartTile: Tile{2, 2},
		GoalTile:  Tile{2, 2},
	}
	f := Stack{}
	solution := UninformedSearch(p, &f)

	if solution.State != p.GoalTile {
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

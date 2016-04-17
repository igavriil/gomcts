package main

import (
	"reflect"
	"testing"
)

func TestNewBoard(t *testing.T) {
	expected := &Board{
		Tile{0, 0}: map[Tile]int{Tile{0, 1}: 1, Tile{1, 0}: 1},
		Tile{0, 1}: map[Tile]int{Tile{0, 0}: 1, Tile{1, 1}: 1},
		Tile{1, 0}: map[Tile]int{Tile{0, 0}: 1, Tile{1, 1}: 1},
		Tile{1, 1}: map[Tile]int{Tile{0, 1}: 1, Tile{1, 0}: 1},
	}
	actual := NewBoard(2)
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestActions(t *testing.T) {
	b := *NewBoard(2)
	expected := []BoardAction{
		BoardAction{Tile{0, 0}, Tile{0, 1}},
		BoardAction{Tile{0, 0}, Tile{1, 0}},
	}
	actual := b.Actions(Tile{0, 0})
	equal := reflect.DeepEqual(
		sliceToBoardActionMap(actual),
		sliceToBoardActionMap(expected),
	)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func sliceToBoardActionMap(s []BoardAction) map[BoardAction]int {
	m := make(map[BoardAction]int)
	for _, v := range s {
		m[v] = 1
	}
	return m
}

func TestTransition(t *testing.T) {
	b := *NewBoard(2)
	expected := Tile{0, 1}
	actual, error := b.Transition(
		Tile{0, 0},
		BoardAction{Tile{0, 0}, Tile{0, 1}},
	)
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	if error != nil {
		t.Error(
			"expected no error",
			"got", error,
		)
	}

	_, error = b.Transition(
		Tile{1, 1},
		BoardAction{Tile{0, 0}, Tile{0, 1}},
	)
	if error == nil {
		t.Error(
			"expected error",
		)
	}
}

func TestStepCost(t *testing.T) {
	b := *NewBoard(2)
	expected := 1
	actual, error := b.StepCost(
		Tile{0, 0},
		BoardAction{Tile{0, 0}, Tile{0, 1}},
	)
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	if error != nil {
		t.Error(
			"expected no error",
			"got", error,
		)
	}

	_, error = b.StepCost(
		Tile{1, 1},
		BoardAction{Tile{0, 0}, Tile{0, 1}},
	)
	if error == nil {
		t.Error(
			"expected error",
		)
	}
}

func TestNeighbors(t *testing.T) {
	b := *NewBoard(2)
	expected := []Tile{Tile{0, 1}, Tile{1, 0}}
	actual := b.Neighbors(Tile{0, 0})
	equal := reflect.DeepEqual(
		sliceToTileMap(actual),
		sliceToTileMap(expected),
	)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func sliceToTileMap(s []Tile) map[Tile]int {
	m := make(map[Tile]int)
	for _, v := range s {
		m[v] = 1
	}
	return m
}

func TestAddNeighbor(t *testing.T) {
	b := *NewBoard(2)
	if _, ok := b[Tile{0, 0}][Tile{1, 1}]; ok {
		t.Error(
			"Expected no edge between 0,0 and 1,1",
			"got", b[Tile{0, 0}][Tile{1, 1}],
		)
	}
	if _, ok := b[Tile{1, 1}][Tile{0, 0}]; ok {
		t.Error(
			"Expected no edge between 1,1 and 0,0",
			"got", b[Tile{1, 1}][Tile{0, 0}],
		)
	}
	b.AddNeighbor(Tile{0, 0}, Tile{1, 1})

	if _, ok := b[Tile{0, 0}][Tile{1, 1}]; !ok {
		t.Error(
			"Expected edge between 0,0 and 1,1",
			"got", b[Tile{0, 0}][Tile{1, 1}],
		)
	}
	if _, ok := b[Tile{1, 1}][Tile{0, 0}]; !ok {
		t.Error(
			"Expected edge between 1,1 and 0,0",
			"got", b[Tile{1, 1}][Tile{0, 0}],
		)
	}
}

func TestRemoveNeighbor(t *testing.T) {
	b := *NewBoard(2)

	if _, ok := b[Tile{0, 0}][Tile{0, 1}]; !ok {
		t.Error(
			"Expected edge between 0,0 and 0,1",
			"got", b[Tile{0, 0}][Tile{0, 1}],
		)
	}
	if _, ok := b[Tile{0, 1}][Tile{0, 0}]; !ok {
		t.Error(
			"Expected edge between 0,1 and 0,0",
			"got", b[Tile{0, 1}][Tile{0, 0}],
		)
	}
	b.RemoveNeighbor(Tile{0, 0}, Tile{0, 1})

	if _, ok := b[Tile{0, 0}][Tile{0, 1}]; ok {
		t.Error(
			"Expected no edge between 0,0 and 0,1",
			"got", b[Tile{0, 0}][Tile{0, 1}],
		)
	}
	if _, ok := b[Tile{0, 1}][Tile{0, 0}]; ok {
		t.Error(
			"Expected no edge between 0,1 and 0,0",
			"got", b[Tile{0, 1}][Tile{0, 0}],
		)
	}
}

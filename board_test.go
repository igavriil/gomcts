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
	expected := []Action{
		BoardAction{Tile{0, 0}, Tile{0, 1}},
		BoardAction{Tile{0, 0}, Tile{1, 0}},
	}
	actual := b.Actions(Tile{0, 0})
	equal := reflect.DeepEqual(
		sliceToActionMap(actual),
		sliceToActionMap(expected),
	)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func sliceToActionMap(s []Action) map[Action]int {
	m := make(map[Action]int)
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

var distanceTest = []struct {
	tile     Tile
	goal     Tile
	distance int
}{
	{Tile{0, 0}, Tile{0, 0}, 0},
	{Tile{1, 1}, Tile{1, 2}, 1},
	{Tile{1, 1}, Tile{2, 1}, 1},
	{Tile{2, 3}, Tile{5, 4}, 4},
}

func TestManhattanDisact(t *testing.T) {
	for _, tt := range distanceTest {
		s := tt.tile.ManhattanDistance(tt.goal)
		if s != tt.distance {
			t.Errorf("%v.ManhattanDistance(%v)=>%v, want %v", tt.tile, tt.goal, s, tt.distance)
		}
	}
}

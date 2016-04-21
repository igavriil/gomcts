package main

import (
	"fmt"
	"math"
)

type Tile struct {
	i, j int
}

func (t Tile) ManhattanDistance(g Tile) int {
	return Abs(t.i-g.i) + Abs(t.j-g.j)
}

func (t Tile) DiagonalDistance(g Tile) int {
	return Min(Abs(t.i-g.i), Abs(t.j-g.j))
}

func (t Tile) CrossProductDistance(s Tile, g Tile) int {
	dx1 := t.i - g.i
	dy1 := t.j - g.j
	dx2 := s.i - g.i
	dy2 := s.j - g.j
	return Abs(dx1*dy2 - dx2*dy1)
}

type BoardAction struct {
	from Tile
	to   Tile
}

type Board map[Tile]map[Tile]int

type BoardProblem struct {
	Board
	StartTile Tile
	GoalTiles []Tile
}

func NewBoard(dimension int) *Board {
	b := make(Board)
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			b[Tile{i, j}] = make(map[Tile]int)
		}
	}
	b.InitializeNeighbors(dimension)

	return &b
}

func (b Board) InitializeNeighbors(dimension int) {
	for i := 0; i < dimension; i++ {
		for j := 0; j < dimension; j++ {
			if j+1 >= 0 && j+1 < dimension {
				b.AddNeighbor(Tile{i, j}, Tile{i, j + 1})
			}
			if j-1 >= 0 && j-1 < dimension {
				b.AddNeighbor(Tile{i, j}, Tile{i, j - 1})
			}
			if i+1 >= 0 && i+1 < dimension {
				b.AddNeighbor(Tile{i, j}, Tile{i + 1, j})
			}
			if i-1 >= 0 && i-1 < dimension {
				b.AddNeighbor(Tile{i, j}, Tile{i - 1, j})
			}
		}
	}
}

func (b Board) AddNeighbor(s Tile, g Tile) {
	b[s][g] = 1
	b[g][s] = 1
}

func (b Board) RemoveNeighbor(s Tile, g Tile) {
	delete(b[s], g)
	delete(b[g], s)
}

func (b Board) Neighbors(t Tile) []Tile {
	var neighbors []Tile
	for n, _ := range b[t] {
		neighbors = append(neighbors, n)
	}
	return neighbors
}

func (b Board) Actions(s State) []Action {
	t := s.(Tile)
	var actions []Action
	for _, n := range b.Neighbors(t) {
		actions = append(actions, BoardAction{t, n})
	}
	return actions
}

func (b Board) Transition(s State, a Action) (State, error) {
	t := s.(Tile)
	m := a.(BoardAction)
	if t != m.from {
		return t, fmt.Errorf("BoardAction: state and action's from missmatch %v -%v", t, m.from)
	}
	return m.to, nil
}

func (b Board) StepCost(s State, a Action) (int, error) {
	t := s.(Tile)
	m := a.(BoardAction)
	if t != m.from {
		return 0, fmt.Errorf("BoardAction: state and action's from missmatch %v -%v", t, m.from)
	}
	return m.from.ManhattanDistance(m.to), nil
}

func (b BoardProblem) InitialState() State {
	return b.StartTile
}

func (b BoardProblem) GoalTest(s State) bool {
	for _, goalTile := range b.GoalTiles {
		if goalTile == s.(Tile) {
			return true
		}
	}
	return false
}

func (b BoardProblem) Heuristic(s State) int {
	min := math.MaxInt64
	for _, goalTile := range b.GoalTiles {
		straightDistance := s.(Tile).ManhattanDistance(goalTile)
		diagonalDistance := s.(Tile).DiagonalDistance(goalTile)
		heuristicDistance := 1000*diagonalDistance + 1414*(straightDistance-2*diagonalDistance)
		heuristicDistance += 1000 * s.(Tile).CrossProductDistance(b.StartTile, goalTile)
		min = Min(min, heuristicDistance)
	}
	return min
}

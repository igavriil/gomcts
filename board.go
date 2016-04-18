package main

import "fmt"

type Tile struct {
	i, j int
}

func (t Tile) ManhattanDistance(g Tile) int {
	return Abs(t.i-g.i) + Abs(t.j-g.j)
}

type BoardAction struct {
	from Tile
	to   Tile
}

type Board map[Tile]map[Tile]int

type BoardProblem struct {
	Board
	StartTile Tile
	GoalTile  Tile
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
	return b.GoalTile == s.(Tile)
}

func (b BoardProblem) Heuristic(s State) int {
	return b.GoalTile.ManhattanDistance(s.(Tile))
}

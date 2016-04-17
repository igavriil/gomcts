package main

import "math"

type Game interface {
	InitialState() State
	Actions(s State) []Action
	Result(s State, a Action) State
	TerminalTest(s State) bool
	Player(s State) Player
	Utility(s State, p Player) int
}

func GameSuccessors(g Game, s State) []State {
	successors := make([]State, len(g.Actions(s)))
	for _, a := range g.Actions(s) {
		child := g.Result(s, a)
		successors = append(successors, child)
	}
	return successors
}

func MinMax(g Game, s State, depth int) int {
	p := g.Player(s)

	if (depth == 0) || g.TerminalTest(s) {
		return g.Utility(s, p)
	}
	if p.IsMax() {
		max := math.MinInt64
		for _, child := range GameSuccessors(g, s) {
			v := MinMax(g, child, depth-1)
			max = Max(max, v)
		}
		return max
	} else {
		min := math.MaxInt64
		for _, child := range GameSuccessors(g, s) {
			v := MinMax(g, child, depth-1)
			min = Min(min, v)
		}
		return min
	}
}

func AlphaBeta(g Game, s State, depth int, alpha int, beta int) int {
	p := g.Player(s)

	if (depth == 0) || g.TerminalTest(s) {
		return g.Utility(s, p)
	}
	if p.IsMax() {
		max := math.MinInt64
		for _, child := range GameSuccessors(g, s) {
			max = Max(max, AlphaBeta(g, child, depth-1, alpha, beta))
			alpha := Max(alpha, max)
			if beta <= alpha {
				break
			}
		}
		return max
	} else {
		min := math.MaxInt64
		for _, child := range GameSuccessors(g, s) {
			min = Min(min, AlphaBeta(g, child, depth-1, alpha, beta))
			beta := Min(beta, min)
			if beta <= alpha {
				break
			}
		}
		return min
	}
}

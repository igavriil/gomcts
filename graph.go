package main

type Graph interface {
	InitialState() State
	Actions(s State) []Action
	Transition(s State, a Action) State
	GoalTest(s State) bool
	StepCost(s State, a Action) int
}

func GraphSuccessors(g Graph, s State) []State {
	successors := make([]State, len(g.Actions(s)))
	for _, a := range g.Actions(s) {
		child := g.Transition(s, a)
		successors = append(successors, child)
	}
	return successors
}

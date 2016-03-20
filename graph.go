package main

type Graph interface {
	InitialState() State
	Actions(s State) []Action
	Transition(s State, a Action) State
	GoalTest(s State) bool
	StepCost(s State, a Action) int
}

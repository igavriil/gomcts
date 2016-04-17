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

func UninformedSearch(g Graph, f Frontier) *Node {
	state := g.InitialState()
	node := &Node{State: state}
	if g.GoalTest(state) {
		return node
	}
	f.Push(node)
	explored := make(map[State]bool)
	queued := make(map[State]bool)

	for f.Len() > 0 {
		node = f.Pop().(*Node)
		state = (*node).State
		explored[state] = true
		for _, action := range g.Actions(state) {
			childNode := node.ChildNode(g, action)
			state = (*childNode).State

			if e, _ := explored[state]; !e {
				if q, _ := queued[state]; !q {
					if g.GoalTest(state) {
						return childNode
					}
					f.Push(childNode)
					queued[state] = true
				}
			}
		}
	}
}

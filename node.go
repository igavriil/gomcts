package main

type GameNode struct {
	state  State
	parent *GameNode
	action Action
}

func (n GameNode) ChildNode(g Game, a Action) *GameNode {
	s := g.Result(n.state, a)
	return &GameNode{s, &n, a}
}

type GraphNode struct {
	State    State
	Parent   *GraphNode
	Action   Action
	PathCost int
}

func (n GraphNode) ChildNode(g Graph, a Action) *GraphNode {
	s := g.Transition(n.State, a)
	c := n.PathCost + g.StepCost(n.State, a)
	return &GraphNode{s, &n, a, c}
}

package main

import "container/heap"

type Graph interface {
	InitialState() State
	Actions(s State) []Action
	Transition(s State, a Action) (State, error)
	GoalTest(s State) bool
	StepCost(s State, a Action) (int, error)
}

type ProblemSpecificGraph interface {
	Graph
	Heuristic(s State) int
}

func UninformedSearch(g Graph, f Frontier) *GraphNode {
	state := g.InitialState()
	node := &GraphNode{State: state}
	if g.GoalTest(state) {
		return node
	}
	f.Push(node)
	explored := make(map[State]bool)
	queued := make(map[State]bool)

	for f.Len() > 0 {
		node = f.Pop().(*GraphNode)
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
	return nil
}

func AstarSearch(g ProblemSpecificGraph, pq PriorityQueue) *GraphNode {
	state := g.InitialState()
	node := &GraphNode{State: state}
	if g.GoalTest(state) {
		return node
	}
	item := &Item{
		value:    node,
		priority: node.PathCost + g.Heuristic(node.State),
	}
	heap.Push(&pq, item)
	explored := make(map[State]bool)
	queued := make(map[State]bool)

	for pq.Len() > 0 {
		node = heap.Pop(&pq).(*Item).value.(*GraphNode)
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
					childItem := &Item{
						value:    childNode,
						priority: childNode.PathCost + g.Heuristic(childNode.State),
					}
					heap.Push(&pq, childItem)
					queued[state] = true
				}
			}
		}
	}
	return nil
}

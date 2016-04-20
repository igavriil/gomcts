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
		state = node.State
		explored[state] = true
		for _, action := range g.Actions(state) {
			childNode := node.ChildNode(g, action)
			childState := childNode.State

			if _, e := explored[childState]; !e {
				if _, q := queued[childState]; !q {
					if g.GoalTest(childState) {
						return childNode
					}
					f.Push(childNode)
					queued[childState] = true
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
	queued := make(map[State]*Item)

	queued[state] = item

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		node = item.value.(*GraphNode)
		state = node.State

		explored[state] = true
		delete(queued, state)
		for _, action := range g.Actions(state) {
			childNode := node.ChildNode(g, action)
			childState := childNode.State
			if g.GoalTest(childState) {
				return childNode
			}

			childItem := &Item{
				value:    childNode,
				priority: childNode.PathCost + g.Heuristic(childState),
			}

			if queuedItem, q := queued[childState]; q {
				if queuedItem.priority < childItem.priority {
					heap.Remove(&pq, queuedItem.index)
					delete(explored, childState)
					delete(queued, childState)
				}
			}
			if _, e := explored[childState]; !e {
				if _, q := queued[childState]; !q {
					heap.Push(&pq, childItem)
					queued[childState] = childItem
				}
			}
		}
	}
	return nil
}

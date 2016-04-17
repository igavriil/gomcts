package main

import "container/heap"

type Element interface{}

type Frontier interface {
	Push(n Element)
	Pop() (n Element)
	Top() (n Element)
	Len() int
}

type Queue []Element

func (q *Queue) Push(n Element) {
	*q = append(*q, n)
}

func (q *Queue) Pop() (n Element) {
	n = (*q)[0]
	*q = (*q)[1:]
	return n
}

func (q *Queue) Len() int {
	return len(*q)
}

func (q *Queue) Top() (n Element) {
	n = (*q)[0]
	return n
}

type Stack []Element

func (s *Stack) Push(n Element) {
	*s = append(*s, n)
}

func (s *Stack) Pop() (n Element) {
	x := s.Len() - 1
	n = (*s)[x]
	*s = (*s)[:x]
	return n
}

func (s *Stack) Len() int {
	return len(*s)
}

func (s *Stack) Top() (n Element) {
	x := s.Len() - 1
	n = (*s)[x]
	return n
}

type Item struct {
	value    interface{}
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	item.index = -1
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

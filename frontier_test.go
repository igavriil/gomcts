package main

import (
	"container/heap"
	"reflect"
	"testing"
)

func TestQueuePush(t *testing.T) {
	f := Queue{}
	f.Push(1)
	f.Push(2)
	expected := Queue{1, 2}
	equal := reflect.DeepEqual(f, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", f,
		)
	}
}

func TestQueuePop(t *testing.T) {
	f := Queue{1, 2, 3}
	expected := 1
	actual := f.Pop()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	expectedQueue := Queue{2, 3}
	equal = reflect.DeepEqual(f, expectedQueue)
	if !equal {
		t.Error(
			"expected", expectedQueue,
			"got", f,
		)
	}
}

func TestQueueTop(t *testing.T) {
	f := Queue{1, 2, 3}
	expected := 1
	actual := f.Top()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	expectedQueue := Queue{1, 2, 3}
	equal = reflect.DeepEqual(f, expectedQueue)
	if !equal {
		t.Error(
			"expected", expectedQueue,
			"got", f,
		)
	}
}

func TestQueueLen(t *testing.T) {
	f := Queue{1, 2, 3}
	expected := 3
	actual := f.Len()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestStackPush(t *testing.T) {
	f := Stack{}
	f.Push(1)
	f.Push(2)
	expected := Stack{1, 2}
	equal := reflect.DeepEqual(f, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", f,
		)
	}
}

func TestStackPop(t *testing.T) {
	f := Stack{1, 2, 3}
	expected := 3
	actual := f.Pop()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	expectedQueue := Stack{1, 2}
	equal = reflect.DeepEqual(f, expectedQueue)
	if !equal {
		t.Error(
			"expected", expectedQueue,
			"got", f,
		)
	}
}

func TestStackTop(t *testing.T) {
	f := Stack{1, 2, 3}
	expected := 3
	actual := f.Top()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	expectedQueue := Stack{1, 2, 3}
	equal = reflect.DeepEqual(f, expectedQueue)
	if !equal {
		t.Error(
			"expected", expectedQueue,
			"got", f,
		)
	}
}

func TestStackLen(t *testing.T) {
	f := Stack{1, 2, 3}
	expected := 3
	actual := f.Len()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPriorityQueuePush(t *testing.T) {
	pq := setupPriorityQueue()

	expected := []Item{
		Item{value: Tile{1, 1}, priority: 1, index: 0},
		Item{value: Tile{2, 2}, priority: 2, index: 1},
		Item{value: Tile{3, 3}, priority: 3, index: 2},
	}
	actual := derefencePriorityQueue(pq)
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPriorityQueuePop(t *testing.T) {
	pq := setupPriorityQueue()
	expected := Tile{1, 1}
	actual := heap.Pop(&pq).(*Item).value
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	expected = Tile{2, 2}
	actual = heap.Pop(&pq).(*Item).value
	equal = reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPriorityQueueLen(t *testing.T) {
	pq := setupPriorityQueue()
	expected := 3
	actual := pq.Len()
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	item := &Item{value: Tile{4, 4}, priority: 4}
	heap.Push(&pq, item)
	expected = 4
	actual = pq.Len()
	equal = reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
	heap.Pop(&pq)
	expected = 3
	actual = pq.Len()
	equal = reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}

func TestPriorityQueueUpdate(t *testing.T) {
	pq := setupPriorityQueue()
	expected := Tile{5, 5}
	item := pq[pq.Len()-1]
	pq.update(item, Tile{5, 5}, 0)

	actual := heap.Pop(&pq).(*Item).value
	equal := reflect.DeepEqual(actual, expected)
	if !equal {
		t.Error(
			"expected", expected,
			"got", actual,
		)
	}
}
func setupPriorityQueue() PriorityQueue {
	pq := make(PriorityQueue, 0)
	item3 := &Item{value: Tile{3, 3}, priority: 3}
	item1 := &Item{value: Tile{1, 1}, priority: 1}
	item2 := &Item{value: Tile{2, 2}, priority: 2}
	heap.Push(&pq, item1)
	heap.Push(&pq, item2)
	heap.Push(&pq, item3)
	return pq
}

func derefencePriorityQueue(pq PriorityQueue) []Item {
	result := make([]Item, 0)
	for _, i := range pq {
		result = append(result, *i)
	}
	return result
}

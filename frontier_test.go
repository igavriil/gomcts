package main

import (
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

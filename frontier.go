package main

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

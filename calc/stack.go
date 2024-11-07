package calc

import (
	"errors"
)

var (
	ErrStackOverflow  = errors.New("stack overflow")
	ErrStackUnderflow = errors.New("stack underflow")
)

type stack struct {
	elements []int
	top      int
}

func newStack(size int) *stack {
	return &stack{
		elements: make([]int, size),
		top:      0,
	}
}

func (s *stack) Push(el int) error {
	if s.top >= len(s.elements) {
		return ErrStackOverflow
	}
	s.elements[s.top] = el
	s.top++
	return nil
}

func (s *stack) Pop() (int, error) {
	if s.top == 0 {
		return 0, ErrStackUnderflow
	}
	s.top--
	return s.elements[s.top], nil
}

func (s *stack) Pop2() (int, int, error) {
	if s.top < 2 {
		return 0, 0, ErrStackUnderflow
	}
	s.top -= 2
	return s.elements[s.top+1], s.elements[s.top], nil
}

func (s *stack) Clear() {
	s.top = 0
}

func (s *stack) Snapshot() []int {
	return s.elements[:s.top]
}

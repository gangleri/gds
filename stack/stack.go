package stack

import (
	"errors"
)

type Stack interface {
	Push(interface{})
	Pop() interface{}
}

type ArrayStack struct {
	index uint
	stuff []interface{}
}

func NewArrayStack(size uint) (s *ArrayStack) {
	s = &ArrayStack{
		index: 0,
		stuff: make([]interface{}, size),
	}
	return
}

func (s *ArrayStack) Push(item interface{}) {
	if s.index == uint(len(s.stuff)) {
		s.stuff = resize(uint(len(s.stuff))*2, s.stuff)
	}

	s.stuff[s.index] = item
	s.index++
}

func (s *ArrayStack) Pop() (i interface{}, e error) {
	if s.index < 1 {
		i, e = nil, errors.New("StackEmpty")
		return
	}

	s.index--
	i = s.stuff[s.index]
	s.stuff[s.index] = nil

	if s.index > 0 && s.index <= uint(len(s.stuff)/4) {
		s.stuff = resize(uint(len(s.stuff)/2), s.stuff)
	}

	return
}

func resize(newCapacity uint, o []interface{}) (a []interface{}) {
	a = make([]interface{}, newCapacity)
	for i := 0; i < len(o); i++ {
		a[i] = o[i]
	}
	return
}

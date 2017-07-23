package stack

import (
	"testing"
)

func TestNewArrayStatck(t *testing.T) {
	s := NewArrayStack(10)
	if len(s.stuff) != 10 {
		t.Errorf("Expected stack to have length 10 but got length [%d]", len(s.stuff))
	}
}

func TestPushToArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	s.Push("Hello!")
	if s.stuff[0] != "Hello!" {
		t.Errorf("Expected to retrieve \"Hello!\" but got [%s]", s.stuff[0])
	}

	if s.index != 1 {
		t.Errorf("Expected to have statck index incremented.")
	}
}

func TestPopFromArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	s.Push("Thing to pop")
	i, e := s.Pop()

	if i != "Thing to pop" {
		t.Errorf("Expected to pop [Thing to pop] but got [%v]", i)
	}

	if e != nil {
		t.Errorf("There should be no error during Pop() [%v]", e)
	}

	if s.index != 0 {
		t.Errorf("Expected to have index decremented")
	}

	if s.stuff[0] != nil {
		t.Errorf("Statck element should contain nil")
	}
}

func TestPushIntToArrayStack(t *testing.T) {
	s := NewArrayStack(10)
	s.Push(13)
	if s.stuff[0] != 13 {
		t.Errorf("Expected to retrieve [13] but got [%d]", s.stuff[0])
	}

	if s.index != 1 {
		t.Errorf("Expected to have statck index incremented.")
	}
}

func TestPushAndGrow(t *testing.T) {
	s := NewArrayStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	if len(s.stuff) != 3 {
		t.Error("Expected to have a 3 element stack")
	}

	s.Push(4)

	if len(s.stuff) != 6 {
		t.Error("Expected to have a capactity of 6")
	}
}

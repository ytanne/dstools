package stack

type stack struct {
	vals   []any
	length int
}

func NewStack() *stack {
	return &stack{
		vals: make([]any, 0),
	}
}

func (s *stack) Push(x any) {
	s.vals = append(s.vals, x)
	s.length++
}

func (s *stack) Peek() (any, bool) {
	if s.length == 0 {
		return nil, false
	}

	l := len(s.vals)
	return s.vals[l-1], true
}

func (s *stack) Pop() (any, bool) {
	if s.length == 0 {
		return nil, false
	}

	l := len(s.vals)
	val := s.vals[l-1]
	s.vals = s.vals[:l-1]
	s.length--

	return val, true
}

func (s *stack) Size() int {
	return s.length
}

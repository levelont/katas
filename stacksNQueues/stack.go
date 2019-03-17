package stacksNQueues

type stacker interface {
	Add(v string)
	Pop() string
}

type stack struct {
	values []string
}

func (s *stack) Add(v string) {
	s.values = append(s.values, v)
}

func (s *stack) Pop() string {
	if len(s.values) == 0 {
		return ""
	}

	top := s.values[len(s.values)-1]

	if len(s.values) > 1 {
		s.values = s.values[:len(s.values)-1]
	} else {
		s.values = make([]string, 0)
	}
	return top
}

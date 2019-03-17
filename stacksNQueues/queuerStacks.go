package stacksNQueues

//A queue implemented with two stacks

type queuerStacks struct {
	first  stack
	second stack
}

func (qs *queuerStacks) Add(v string) {
	qs.first.Add(v)
}

func (qs *queuerStacks) Deque() string {
	numSwaps := len(qs.first.values)
	for i := 0; i < numSwaps; i++ {
		v := qs.first.Pop()
		qs.second.Add(v)
	}

	top := qs.second.Pop()

	numSwaps = len(qs.second.values)
	for i := 0; i < numSwaps; i++ {
		v := qs.second.Pop()
		qs.first.Add(v)
	}

	return top
}

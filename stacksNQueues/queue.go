package stacksNQueues

type queuer interface {
	Add(v string)
	Deque() string
}

type queue struct {
	values []string
}

func (q *queue) Add(v string) {
	q.values = append([]string{v}, q.values...)
}

func (q *queue) Deque() string {
	if len(q.values) == 0 {
		return ""
	}

	top := q.values[0]

	if len(q.values) > 1 {
		q.values = q.values[1:]
	} else {
		q.values = make([]string, 0)
	}
	return top
}

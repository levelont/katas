package stacksNQueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_QueuerStacks_Deque(t *testing.T) {
	type fields struct {
		values []string
	}
	tests := []struct {
		name      string
		fields    fields
		want      string
		postState fields
	}{
		{
			name:      "single value",
			fields:    fields{[]string{"1"}},
			want:      "1",
			postState: fields{[]string{}},
		},

		{
			name:      "two values",
			fields:    fields{[]string{"1", "2"}},
			want:      "1",
			postState: fields{[]string{"2"}},
		},

		{
			name:      "three values",
			fields:    fields{[]string{"1", "2", "3"}},
			want:      "1",
			postState: fields{[]string{"2", "3"}},
		},

		{
			name:      "empty stack",
			fields:    fields{[]string{}},
			want:      "",
			postState: fields{[]string{}},
		},

		{
			name:      "nil stack",
			fields:    fields{nil},
			want:      "",
			postState: fields{nil},
		},
	}
	for _, tt := range tests {
		queue := &queuerStacks{
			first:  stack{tt.fields.values},
			second: stack{make([]string, 0)},
		}
		assert.Equal(t, tt.want, queue.Deque(), tt.name)
		assert.Equal(t, tt.postState.values, queue.first.values, tt.name)
		assert.Equal(t, []string{}, queue.second.values, tt.name)
	}
}

package stacksNQueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_queue_Deque(t *testing.T) {
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
		t.Run(tt.name, func(t *testing.T) {
			queue := &queue{
				values: tt.fields.values,
			}
			if got := queue.Deque(); got != tt.want {
				t.Errorf("queue.Deque() = %v, want %v", got, tt.want)
			}

			assert.Equal(t, tt.postState.values, queue.values, tt.name)
		})
	}
}

func Test_queue_Add(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		postState []string
	}{
		{
			name:      "Add three elements",
			args:      []string{"1", "2", "3"},
			postState: []string{"3", "2", "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var q queue
			for _, v := range tt.args {
				q.Add(v)
			}
			assert.Equal(t, tt.postState, q.values)
		})
	}
}

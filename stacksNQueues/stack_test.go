package stacksNQueues

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_stack_Pop(t *testing.T) {
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
			want:      "2",
			postState: fields{[]string{"1"}},
		},

		{
			name:      "three values",
			fields:    fields{[]string{"1", "2", "3"}},
			want:      "3",
			postState: fields{[]string{"1", "2"}},
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
		s := &stack{
			values: tt.fields.values,
		}
		assert.Equal(t, tt.want, s.Pop(), tt.name)
		assert.Equal(t, tt.postState.values, s.values, tt.name)
	}
}

func Test_stack_Add(t *testing.T) {
	tests := []struct {
		name      string
		args      []string
		postState []string
	}{
		{
			name:      "Add three elements",
			args:      []string{"1", "2", "3"},
			postState: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		var s stack
		for _, v := range tt.args {
			s.Add(v)
		}
		assert.Equal(t, tt.postState, s.values, tt.name)
	}
}

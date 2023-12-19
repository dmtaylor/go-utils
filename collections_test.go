package go_utils

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStableMap(t *testing.T) {
	type testCase[K any, V any] struct {
		name string
		in   []K
		f    func(K) V
		want []V
	}
	tests := []testCase[any, any]{
		{
			"empty",
			[]any{},
			func(s any) any {
				return s
			},
			[]any{},
		},
		{
			"int_to_string",
			[]any{int64(10), int64(20), int64(30)},
			func(i any) any {
				return strconv.FormatInt(i.(int64), 10)
			},
			[]any{"10", "20", "30"},
		},
		{
			"add1",
			[]any{10, 20, 30},
			func(i any) any {
				return i.(int) + 1
			},
			[]any{11, 21, 31},
		},
		// TODO add additional tests
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StableMap(tt.in, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestConcurrentMap(t *testing.T) {
	type testCase[K any, V any] struct {
		name string
		in   []K
		f    func(K) V
		want []V
	}
	tests := []testCase[any, any]{
		{
			"empty",
			[]any{},
			func(s any) any {
				return s
			},
			[]any{},
		},
		{
			"int_to_string",
			[]any{int64(10), int64(20), int64(30)},
			func(i any) any {
				return strconv.FormatInt(i.(int64), 10)
			},
			[]any{"10", "20", "30"},
		},
		{
			"add1",
			[]any{10, 20, 30},
			func(i any) any {
				return i.(int) + 1
			},
			[]any{11, 21, 31},
		},
		// TODO add additional test cases
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ConcurrentMap(tt.in, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStableFilter(t *testing.T) {
	type testCase[K any] struct {
		name string
		in   []K
		f    func(K) bool
		want []K
	}
	tests := []testCase[any]{
		{
			"empty",
			[]any{},
			func(a any) bool {
				return true
			},
			[]any{},
		},
		{
			"clear_list",
			[]any{10, 20, 30},
			func(i any) bool {
				return false
			},
			[]any{},
		},
		{
			"evens_only",
			[]any{2, 3, 4},
			func(i any) bool {
				return i.(int)%2 == 0
			},
			[]any{2, 4},
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StableFilter(tt.in, tt.f)
			assert.Equal(t, tt.want, got)
		})
	}
}

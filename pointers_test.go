package go_utils

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPtr(t *testing.T) {
	type testStruct struct {
		value int
	}
	type testCase[A any] struct {
		name string
		arg  A
	}
	tests := []testCase[any]{
		{
			"string",
			"foo",
		},
		{
			"int",
			42,
		},
		{
			"struct",
			testStruct{value: 42},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Ptr(tt.arg)
			pType := reflect.TypeOf(got)
			if assert.Equal(t, reflect.Ptr, pType.Kind()) {
				assert.Equal(t, tt.arg, *got)
			}
		})
	}
}

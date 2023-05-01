package gogenerics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type GetterSetter[T any] interface {
	GetValue() T
	SetValue(value T)
}

func ChangeValue[T any](param GetterSetter[T], value T) T {
	param.SetValue(value)
	return param.GetValue()
}

type MyData[T any] struct {
	Value T
}

func (d *MyData[T]) GetValue() T {
	return d.Value
}

func (d *MyData[T]) SetValue(value T) {
	d.Value = value
}

func TestInterface(t *testing.T) {
	MyData := MyData[string]{}
	result := ChangeValue[string](&MyData, "name")

	assert.Equal(t, "name", result)
	assert.Equal(t, "name", MyData.Value)
}

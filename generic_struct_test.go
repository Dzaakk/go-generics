package gogenerics

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type Data[T any] struct {
	First  T
	Second T
}

func (d *Data[_]) SayHello(name string) string {
	return "Hello " + name
}

func (d *Data[T]) ChangeFirst(first T) T {
	d.First = first
	return d.First
}
func TestData(t *testing.T) {
	data := Data[string]{
		First:  "First Name",
		Second: "Last Name",
	}
	fmt.Println(data)
}
func TestGenericMethod(t *testing.T) {
	data := Data[string]{
		First:  "First Name",
		Second: "Last Name",
	}
	assert.Equal(t, "Name", data.ChangeFirst("Name"))
	assert.Equal(t, "Hello User1", data.SayHello("User1"))
}

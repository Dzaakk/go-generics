package gogenerics

import (
	"fmt"
	"testing"
)

type Data[T any] struct {
	First  T
	Second T
}

func TestData(t *testing.T) {
	data := Data[string]{
		First:  "First Name",
		Second: "Last Name",
	}
	fmt.Println(data)
}

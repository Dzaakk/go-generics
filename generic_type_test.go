package gogenerics

import (
	"fmt"
	"testing"
)

type Bag[T any] []T

func PrintBag[T any](bag Bag[T]) {
	for _, value := range bag {
		fmt.Println(value)
	}
}

func TestBagString(t *testing.T) {
	names := Bag[string]{"name1", "name2"}
	PrintBag(names)
}

func TestBagInt(t *testing.T) {
	numbers := Bag[int]{1, 2, 3}
	PrintBag(numbers)
}

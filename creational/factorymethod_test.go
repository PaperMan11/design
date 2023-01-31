package creational

import (
	"testing"
)

func compute(o OperatorFactory, a, b int) int {
	operator := o.Create()
	operator.SetA(a)
	operator.SetB(b)
	return operator.Result()
}

func TestFactory(t *testing.T) {
	var factory OperatorFactory

	factory = PlusOperatorFactory{}
	if compute(factory, 1, 2) != 3 {
		t.Fatal("error with factory method pattern")
	}

	factory = MinusOperatorFactory{}
	if compute(factory, 1, 2) != -1 {
		t.Fatal("error with factory method pattern")
	}

}

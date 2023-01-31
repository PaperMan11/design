package creational

import "testing"

func TestBuilder(t *testing.T) {
	b1 := &Builder1{res: ""}
	d1 := NewDirector(b1)
	d1.Construct()
	if b1.GetRes() != "123" {
		t.Fatal("fail")
	}

	b2 := &Builder2{res: 0}
	d2 := NewDirector(b2)
	d2.Construct()
	if b2.GetRes() != 6 {
		t.Fatal("fail")
	}
}

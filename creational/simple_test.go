package creational

import "testing"

func TestHi(t *testing.T) {
	a := NewAPI(1)
	str := a.Say("tan")
	if str != "Hi: tan\n" {
		t.Fatal("Type1 test fail")
	}
}

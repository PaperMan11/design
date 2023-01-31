package structural

import "testing"

func TestAdapter(t *testing.T) {
	phone := NewPhone(NewAdapter(&V220{}))
	phone.Charge()
}

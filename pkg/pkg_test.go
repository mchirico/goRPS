package pkg

import "testing"

func TestAdd(t *testing.T) {
	a:=Add(1,2)
	if a != 3 {
		t.FailNow()
	}
}
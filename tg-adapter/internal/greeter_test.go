package internal

import (
	"testing"
)

func TestGreet(t *testing.T) {
	s := Greet()

	if s != "haha" {
		t.Errorf("It's not hehe %v", s)
	}
}

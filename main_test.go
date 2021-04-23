package main

import (
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := add(5, 3)
	if ans != 8 {
		t.Errorf("add(5, 3) = %d; want 8", ans)
	}

}

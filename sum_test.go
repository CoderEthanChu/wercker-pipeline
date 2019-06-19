package main

import "testing"

func TestAdd(t *testing.T) {
	actual := Add(2, 3)
	if actual != 5 {
		t.Fail()
	}
}

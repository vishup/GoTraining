package main

import "testing"

func TestExample(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	l := len(arr)

	if l != 10 {
		t.Error("Expected 10, found ", l)
	}
}

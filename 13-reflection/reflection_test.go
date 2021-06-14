package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, want %d but got %d", 1, len(got))
	}

	if got[0] != expected {
		t.Errorf("want %q but got %q", expected, got[0])
	}
}

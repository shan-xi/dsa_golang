package main

import "testing"

func TestCase1Success(t *testing.T) {
	got := palindrome_number(121)
	want := true
	if got != want {
		t.Errorf("%v != %v\n", got, want)
	}
}

func TestCase2Success(t *testing.T) {
	got := palindrome_number(-121)
	want := false
	if got != want {
		t.Errorf("%v != %v\n", got, want)
	}
}
func TestCase3Success(t *testing.T) {
	got := palindrome_number(10)
	want := false
	if got != want {
		t.Errorf("%v != %v\n", got, want)
	}
}

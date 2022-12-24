package main

import "testing"

func TestAddTwoNumbersCase1(t *testing.T) {
	n1 := buildSingleLinkedListFromArray([]int{3, 4, 2})
	n2 := buildSingleLinkedListFromArray([]int{4, 6, 5})
	got := stringResult(addTwoNumbers(n1, n2))
	want := "708"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddTwoNumbersCase2(t *testing.T) {
	n1 := buildSingleLinkedListFromArray([]int{0})
	n2 := buildSingleLinkedListFromArray([]int{0})
	got := stringResult(addTwoNumbers(n1, n2))
	want := "0"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

func TestAddTwoNumbersCase3(t *testing.T) {
	n1 := buildSingleLinkedListFromArray([]int{9, 9, 9, 9, 9, 9, 9})
	n2 := buildSingleLinkedListFromArray([]int{9, 9, 9, 9})
	got := stringResult(addTwoNumbers(n1, n2))
	want := "89990001"
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}
}

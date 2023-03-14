package main

import "testing"

func Test_case1_success(t *testing.T) {

	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	want := 5
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func Test_case2_success(t *testing.T) {

	got := maxProfit([]int{1, 2})
	want := 1
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

func Test_case3_success(t *testing.T) {

	got := maxProfit([]int{2, 1, 2, 1, 0, 1, 2})
	want := 2
	if got != want {
		t.Errorf("got %q, wanted %q", got, want)
	}

}

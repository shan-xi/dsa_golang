package main

import "testing"

func Test_case1_success(t *testing.T) {
	grid := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}
	got := numEnclaves(grid)
	want := 3
	if got != want {
		t.Errorf("got %v != want %v", got, want)
	}
}

func Test_case2_success(t *testing.T) {
	grid := [][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}
	got := numEnclaves(grid)
	want := 0
	if got != want {
		t.Errorf("got %v != want %v", got, want)
	}
}

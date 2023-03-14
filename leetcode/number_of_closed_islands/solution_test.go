package main

import "testing"

func Test_case1_success(t *testing.T) {
	got := closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 1, 0},
		{1, 0, 1, 0, 1, 1, 1, 0},
		{1, 0, 0, 0, 0, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1, 0}})
	want := 2
	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}

func Test_case2_success(t *testing.T) {
	got := closedIsland([][]int{
		{0, 0, 1, 0, 0},
		{0, 1, 0, 1, 0},
		{0, 1, 1, 1, 0},
	})
	want := 1
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func Test_case3_success(t *testing.T) {
	got := closedIsland([][]int{
		{1, 1, 1, 1, 1, 1, 1},
		{1, 0, 0, 0, 0, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 0, 1, 0, 1, 0, 1},
		{1, 0, 1, 1, 1, 0, 1},
		{1, 1, 1, 1, 1, 1, 1},
	})
	want := 2
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

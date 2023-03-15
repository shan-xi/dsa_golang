package main

import "testing"

func Test_case1_success(t *testing.T) {
	got := generate(5)

	want := make([][]int, 5)
	for i := range want {
		want[i] = make([]int, i+1)
	}
	want[0] = []int{1}
	want[1] = []int{1, 1}
	want[2] = []int{1, 2, 1}
	want[3] = []int{1, 3, 3, 1}
	want[4] = []int{1, 4, 6, 4, 1}

	for i := range got {
		for j := range got[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("got %v != want %v", got[i][j], want[i][j])
			}
		}
	}
}

func Test_case2_success(t *testing.T) {
	got := generate(1)

	want := make([][]int, 1)
	for i := range want {
		want[i] = make([]int, i+1)
	}
	want[0] = []int{1}
	for i := range got {
		for j := range got[i] {
			if got[i][j] != want[i][j] {
				t.Errorf("got %v != want %v", got[i][j], want[i][j])
			}
		}
	}
}

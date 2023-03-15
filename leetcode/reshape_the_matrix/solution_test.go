package main

import (
	"testing"
)

func Test_case1_success(t *testing.T) {
	mat := [][]int{{1, 2}, {3, 4}}
	r := 1
	c := 4
	got := matrixReshape(mat, r, c)
	want := [][]int{{1, 2, 3, 4}}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("got %v != want %v", got[i][j], want[i][j])
			}
		}
	}
}

func Test_case2_success(t *testing.T) {
	mat := [][]int{{1, 2}, {3, 4}}
	r := 2
	c := 4
	got := matrixReshape(mat, r, c)
	want := [][]int{{1, 2}, {3, 4}}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("got %v != want %v", got[i][j], want[i][j])
			}
		}
	}
}

func Test_case3_success(t *testing.T) {
	mat := [][]int{{1, 2}, {3, 4}}
	r := 4
	c := 1
	got := matrixReshape(mat, r, c)
	want := [][]int{{1}, {2}, {3}, {4}}
	for i := 0; i < len(got); i++ {
		for j := 0; j < len(got[i]); j++ {
			if got[i][j] != want[i][j] {
				t.Errorf("got %v != want %v", got[i][j], want[i][j])
			}
		}
	}
}

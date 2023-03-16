package main

import "testing"

func Test_case1_success(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	a := searchMatrix(matrix, 3)
	e := true
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}
func Test_case2_success(t *testing.T) {
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	a := searchMatrix(matrix, 13)
	e := false
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

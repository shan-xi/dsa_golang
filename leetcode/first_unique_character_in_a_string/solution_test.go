package main

import "testing"

func Test_case1_success(t *testing.T) {
	a := firstUniqChar("leetcode")
	e := 0
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}
func Test_case2_success(t *testing.T) {
	a := firstUniqChar("loveleetcode")
	e := 2
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}
func Test_case3_success(t *testing.T) {
	a := firstUniqChar("aabb")
	e := -1
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

package main

import "testing"

func Test_case1_success(t *testing.T) {
	a := canConstruct("a", "b")
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case2_success(t *testing.T) {
	a := canConstruct("aa", "ab")
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case3_success(t *testing.T) {
	a := canConstruct("aa", "aab")
	e := true
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

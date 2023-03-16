package main

import "testing"

func Test_case1_success(t *testing.T) {
	a := fib(2)
	e := 1
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

func Test_case2_success(t *testing.T) {
	a := fib(3)
	e := 2
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

func Test_case3_success(t *testing.T) {
	a := fib(4)
	e := 3
	if e != a {
		t.Errorf("e %v != a %v", e, a)
	}
}

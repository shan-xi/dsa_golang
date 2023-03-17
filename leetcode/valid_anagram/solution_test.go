package main

import "testing"

func Test_case1_success(t *testing.T) {
	a := isAnagram("anagram", "nagaram")
	e := true
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

func Test_case2_success(t *testing.T) {
	a := isAnagram("rat", "car")
	e := false
	if a != e {
		t.Errorf("a %v != e %v", a, e)
	}
}

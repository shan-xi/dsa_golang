package main

import "testing"

func TestCase1Success(t *testing.T) {
	got := longestPalindrome("babad")
	want := "bab"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase2Success(t *testing.T) {
	got := longestPalindrome("cbbd")
	want := "bb"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase3Success(t *testing.T) {
	got := longestPalindrome("abaceudekq")
	want := "aba"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase4Success(t *testing.T) {
	got := longestPalindrome("ababab")
	want := "ababa"
	want2 := "babab"
	if got != want && got != want2 {
		t.Errorf("%v != %v or %v", got, want, want2)
	}
}

func TestCase5Success(t *testing.T) {
	got := longestPalindrome("ceudekqaba")
	want := "aba"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase6Success(t *testing.T) {
	got := longestPalindrome("a")
	want := "a"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase7Success(t *testing.T) {
	got := longestPalindrome("ab")
	want := "a"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase8Success(t *testing.T) {
	got := longestPalindrome("ac")
	want := "a"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase9Success(t *testing.T) {
	got := longestPalindrome("aacabdkacaa")
	want := "aca"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase10Success(t *testing.T) {
	got := longestPalindrome("ccc")
	want := "ccc"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase11Success(t *testing.T) {
	got := longestPalindrome("aaabaaaa")
	want := "aaabaaa"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

func TestCase12Success(t *testing.T) {
	got := longestPalindrome("tattarrattat")
	want := "tattarrattat"
	if got != want {
		t.Errorf("%v != %v", got, want)
	}
}

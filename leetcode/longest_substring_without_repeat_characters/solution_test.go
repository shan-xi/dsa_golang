package main

import "testing"

func TestCase1Success(t *testing.T) {
	got := lengthOfLongestSubstring("abcabcbb")
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase2Success(t *testing.T) {
	got := lengthOfLongestSubstring("bbbbb")
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase3Success(t *testing.T) {
	got := lengthOfLongestSubstring("pwwkew")
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase4Success(t *testing.T) {
	got := lengthOfLongestSubstring("aab")
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCase5Success(t *testing.T) {
	got := lengthOfLongestSubstring("dvdf")
	want := 3
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCaseˊSuccess(t *testing.T) {
	got := lengthOfLongestSubstring("abba")
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

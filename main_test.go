package main

import "testing"

func TestIsInterleave(t *testing.T) {
	cases := []struct {
		s1, s2, s3 string
		want       bool
	}{
		{"aabcc", "dbbca", "aadbbcbcac", true},
		{"aabcc", "dbbca", "aadbbbaccc", false},
		{"", "", "", true},
		{"", "abc", "abc", true},
		{"abc", "", "abc", true},
		{"a", "b", "abx", false}, // length mismatch
		{"ab", "ab", "abab", true},
		{"aa", "ab", "aaba", true},
	}
	for _, c := range cases {
		if got := IsInterleave(c.s1, c.s2, c.s3); got != c.want {
			t.Errorf("IsInterleave(%q,%q,%q) = %v, want %v", c.s1, c.s2, c.s3, got, c.want)
		}
	}
}

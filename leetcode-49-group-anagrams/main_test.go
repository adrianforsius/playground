package main

import (
	"slices"
	"testing"
)

func TestHello(t *testing.T) {
	in := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	out := GroupAnagrams(in)
	expected := [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}}

	for i := 0; i < len(expected); i++ {
		slices.Equal(s1 S, s2 S)
		if !slices.Equal(out[i], expected[i]) {
			t.Fail()
		}
	}
}

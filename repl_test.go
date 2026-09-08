package main

import "testing"

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input string
		expected []string
	}{
		{
			input:"helLo woRld",
			expected:[]string{"hello" , "world"},
		},
		{
			input:" pikachu Chu    ",
			expected:[]string{"pikachu" , "chu"},
		},
		{
			input:"creating a        slice of test case",
			expected:[]string{"creating" , "a","slice","of","test","case"},
		},
	}

	for _, tc:= range cases {
		actual := cleanInput(tc.input)

		if len(actual) != len(tc.expected) {
           t.Errorf("expected  : %v, got: %v",tc.expected, actual)
		}

		for i,word := range actual {
			if word != tc.expected[i]{
               t.Errorf("expected: %v, got: %v", tc.expected[i], word)
			}
		}
 	}
}
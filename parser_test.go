package main

import "testing"

func TestParser(t *testing.T) {
	var tests = []struct {
		tokens []string
		ast Exp
	}{
		// Need to think about whether []string{"42"} would be wrapped in an expression or not
		{
			[]string{"42"},
			Int{42},
		},
		{
			[]string{"(", "+", "1", "2", ")"}, 
			Exp{[]interface{
				Identifier{"+"},
				Int{1},
				Int{2},
			},
		},
	}

	for _, test := range tests {
		if got, _ := Parse(test.tokens); !reflect.DeepEqual(got, test.ast) {
			t.Errorf("got %q, want %q", got, test.ast)
		}
	}
}
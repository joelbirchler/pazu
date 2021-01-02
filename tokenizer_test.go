package main

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	var tests = []struct {
		s    string
		want []string
	}{
		{"42", []string{"42"}},
		{"(+ 1 2 3)", []string{"(", "+", "1", "2", "3", ")"}},
		{`(str "paren (" "trap haha")`, []string{"(", "str", `"paren ("`, `"trap haha"`, ")"}},
		{"(fn (x)\n  x)", []string{"(", "fn", "(", "x", ")", "x", ")"}},
		{"(fn (x)\n  (+ x 1))   41 )", []string{"(", "fn", "(", "x", ")", "(", "+", "x", "1", ")", ")", "41", ")"}},
	}

	for _, test := range tests {
		if got, _ := Tokenize(test.s); !reflect.DeepEqual(got, test.want) {
			t.Errorf("got %q, want %q", got, test.want)
		}
	}
}

func TestTokenizeErr(t *testing.T) {
	var tests = []struct {
		s      string
		msg    string
		line   int
		column int
	}{
		{`(unclosed "quote yo)`, "literal not terminated", 1, 21},
		{`error on second line
			"test`, "literal not terminated", 2, 9},
	}

	for _, test := range tests {
		_, err := Tokenize(test.s)

		if serr, ok := err.(*SyntaxError); ok {
			if serr.Msg != test.msg || serr.Line != test.line || serr.Column != test.column {
				t.Errorf("got `%s` (%d, %d), want `%s` (%d, %d)", serr.Msg, serr.Line, serr.Column, test.msg, test.line, test.column)
			}
		} else {
			t.Errorf("got %v, expected SyntaxError", err)
		}
	}
}

package main

import (
	"testing"
	"reflect"
)

func TestParser(t *testing.T) {
	var tests = []struct {
		tokens []string
		ast interface{}
	}{
		{
			[]string{"42"},
			Atom{42},
		},
		{
			[]string{"(", "+", "1", "2", ")"}, 
			List{
				Elements: []interface{}{
					Atom{Symbol("+")},
					Atom{1},
					Atom{2},
				},
			},
		},
		{
			[]string{"(", "if", "foo", "(", "+", "1", "2", ")", "(", "+", "3", "4", ")", ")"}, 
			List{
				Elements: []interface{}{
					Atom{Symbol("if")},
					Atom{Symbol("foo")},
					List{
						Elements: []interface{}{
							Atom{Symbol("+")},
							Atom{1},
							Atom{2},
						},
					},
					List{
						Elements: []interface{}{
							Atom{Symbol("+")},
							Atom{3},
							Atom{4},
						},
					},
				},
			},
		},
		{
			[]string{"(", "hey", "(", "foo", "1", "(", "+", "2", "3", ")", ")", "(", "test", "test", ")", ")"}, 
			List{
				Elements: []interface{}{
					Atom{Symbol("hey")},
					List{
						Elements: []interface{}{
							Atom{Symbol("foo")},
							Atom{1},
							List{
								Elements: []interface{}{
									Atom{Symbol("+")},
									Atom{2},
									Atom{3},
								},
							},
						},
					},
					List{
						Elements: []interface{}{
							Atom{Symbol("test")},
							Atom{Symbol("test")},
						},
					},
				},
			},
		},
	}

	for _, test := range tests {
		if got, _ := Parse(test.tokens); !reflect.DeepEqual(got, test.ast) {
			t.Errorf("got %q, want %q", got, test.ast)
		}
	}
}
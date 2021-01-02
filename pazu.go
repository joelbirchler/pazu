package main

import (
	"fmt"
	"os"
)

func main() {
	eval("fn")
	eval("42")
	eval(`"hey"`)
	eval("(+ 1 2)")
	eval("(if foo (+ 1 2) (+ 3 4))")
	eval("(hey (foo 1 (+ 2 3)) (test test)")
}

func eval(s string) {
	tokens, err := Tokenize(s)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}

	ast, err := Parse(tokens)
	if err != nil {
		fmt.Printf("Error: %s", err)
	}
	print(ast)

	// returns something
}

func print(ast interface{}) {
	_print(ast, "\n")
}

func _print(ast interface{}, separator string) {
	switch v := ast.(type) {
	case List:
		fmt.Printf("( ")
		for _, e := range v.Elements {
			_print(e, " ")
		}
		fmt.Printf(")")
	case Atom:
		fmt.Printf("%T:%v ", v.Value, v.Value)
	default:
		fmt.Printf("?")
	}

	fmt.Printf(separator)
}

// TODO: clean it up and add tests
// TODO: ParseError
// TODO: add position to ast structs for future error messages
// TODO: improve printing (Stringer)
// TODO: repl w/ just read and parse (rppl)
//
// TODO: execute
// TODO: tokenizer should return multiple errors
// TODO: cons car cdr head tail
// TODO: quote lists '(a b c)
// TODO: pattern matching

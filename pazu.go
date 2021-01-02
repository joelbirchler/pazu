package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{} // TODO
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	fmt.Println("Pazu REPL v0.0.1 (Ctrl+D to exit)")

	p := prompt.New(
		executor,
		completer,
		prompt.OptionTitle("Pazu REPL"),
		prompt.OptionPrefixTextColor(prompt.Cyan),
		prompt.OptionPrefix("> "),
	)

	p.Run()
}

func executor(s string) {
	tokens, err := Tokenize(s)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	ast, err := Parse(tokens)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	print(ast)
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

// TODO: ParseError and tests
// TODO: add position to ast structs for future error messages
// TODO: improve printing (Stringer)
// TODO: organize a bit... main should be a repl that uses the pazu package
//
// TODO: execute
// TODO: tokenizer should return multiple errors
// TODO: cons car cdr head tail
// TODO: quoted lists '(a b c)
// TODO: pattern matching

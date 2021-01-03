package main

import (
	"fmt"

	"github.com/c-bata/go-prompt"
	"github.com/joelbirchler/pazu/pkg/pazu"
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
	tokens, err := pazu.Tokenize(s)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	ast, err := pazu.Parse(tokens)
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
	case pazu.List:
		fmt.Printf("( ")
		for _, e := range v.Elements {
			_print(e, " ")
		}
		fmt.Printf(")")
	case pazu.Atom:
		fmt.Printf("%T:%v ", v.Value, v.Value)
	default:
		fmt.Printf("?")
	}

	fmt.Printf(separator)
}

// TODO: ParseError and tests
// TODO: add position to ast structs for future error messages
//
// TODO: execute
// TODO: printer should return pazulang (pazu fmt = read-parse-print)
// TODO: tokenizer should return multiple errors
// TODO: cons car cdr head tail
// TODO: more lispy tokenizer (+a 1 2) should be ( +a 1 2 ) not ( + a 1 2 )
// TODO: quoted lists '(a b c)
// TODO: more number types
// TODO: language features: pattern matching, channels/actors, gpio, custom types, algebraic data types

package main

import (
	"fmt"
	"strconv"
	"unicode"
)

type List struct {
	Elements []interface{} // Elements may be Atoms or Lists
}

type Atom struct {
	Value interface{}
}

type Symbol string

// Parse takes an s-expression of string tokens and returns an abstract syntax tree. It returns
// a single item which may be an Atom (int, string, symbol) or a List.
func Parse(tokens []string) (interface{}, error) {
	if len(tokens) == 0 {
		return List{}, nil
	}
	
	if len(tokens) == 1 {
		return parseAtom(tokens[0])
	}

	if tokens[0] == "(" && tokens[len(tokens) - 1] == ")" {
		ast, remainder, err := parseList(tokens)
		if len(remainder) != 0 {
			fmt.Printf("REMAINDER: %v", remainder) // FIXME
		}
		return ast, err
	}

	return nil, fmt.Errorf("invalid expression")
}

func parseAtom(token string) (interface{}, error) {
	switch {
	case isString(token):
		return parseString(token)
	case isInt(token):
		return parseInt(token)
	case token == "(" || token == ")":
		return nil, fmt.Errorf("invalid symbol: %s", token)
	}

	return Atom{Symbol(token)}, nil
}

func isString(token string) bool {
	first := []rune(token)[0]
	last := []rune(token)[len(token) - 1]
	return first == '"' && last == '"'
}

func isInt(token string) bool {
	return unicode.IsDigit([]rune(token)[0]) // TODO: This could be more robust
}

func parseString(token string) (Atom, error) {
	s, err := strconv.Unquote(token)
	if err != nil {
		return Atom{}, fmt.Errorf("quote mismatch")
	}

	return Atom{s}, nil
}

func parseInt(token string) (Atom, error) {
	i, err := strconv.Atoi(token)
	if err != nil {
		return Atom{}, fmt.Errorf("could not '%s' convert to integer", token)
	}

	return Atom{i}, nil
}

func parseList(tokens []string) (List, []string, error) {
	tokens = tokens[1:]
	li := List{}

	var (
		err error
		el interface{}
	)

	for len(tokens) > 0 {
		head := tokens[0]
		tail := tokens[1:]
		
		switch head {
		case "(":
			el, tokens, err = parseList(tokens)
		case ")":
			return li, tail, nil
		default:
			el, err = parseAtom(head)
			tokens = tail
		}		

		if err != nil {
			return List{}, []string{}, err
		}

		li.Elements = append(li.Elements, el)
	}
	
	return li, tokens, fmt.Errorf("parenthesis mismatch")
}

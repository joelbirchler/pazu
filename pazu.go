package main

import (
	"fmt"
	"strings"
	"text/scanner"
)

func main() {
	fmt.Println("hi!")
}

type SyntaxError struct {
	Msg string
	Line int
	Column int
}
func (e *SyntaxError) Error() string { 
	return fmt.Sprintf("%s (line %d, col %d)", e.Msg, e.Line, e.Column)
}

func Tokenize(sexp string) (tokens []string, err error) {
	var s scanner.Scanner
	s.Init(strings.NewReader(sexp))
	s.Error = func(s *scanner.Scanner, msg string) {
		err = &SyntaxError{msg, s.Pos().Line, s.Pos().Column}
	}
	
	for tok := s.Scan(); tok != scanner.EOF; tok = s.Scan() {
		tokens = append(tokens, s.TokenText())
	}

	return
}

// type Exp struct {
// 	Value {}interface
// }

// func Parse() Exp {
// }


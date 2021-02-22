package peg

import (
	"fmt"
	"github.com/yhirose/go-peg"
	"strings"
)

func TestParser(grammar string, test string) (ast string) {

	tests := strings.Split(test, "~")
	results := []string{}
	formulaParser, err := peg.NewParser(grammar)
	if err != nil {
		results = append(results, err.Error())
		return strings.Join(results, "~")
	}

	for _, t := range tests {
		s, errP := formulaParser.ParseAndGetValue(t, nil)
		if errP != nil {
			results = append(results, errP.Error())
		} else {
			results = append(results, fmt.Sprint(s))
		}
	}
	return strings.Join(results, "~")
}

func Example() {
	grammar := GrammarExample()

	r := strings.Split(TestParser(grammar, TestExample()), "~")
	tests := strings.Split(TestExample(), "~")
	for i, s := range r {
		fmt.Println(tests[i], "=>", s)
	}
}

func GrammarExample() string {
	return `
# Add Comment
# Simple calculator
EXPR         ←  ATOM (BINOP ATOM)*
ATOM         ←  NUMBER / ('(' EXPR ')') / ('"' TEXT '"')
BINOP        ←  < [-+/*&] >
NUMBER       ←  < [0-9]+ >
TEXT         ←  < [A-Za-Z]+ >
%whitespace  ←  [ \t]*
---
# Expression parsing option
%expr  = EXPR   # Rule to apply 'precedence climbing method' to
%binop = L + -  # Precedence level 1
%binop = L * /  # Precedence level 2
%binop = L &    # Precedence level 3
`
}

func TestExample() string {
	return `1~"hello"&world~"hello"&"cool"&"world"~"hi "&"world"`
}

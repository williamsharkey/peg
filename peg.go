package peg

import (
	"fmt"
	"github.com/yhirose/go-peg"
	"strings"
)

func TestParser(grammar string, test string) (ast string) {

	tests := strings.Split(test, "\n")
	results := []string{}
	parser, err := peg.NewParser(grammar)
	if err != nil {
		results = append(results, "Grammar Parse Error: "+err.Error())
		return strings.Join(results, "\n")
	}

	parser.EnableAst()

	for _, t := range tests {
		s, errP := parser.ParseAndGetAst(t, nil)
		//s, errP := formulaParser.ParseAndGetValue(t, nil)
		if errP != nil {
			results = append(results, t+"\nError: \n"+errP.Error())
		} else {
			results = append(results, t+":\n"+fmt.Sprint(s))
		}
	}
	return strings.Join(results, "\n")
}

func GrammarExample() string {
	return `

# Grammar with strings
EXPR         ←  ATOM (BINOP ATOM)*
ATOM         ←  NUMBER / STRING / '(' EXPR ')' / REF
BINOP        ←  < [-+/*&] >
NUMBER       ←  < [0-9]+ >
REF          ←  < (!(BINOP).)* >
STRING       ←  < ["] (!('"')./'""')* ["] >
%whitespace  ←  [ \t]*
---
# Expression parsing option
%expr  = EXPR   # Rule to apply 'precedence climbing method' to
%binop = L + - &  # Precedence level 1
%binop = L * /  # Precedence level 2

`
}

func TestExample() string {
	return `1+1
hello
"hello world"&"abc"`
}

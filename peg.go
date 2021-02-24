package peg

import (
	"fmt"
	"github.com/yhirose/go-peg"
	"strings"
)

func TestParser(grammar string, test string) (results string) {

	parser, err := peg.NewParser(grammar)
	if err != nil {
		return "Grammar Parse Error: " + err.Error()
	}
	parser.EnableAst()

	for _, t := range strings.Split(test, "\n") {
		s, errP := parser.ParseAndGetAst(t, nil)
		if errP != nil {
			results += fmt.Sprintf("s:\nError: %s\n", t, errP)
		} else {
			results += fmt.Sprintf("%s:\n%s\n", t, s)
		}
	}
	return results
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

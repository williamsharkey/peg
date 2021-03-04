package peg

import (
	"fmt"
	"github.com/williamsharkey/go-peg"
	"strings"
)

func TestParser(grammar string, test string) (results string) {
	peg.CommentCharacterSet(":note:")
	parser, err := peg.NewParser(grammar)
	if err != nil {
		return "Grammar Parse Error: " + err.Error()
	}
	parser.EnableAst()

	for _, t := range strings.Split(test, "\n") {
		s, errP := parser.ParseAndGetAst(t, nil)
		if errP != nil {
			results += fmt.Sprintf("s:\nError: %s\n", t, errP.Error())
		} else {
			results += fmt.Sprintf("%s:\n%s\n", t, s)
		}
	}
	return results
}

func GrammarExample() string {
	return `

:note: Grammar with strings
EXPR         ←  ATOM (BINOP ATOM)*
ATOM         ←  NUMBER / STRING / ATFN / REF / '(' EXPR ')'
ATFN         ←  '@' <[A-Za-z]*> '(' <ATOM> ( ',' <ATOM>)*  ')' 
BINOP        ←  < [-+/*&] / '#OR#' >
NUMBER       ←  < [0-9]+ >
REF          ←  < (!(BINOP/'"'/ [ \t]).)+ >
STRING       ←  ["] < (!('"')./'""')*  > ["] [ \t]* 
%whitespace  ←  [ \t]*
---
:note: Expression parsing option
%expr  = EXPR     :note: Rule to apply 'precedence climbing method' to
%binop = L #OR#   :note: Precedence level 1
%binop = L + - &  :note: Precedence level 2
%binop = L * /    :note: Precedence level 3

`
}

func TestExample() string {
	return `1+1
hello
"hello world"&"abc"
"a"#OR#"b"`
}

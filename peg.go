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
ATFN         ←  '@' FN_NAME  ( '(' ( EXPR ( ',' EXPR)* )? ')' )?  
FN_NAME      ←  <[A-Za-z]*>
BINOP        ←  '<>' / '<=' / '>=' / '#OR#'/ [-+/*&=<>]
NUMBER       ←  < [0-9]+ ([.] [0-9]* )? >
REF          ←  [a-zA-Z$:0-9_]+
STRING       ←  ["] < (!('"')./'""')*  > ["] [ \t]* 
%whitespace  ←  [ \t]*
---
:note: Expression parsing option
%expr  = EXPR                  :note: apply precedence climbing to EXPR
%binop = L = #OR# < > <= >= <> :note: Precedence level 1
%binop = L + - &               :note: Precedence level 2
%binop = L * /                 :note: Precedence level 3

`
}

func TestExample() string {
	return `1+1
hello
"hello world"&"abc"
"a"#OR#"b"
1<>2`
}

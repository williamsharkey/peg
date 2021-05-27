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
COL ← [A-Z][A-Z]?
ADDR ← <COL ROW> 
LOCAL_ADDR ← (SHEET ':')? ADDR
FN_ADDR ← '<<' [a-zA-Z$:0-9_/]+ '>>' LOCAL_ADDR
SHEET ←  [A-Z]
ROW ← NUMBER
REF_FREE  ←  [a-zA-Z$:0-9_]+
REF ← FN_ADDR / LOCAL_ADDR > REF_FREE
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
	return `A
1
A1
AA1
AAA1
B:A
1+1
hello
"hello world"&"abc"
"a"#OR#"b"
1<>2
($CHECKDATE<=22.2)#OR#($INCLUDEWAIT=0)#OR#($QUARFLAG<>1)
($CHECKDATE<=<<L:\\Vru\\Flags\\Lincflag.Wk3>>$A$2)#OR#($INCLUDEWAIT=0)#OR#($QUARFLAG<>1)`
}

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
	printTrees := false
	for _, t := range strings.Split(test, "\n") {
		s, errP := parser.ParseAndGetAst(t, nil)
		if errP != nil {
			if printTrees {
				results += fmt.Sprintf("s:\nError: %s\n", t, errP.Error())
			} else {
				results += fmt.Sprintf("fail: %s\n", t)
			}

		} else {
			if printTrees {
				results += fmt.Sprintf("%s:\n%s\n", t, s)
			} else {
				results += fmt.Sprintf("ok: %s\n", t)
			}
		}
	}
	return results
}

func TestParserPrintTrees(grammar string, test string) (results string) {
	peg.CommentCharacterSet(":note:")
	parser, err := peg.NewParser(grammar)
	if err != nil {
		return "Grammar Parse Error: " + err.Error()
	}
	parser.EnableAst()
	printTrees := true
	for _, t := range strings.Split(test, "\n") {
		s, errP := parser.ParseAndGetAst(t, nil)
		if errP != nil {
			if printTrees {
				results += fmt.Sprintf("fail: %s\n", t)
				results += fmt.Sprintf("%s\n", errP.Error())
			} else {
				results += fmt.Sprintf("fail: %s\n", t)
			}

		} else {
			if printTrees {
				results += fmt.Sprintf("ok: %s\n\n%s\n\n\n", t, s)
			} else {
				results += fmt.Sprintf("ok: %s\n", t)
			}
		}
	}
	return results
}

func GrammarExample() string {
	return `

:note: Grammar with strings
FORMULA      ←  '+'? EXPR
EXPR         ←  ATOM (BINOP ATOM)*
ATOM         ← [ \t]* ( NUMBER / STRING / ATFN / URNEG / RANGE / REF / '(' EXPR ')' ) [ \t]*
ATFN         ←  '@' FN_NAME  ( '(' ( EXPR ( ',' EXPR)* )? ')' )?  
URNEG        ←  '-' EXPR
FN_NAME      ←  <[A-Za-z]*>
BINOP        ←  '<>' / '<=' / '>=' / '#OR#'/ [-+/*&=<>]
NUMBER       ←  < ([-])? [0-9]+ ([.] [0-9]* )? >
INTEGER      ←  < [0-9]+ >
COL          ← [A-Z][A-Z]?
ADDR         ← <'$'? COL '$'? ROW> 
LOCAL_ADDR   ← <( '$'? SHEET ':')? ADDR>
FN_ADDR      ← <'<<' [a-zA-Z$:0-9_/.\\]+ '>>' LOCAL_ADDR>
SHEET        ←  [A-Z]
ROW          ← INTEGER
REF_FREE     ←  [a-zA-Z$:0-9_.\\]+
REF          ← FN_ADDR / LOCAL_ADDR / REF_FREE
RANGE        ← REF'..'REF
STRING       ←  ["] < (!('"')./'""')*  > ["]
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
	return `@SUM()
@SUM(S6)
@SUM(S6..S8)
A
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
($CHECKDATE<=<<L:\\Vru\\Flags\\Lincflag.Wk3>>$A$2)#OR#($INCLUDEWAIT=0)#OR#($QUARFLAG<>1)
+"{goto}"&@CELL("coord",$WAITLOOPTIME)&"~{esc}@NOW+@TIME(0,10,0)~{edit}{calc}~{recalc waitparts}"
+"{if ($CHECKDATE<=<<"&$FLAGFILE&">>$A$2)#OR#($INCLUDEWAIT=0)#OR#($QUARFLAG<>1)}{branch startrun}"
+"{indicate "" "&$FLAGFILE&" Wrong Date, Need "&$I$7&", Waiting Until "&$I$5&":"&$I$6&" ""}{wait "&@STRING($WAITLOOPTIME,5)&"}{branch \B}"
+"{recalc \B}{system """&"md "&$B$303&@RIGHT("0000"&@STRING(@YEAR($ENDDATE)+1900,0),4)&""""&"}"
@IF(@HOUR($WAITLOOPTIME)<10,"0"&@STRING(@HOUR($WAITLOOPTIME),0),@STRING(@HOUR($WAITLOOPTIME),0))
@IF(@MINUTE($WAITLOOPTIME)<10,"0"&@STRING(@MINUTE($WAITLOOPTIME),0),@STRING(@MINUTE($WAITLOOPTIME),0))
@STRING(@MONTH($CHECKDATE),0)&"-"&@STRING(@DAY($CHECKDATE),0)&"-"&@STRING(@YEAR($CHECKDATE)+1900,0)`
}

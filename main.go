package main

import (
	"fmt"
	"github.com/yhirose/go-peg"
	"strings"
)

func TestParser(grammar string, testsJoined string) (resultsJoined string) {

	tests := strings.Split(testsJoined, "~")
	results := []string{}
	formulaParser, err := peg.NewParser(grammar)
	if err != nil {
		results = append(results, err.Error())
		return
	}

	for _, t := range tests {
		ast, err := formulaParser.ParseAndGetAst(t, nil)
		if err != nil {
			results = append(results, err.Error())
		} else {
			results = append(results, ast.String())
		}
	}
	return strings.Join(results, "~")
}

func Example() {
	grammar := `
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
	tests := []string{`"hello"&world`, `"hello"&"cool"&"world"`, `"hi "&"world"`}
	r := strings.Split(TestParser(grammar, strings.Join(tests, "~")), "~")

	for i, s := range r {
		fmt.Println(tests[i], "=>", s)
	}
}

func main() {
	Example()
}

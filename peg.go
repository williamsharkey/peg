package peg

import (
	"fmt"
	"github.com/yhirose/go-peg"
	"strings"
)

func TestParser(grammar string, test string) (ast string) {

	tests := strings.Split(test, "~")
	results := []string{}
	parser, err := peg.NewParser(grammar)
	if err != nil {
		results = append(results, err.Error())
		return strings.Join(results, "~")
	}

	parser.EnableAst()
	//g := parser.Grammar
	//g["EXPR"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
	//	val := v.ToInt(0)
	//	if v.Len() > 1 {
	//		ope := v.ToStr(1)
	//		rhs := v.ToInt(2)
	//		switch ope {
	//		case "+": val += rhs
	//		case "-": val -= rhs
	//		case "*": val *= rhs
	//		case "/": val /= rhs
	//		}
	//	}
	//	return val, nil
	//}
	//g["BINOP"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
	//	return v.Token(), nil
	//}
	//g["NUMBER"].Action = func(v *peg.Values, d peg.Any) (peg.Any, error) {
	//	return strconv.Atoi(v.Token())
	//}

	for _, t := range tests {
		s, errP := parser.ParseAndGetAst(t, nil)
		//s, errP := formulaParser.ParseAndGetValue(t, nil)
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
    # Simple calculator
    EXPR         ←  ATOM (BINOP ATOM)*
    ATOM         ←  NUMBER / '(' EXPR ')'
    BINOP        ←  < [-+/*] >
    NUMBER       ←  < [0-9]+ >
    %whitespace  ←  [ \t]*
    ---
    # Expression parsing option
    %expr  = EXPR   # Rule to apply 'precedence climbing method' to
    %binop = L + -  # Precedence level 1
    %binop = L * /  # Precedence level 2
`
}

func TestExample() string {
	return `1+1`
}

func main() {
	Example()
}

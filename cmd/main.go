package main

import (
	"fmt"
	"github.com/williamsharkey/peg"
)

func main() {
	fmt.Println(peg.TestParser(peg.GrammarExample(), peg.TestExample()))
}

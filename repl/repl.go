package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/vladwithcode/monkey-interpreter/lexer"
	"github.com/vladwithcode/monkey-interpreter/token"
)

const PROMPT = "$ "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Bytes()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("{Type: %s Literal: %q}\n", token.TokenTypeToString(tok.Type), tok.Literal)
		}
	}
}

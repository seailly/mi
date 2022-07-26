package repl

import (
	"bufio"
	"fmt"
	"github.com/seailly/mi/object"
	"io"

	"github.com/seailly/mi/evaluator"
	"github.com/seailly/mi/lexer"
	"github.com/seailly/mi/parser"
)

const PROMPT = "> "

// Start Run a REPL instance, reading io content into the lexer
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			_, err := io.WriteString(out, evaluated.Inspect())
			if err != nil {
				panic("Failed to write string")
			}

			_, err = io.WriteString(out, "\n")
			if err != nil {
				panic("Failed to write string")
			}
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		_, err := io.WriteString(out, "\t"+msg+"\n")
		if err != nil {
			panic("Failed to write string")
		}
	}
}

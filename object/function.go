package object

import (
	"fmt"
	"github.com/seailly/mi/ast"
	"strings"
)

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType {
	return FUNCTION_OBJECT
}

func (f *Function) Inspect() string {
	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	return fmt.Sprintf("fn(%s) {\n%s\n}", strings.Join(params, ", "), f.Body.String())
}

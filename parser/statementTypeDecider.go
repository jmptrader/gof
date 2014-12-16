package parser

import (
	"errors"
	"github.com/apoydence/GoF/parser/statementTypes"
)

func StatementTypeDecider(tokens []string) (statementTypes.Statement, error) {
	switch tokens[0] {
	case "func":
		return statementTypes.NewFunctionStatement(tokens), nil
	case "if":
		return statementTypes.NewIfStatement(tokens), nil
	case "match":
		return statementTypes.NewMatchStatement(tokens), nil
	default:
		if tokens[1] != "=" {
			return nil, errors.New("Unknown statement: " + tokens[0])
		}

		return statementTypes.NewDeclarationStatement(tokens), nil
	}
}

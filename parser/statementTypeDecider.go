package parser

import (
	"errors"
	"github.com/apoydence/GoF/parser/statementTypes"
)

func StatementTypeDecider(tokens []string) (statementTypes.Statement, error) {
	var statement statementTypes.Statement
	switch tokens[0] {
	case "func":
		statement = statementTypes.NewFunctionStatement()
	case "if":
		statement = statementTypes.NewIfStatement()
	case "match":
		statement = statementTypes.NewMatchStatement()
	default:
		if tokens[1] != "=" {
			return nil, errors.New("Unknown statement: " + tokens[0])
		}

		statement = statementTypes.NewDeclarationStatement()
	}

	statement.AddLine(tokens)
	return statement, nil
}

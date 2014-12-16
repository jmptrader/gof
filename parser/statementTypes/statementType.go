package statementTypes

type StatementType int

const (
	FunctionType StatementType = iota
	IfType
	MatchType
	DeclarationType
)

type Statement interface {
	Type() StatementType
	AddLine(tokens []string)
}

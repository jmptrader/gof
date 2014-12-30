package expressionParsing

type blockSpec struct {
	block     string
	valueType TypeName
}

func toBlockSpec(opQueue []string) []*blockSpec {
	result := make([]*blockSpec, 0)
	for _, o := range opQueue {
		result = append(result, newBlockSpec(o, ""))
	}
	return result
}

func newBlockSpec(op string, valueType TypeName) *blockSpec {
	return &blockSpec{
		block:     op,
		valueType: valueType,
	}
}

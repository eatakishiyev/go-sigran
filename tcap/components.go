package tcap

type ComponentType uint

const (
	InvokeOp ComponentType = iota + 1
	ReturnResultLastOp
)

type ComponentPortion struct {
	Component interface{} `asn1:"choice:components"`
}

type Component interface {
	GetType() ComponentType
}

type Invoke struct {
	InvokeId      int    `asn1:"tag:2,universal"`
	LinkedId      int    `asn1:"tag:0,optional"`
	OperationCode []byte `asn1:"tag:2,universal"`
	Parameter     []byte `asn1:"tag:16,universal,optional"`
}

func (i Invoke) GetType() ComponentType {
	return InvokeOp
}

type ReturnResultLast struct {
	InvokeId int `asn1:"tag:2,universal"`
}

func (returnResultLast ReturnResultLast) GetType() ComponentType {
	return ReturnResultLastOp
}

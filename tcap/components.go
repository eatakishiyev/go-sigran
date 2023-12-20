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

/* Invoke */

type Invoke struct {
	InvokeId      int    `asn1:"tag:2,universal"`
	LinkedId      int    `asn1:"tag:0,optional"`
	OperationCode int    `asn1:"tag:2,universal"`
	Parameter     []byte `asn1:"tag:16,universal,optional"`
}

func (i *Invoke) GetType() ComponentType {
	return InvokeOp
}

/* ReturnResultLast */

type ReturnResultLast struct {
	InvokeId             int                  `asn1:"tag:2,universal"`
	ResultReturnResponse ResultReturnResponse `asn1:"tag:16,universal,optional"`
}

type ResultReturnResponse struct {
	OperationCode int    `asn1:"tag:2,universal"`
	Parameter     []byte `asn1:"tag:16,universal,optional"`
}

func (returnResultLast *ReturnResultLast) GetType() ComponentType {
	return ReturnResultLastOp
}

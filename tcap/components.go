package tcap

type Invoke struct {
	InvokeId      int `asn1:"numeric,tag:2"`
	LinkedId      int
	OperationCode []byte `asn1:"tag:6"`
	Parameter     []byte
}

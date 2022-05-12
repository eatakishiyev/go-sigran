package parameters

type Encoding uint

const (
	SingleAsn1 Encoding = iota
	OctetAligned
	Arbitrary
)

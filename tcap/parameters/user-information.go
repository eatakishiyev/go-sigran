package parameters

import (
	"encoding/asn1"
)

type Encoding uint

const (
	SingleAsn1 Encoding = iota
	OctetAligned
	Arbitrary
)

type UserInformation struct {
	DirectReference   []uint64
	IndirectReference uint64
	ObjectDescriptor  string
	Encoding          Encoding
	ExternalData      []byte
}

func (u *UserInformation) Decode(data []byte) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, u, "tag:30")
}

package parameters

import (
	"encoding/asn1"
)

type ApplicationContext struct {
	Oid asn1.ObjectIdentifier
}

func (a *ApplicationContext) Decode(data []byte) ([]byte, error) {
	return asn1.Unmarshal(data, a)
}

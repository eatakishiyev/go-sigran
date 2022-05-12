package parameters

import (
	asn12 "github.com/PromonLogicalis/asn1"
)

type ApplicationContext struct {
	Oid asn12.Oid `asn1:"tag:6,universal"`
}

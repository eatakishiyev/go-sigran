package parameters

import (
	"github.com/PromonLogicalis/asn1"
	"reflect"
)

type IndirectReference int
type ObjDescriptorType string

type External struct {
	Oid               asn1.Oid          `asn1:"tag:6,universal,optional"`
	IndirectReference IndirectReference `asn1:"tag:2,universal,optional"`
	ObjDescriptorType ObjDescriptorType `asn1:"tag:7,universal,optional"`
	Encoding          interface{}       `asn1:"choice:encoding"`
}

func (e *External) Decode(data []byte) ([]byte, error) {

	ctx := asn1.NewContext()
	ctx.SetDer(false, false)
	ctx.AddChoice("encoding", []asn1.Choice{
		{
			Type:    reflect.TypeOf([]byte{}),
			Options: "tag:0",
		},
		{
			Type:    reflect.TypeOf([]byte{}),
			Options: "tag:1",
		},
		{
			Type:    reflect.TypeOf([]byte{}),
			Options: "tag:2",
		},
	})

	return ctx.DecodeWithOptions(data, e, "tag:8,universal")
}

package parameters

import (
	"github.com/PromonLogicalis/asn1"
	"reflect"
)

type IndirectReference int
type ObjDescriptorType string

type AsnEncodingType []byte
type OctetEncodingType []byte
type ArbitraryEncodingType []byte

type External struct {
	Oid               asn1.Oid    `asn1:"tag:6,universal,optional"`
	IndirectReference []byte      `asn1:"tag:2,universal,optional"`
	ObjDescriptorType []byte      `asn1:"tag:7,universal,optional"`
	Encoding          interface{} `asn1:"choice:encoding"`
}

func (e External) Encode(params string) ([]byte, error) {
	var ctx = asn1.NewContext()
	ctx.SetDer(false, false)
	ctx.AddChoice("encoding", []asn1.Choice{
		{
			Type:    reflect.TypeOf(DialogAbortPDU{}),
			Options: "tag:0",
		},
		{
			Type:    reflect.TypeOf(OctetEncodingType{0}),
			Options: "tag:1",
		},
		{
			Type:    reflect.TypeOf(ArbitraryEncodingType{0}),
			Options: "tag:2",
		},
	})
	return ctx.EncodeWithOptions(e, params)
}

func (e External) Decode(data []byte) ([]byte, error) {
	var ctx = asn1.NewContext()
	ctx.SetDer(false, false)
	ctx.AddChoice("encoding", []asn1.Choice{
		{
			Type:    reflect.TypeOf(DialogAbortPDU{}),
			Options: "tag:0",
		},
		{
			Type:    reflect.TypeOf(OctetEncodingType{}),
			Options: "tag:1",
		},
		{
			Type:    reflect.TypeOf(ArbitraryEncodingType{}),
			Options: "tag:2",
		},
	})
	return ctx.DecodeWithOptions(data, &e, "tag:8,universal")
}

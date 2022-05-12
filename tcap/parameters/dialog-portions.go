package parameters

import (
	asn12 "encoding/asn1"
	"fmt"
	"github.com/PromonLogicalis/asn1"
	"reflect"
)

type DialogPortion struct {
	DialogPDU DialogPDU
}

func (dp *DialogPortion) Decode(data []byte) {
	var external External
	_, err := external.Decode(data)
	if err != nil {
		fmt.Printf("error occurred during parse asn.1 external type %s\n", err)
	}
	var raw asn12.RawValue
	asn12.Unmarshal(external.Encoding.([]byte), &raw)
	pduType := DialogPDUType(raw.Tag)

	var pdu DialogPDU
	var options string

	ctx := asn1.NewContext()

	switch pduType {
	case DialogRequestApdu:
		pdu = new(DialogRequestPDU)
		options = "tag:0,application"
	case DialogResponseApdu:
		pdu = new(DialogResponsePDU)

		ctx.AddChoice("associate-source-diagnostic", []asn1.Choice{
			{
				Type:    reflect.TypeOf(DialogueServiceUserDiagnostic{}),
				Options: "tag:1",
			},
			{
				Type:    reflect.TypeOf(DialogueServiceProviderDiagnostic{}),
				Options: "tag:2",
			},
		})
		options = "tag:1,application"
	case DialogAbortApdu:
		pdu = new(DialogAbortPDU)
		options = "tag:4,application"
	}
	_, err = ctx.DecodeWithOptions(raw.FullBytes, pdu, options)
	if err != nil {
		fmt.Printf("error occurred during decode dialog portion %v %s", pduType, err)
	}
	dp.DialogPDU = pdu
}

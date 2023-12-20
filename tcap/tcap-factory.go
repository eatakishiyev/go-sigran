package tcap

import (
	"fmt"
	"github.com/PromonLogicalis/asn1"
	"go-sigtran/tcap/parameters"
	"reflect"
)

type TcapMessageFactory struct {
}

func Decode(data []byte, msg interface{}) {
	ctx := asn1.NewContext()
	ctx.SetDer(false, false)
	err := ctx.AddChoice("message", []asn1.Choice{
		{
			Type:    reflect.TypeOf(BeginMessage{}),
			Options: "tag:2,application",
		},
		{
			Type:    reflect.TypeOf(EndMessage{}),
			Options: "tag:4,application",
		},
		{
			Type:    reflect.TypeOf(ContinueMessage{}),
			Options: "tag:5,application",
		},
		{
			Type:    reflect.TypeOf(AbortMessage{}),
			Options: "tag:7,application",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ctx.AddChoice("dialogPortion", []asn1.Choice{
		{
			Type:    reflect.TypeOf(parameters.DialogRequestPDU{}),
			Options: "tag:0,application",
		},
		{
			Type:    reflect.TypeOf(parameters.DialogResponsePDU{}),
			Options: "tag:1,application",
		},
		{
			Type:    reflect.TypeOf(parameters.DialogAbortPDU{}),
			Options: "tag:4,application",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ctx.AddChoice("components", []asn1.Choice{
		{
			Type:    reflect.TypeOf(Invoke{}),
			Options: "tag:1",
		},
		{
			Type:    reflect.TypeOf(ReturnResultLast{}),
			Options: "tag:2",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	err = ctx.AddChoice("associate-source-diagnostic", []asn1.Choice{
		{
			Type:    reflect.TypeOf(parameters.ServiceProviderDiagnostic(0)),
			Options: "tag:1,universal",
		},
		{
			Type:    reflect.TypeOf(parameters.ServiceUserDiagnostic(0)),
			Options: "tag:2,universal",
		},
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = ctx.DecodeWithOptions(data, msg, "choice:message")
	if err != nil {
		fmt.Println(err)
		return
	}
}

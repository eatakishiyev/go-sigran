package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type AspUpAck struct {
	*messages.MessageHeader
	*parameters.InfoString
	*parameters.AspIdentifier
}

func NewAspUpAck(infoString *parameters.InfoString, identifier *parameters.AspIdentifier) *AspUpAck {
	a := &AspUpAck{
		MessageHeader: messages.NewHeader(messages.Aspsm, messages.AspUpAck),
		InfoString:    infoString,
		AspIdentifier: identifier,
	}
	return a
}

func (a *AspUpAck) Header() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspUpAck) EncodeMessage() []byte {
	var encoded []byte
	if a.InfoString != nil {
		parameters.Encode(encoded, a.InfoString)
	}
	if a.AspIdentifier != nil {
		parameters.Encode(encoded, a.AspIdentifier)
	}
	return encoded
}

func (a *AspUpAck) DecodeMessage(b []byte) {
	if len(b) > 0 {
		params := factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			p := params[idx]
			switch p.GetHeader().Tag {
			case parameters.ParamInfoString:
				a.InfoString = p.(*parameters.InfoString)
			case parameters.ParamAspIdentifier:
				a.AspIdentifier = p.(*parameters.AspIdentifier)
			}
		}
	}
}

func (a *AspUpAck) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspUpAck) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

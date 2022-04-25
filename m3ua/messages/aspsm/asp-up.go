package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
)

type AspUp struct {
	*messages.MessageHeader
	*parameters.AspIdentifier
	*parameters.InfoString
}

func NewAspUp(aspId *parameters.AspIdentifier, infoString *parameters.InfoString) *AspUp {
	a := &AspUp{
		MessageHeader: messages.NewHeader(messages.MessageClassAspsm, messages.MessageTypeAspUp),
		InfoString:    infoString,
		AspIdentifier: aspId,
	}
	return a
}

func (a *AspUp) EncodeMessage() []byte {
	var encoded []byte
	if a.AspIdentifier != nil {
		parameters.Encode(encoded, a.AspIdentifier)
	}
	if a.InfoString != nil {
		parameters.Encode(encoded, a.InfoString)
	}
	return encoded
}

func (a *AspUp) DecodeMessage(b []byte) {
	if len(b) > 0 {
		params := parameters.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			p := params[idx]
			switch p.GetHeader().Tag {
			case parameters.ParamAspIdentifier:
				a.AspIdentifier = p.(*parameters.AspIdentifier)
			case parameters.ParamInfoString:
				a.InfoString = p.(*parameters.InfoString)
			}
		}
	}
}

func (a *AspUp) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspUp) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

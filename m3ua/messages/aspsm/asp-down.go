package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
)

type AspDown struct {
	*messages.MessageHeader
	infoString *parameters.InfoString
}

func NewAspDown(infoString *parameters.InfoString) *AspDown {
	return &AspDown{
		MessageHeader: messages.NewHeader(messages.Aspsm, messages.AspDn),
		infoString:    infoString,
	}
}

func (a *AspDown) EncodeMessage() []byte {
	var encoded []byte
	if a.infoString != nil {
		parameters.Encode(encoded, a.infoString)
	}
	return encoded
}

func (a *AspDown) DecodeMessage(b []byte) {
	availableBytes := len(b)
	if availableBytes > 0 {
		a.infoString = parameters.NewInfoString(string(b))
	}
}

func (a *AspDown) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspDown) InfoString() *parameters.InfoString {
	return a.infoString
}

func (a *AspDown) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

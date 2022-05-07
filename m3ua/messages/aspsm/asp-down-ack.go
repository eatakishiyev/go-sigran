package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
)

type AspDownAck struct {
	*messages.MessageHeader
	infoString *parameters.InfoString
}

func NewAspDownAck(infoString *parameters.InfoString) *AspDownAck {
	return &AspDownAck{
		MessageHeader: messages.NewHeader(messages.Aspsm, messages.AspDnAck),
		infoString:    infoString,
	}
}

func (a *AspDownAck) EncodeMessage() []byte {
	var encoded []byte
	if a.infoString != nil {
		parameters.Encode(encoded, a.infoString)
	}
	return encoded
}

func (a *AspDownAck) DecodeMessage(b []byte) {
	availableBytes := len(b)
	if availableBytes > 0 {
		a.infoString = parameters.NewInfoString(string(b))
	}
}

func (a *AspDownAck) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspDownAck) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

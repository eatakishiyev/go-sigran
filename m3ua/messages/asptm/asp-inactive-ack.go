package asptm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
)

type AspInactiveAck struct {
	*messages.MessageHeader
	*parameters.RoutingContext
	*parameters.InfoString
}

func NewAspInactiveAck(routingContext *parameters.RoutingContext, infoString *parameters.InfoString) *AspInactiveAck {
	return &AspInactiveAck{
		MessageHeader: &messages.MessageHeader{
			MessageType:  messages.MessageTypeAspIaAck,
			MessageClass: messages.MessageClassAsptm,
		},
		RoutingContext: routingContext,
		InfoString:     infoString,
	}
}

func (a *AspInactiveAck) EncodeMessage() []byte {
	var encoded []byte
	if a.RoutingContext != nil {
		parameters.Encode(encoded, a.RoutingContext)
	}

	if a.InfoString != nil {
		parameters.Encode(encoded, a.InfoString)
	}

	return encoded
}

func (a *AspInactiveAck) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = parameters.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamRoutingContext:
				a.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamInfoString:
				a.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (a *AspInactiveAck) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspInactiveAck) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

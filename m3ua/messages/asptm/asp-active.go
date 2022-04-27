package asptm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type AspActive struct {
	*messages.MessageHeader
	*parameters.TrafficMode
	*parameters.RoutingContext
	*parameters.InfoString
}

func NewAspActive(trafficMode *parameters.TrafficMode, routingContext *parameters.RoutingContext, infoString *parameters.InfoString) *AspActive {
	return &AspActive{
		MessageHeader: &messages.MessageHeader{
			MessageType:  messages.MessageTypeAspAc,
			MessageClass: messages.MessageClassAsptm,
		},
		TrafficMode:    trafficMode,
		RoutingContext: routingContext,
		InfoString:     infoString,
	}
}

func (a *AspActive) EncodeMessage() []byte {
	var encoded []byte
	if a.TrafficMode != nil {
		parameters.Encode(encoded, a.TrafficMode)
	}
	if a.RoutingContext != nil {
		parameters.Encode(encoded, a.RoutingContext)
	}
	if a.InfoString != nil {
		parameters.Encode(encoded, a.InfoString)
	}
	return encoded
}

func (a *AspActive) DecodeMessage(b []byte) {
	if len(b) > 0 {
		params := factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			p := params[idx]
			switch p.GetHeader().Tag {
			case parameters.ParamTrafficMode:
				a.TrafficMode = p.(*parameters.TrafficMode)
			case parameters.ParamRoutingContext:
				a.RoutingContext = p.(*parameters.RoutingContext)
			case parameters.ParamInfoString:
				a.InfoString = p.(*parameters.InfoString)
			}
		}
	}
}

func (a *AspActive) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspActive) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

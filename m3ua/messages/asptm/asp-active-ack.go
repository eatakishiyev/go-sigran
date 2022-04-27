package asptm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type AspActiveAck struct {
	*messages.MessageHeader
	*parameters.TrafficMode
	*parameters.RoutingContext
	*parameters.InfoString
}

func NewAspActiveAck(trafficMode *parameters.TrafficMode, routingContext *parameters.RoutingContext, infoString *parameters.InfoString) *AspActiveAck {
	return &AspActiveAck{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassAsptm,
			MessageType:  messages.MessageTypeAspAcAck,
		},
		TrafficMode:    trafficMode,
		RoutingContext: routingContext,
		InfoString:     infoString,
	}
}

func (a *AspActiveAck) EncodeMessage() []byte {
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

func (a *AspActiveAck) DecodeMessage(b []byte) {
	if len(b) > 0 {
		params := factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			param := params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamTrafficMode:
				a.TrafficMode = param.(*parameters.TrafficMode)
			case parameters.ParamRoutingContext:
				a.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamInfoString:
				a.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (a *AspActiveAck) GetHeader() *messages.MessageHeader {
	return a.MessageHeader
}

func (a *AspActiveAck) SetHeader(header *messages.MessageHeader) {
	a.MessageHeader = header
}

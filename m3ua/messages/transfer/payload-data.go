package transfer

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type PayloadData struct {
	*messages.MessageHeader
	*parameters.NetworkAppearance
	*parameters.RoutingContext
	*parameters.ProtocolData
	*parameters.CorrelationId
}

func (p *PayloadData) EncodeMessage() []byte {
	var encoded []byte
	if p.NetworkAppearance != nil {
		parameters.Encode(encoded, p.NetworkAppearance)
	}

	if p.RoutingContext != nil {
		parameters.Encode(encoded, p.RoutingContext)
	}

	if p.ProtocolData != nil {
		parameters.Encode(encoded, p.ProtocolData)
	}
	return encoded
}

func (p *PayloadData) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamNetworkAppearance:
				p.NetworkAppearance = param.(*parameters.NetworkAppearance)
			case parameters.ParamRoutingContext:
				p.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamProtocolData:
				p.ProtocolData = param.(*parameters.ProtocolData)
			}
		}
	}
}

func (p *PayloadData) GetHeader() *messages.MessageHeader {
	return p.MessageHeader
}

func (p *PayloadData) SetHeader(header *messages.MessageHeader) {
	p.MessageHeader = header
}

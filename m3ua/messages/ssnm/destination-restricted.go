package ssnm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type DestinationRestricted struct {
	*messages.MessageHeader
	*parameters.RoutingContext
	*parameters.AffectedPointCode
	*parameters.InfoString
}

func NewDestinationRestricted(routingContext *parameters.RoutingContext, affectedPointCode *parameters.AffectedPointCode, infoString *parameters.InfoString) *DestinationRestricted {
	return &DestinationRestricted{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassSsnm,
			MessageType:  messages.MessageTypeDrst,
		},
		RoutingContext:    routingContext,
		AffectedPointCode: affectedPointCode,
		InfoString:        infoString,
	}
}

func (d *DestinationRestricted) EncodeMessage() []byte {
	var encoded []byte
	if d.RoutingContext != nil {
		parameters.Encode(encoded, d.RoutingContext)
	}

	if d.AffectedPointCode != nil {
		parameters.Encode(encoded, d.AffectedPointCode)
	}

	if d.InfoString != nil {
		parameters.Encode(encoded, d.InfoString)
	}
	return encoded
}

func (d *DestinationRestricted) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamRoutingContext:
				d.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamAffectedPointCode:
				d.AffectedPointCode = param.(*parameters.AffectedPointCode)
			case parameters.ParamInfoString:
				d.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (d *DestinationRestricted) GetHeader() *messages.MessageHeader {
	return d.MessageHeader
}

func (d *DestinationRestricted) SetHeader(header *messages.MessageHeader) {
	d.MessageHeader = header
}

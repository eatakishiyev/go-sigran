package ssnm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type DestinationUnavailable struct {
	*messages.MessageHeader
	*parameters.NetworkAppearance
	*parameters.RoutingContext
	*parameters.AffectedPointCode
	*parameters.InfoString
}

func NewDestinationUnavailable(networkAppearance *parameters.NetworkAppearance, routingContext *parameters.RoutingContext, affectedPointCode *parameters.AffectedPointCode, infoString *parameters.InfoString) *DestinationUnavailable {
	return &DestinationUnavailable{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassSsnm,
			MessageType:  messages.MessageTypeDuna,
		},
		NetworkAppearance: networkAppearance,
		RoutingContext:    routingContext,
		AffectedPointCode: affectedPointCode,
		InfoString:        infoString,
	}
}

func (d *DestinationUnavailable) EncodeMessage() []byte {
	var encoded []byte
	if d.NetworkAppearance != nil {
		parameters.Encode(encoded, d.NetworkAppearance)
	}

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

func (d *DestinationUnavailable) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamNetworkAppearance:
				d.NetworkAppearance = param.(*parameters.NetworkAppearance)
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

func (d *DestinationUnavailable) GetHeader() *messages.MessageHeader {
	return d.MessageHeader
}

func (d *DestinationUnavailable) SetHeader(header *messages.MessageHeader) {
	d.MessageHeader = header
}

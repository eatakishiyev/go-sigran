package ssnm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type DestinationStateAudit struct {
	*messages.MessageHeader
	*parameters.NetworkAppearance
	*parameters.AffectedPointCode
	*parameters.InfoString
}

func NewDestinationStateAudit(networkAppearance *parameters.NetworkAppearance, affectedPointCode *parameters.AffectedPointCode, infoString *parameters.InfoString) *DestinationStateAudit {
	return &DestinationStateAudit{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.Ssnm,
			MessageType:  messages.Daud,
		},
		NetworkAppearance: networkAppearance,
		AffectedPointCode: affectedPointCode,
		InfoString:        infoString,
	}
}

func (d *DestinationStateAudit) EncodeMessage() []byte {
	var encoded []byte
	if d.NetworkAppearance != nil {
		parameters.Encode(encoded, d.NetworkAppearance)
	}

	if d.AffectedPointCode != nil {
		parameters.Encode(encoded, d.AffectedPointCode)
	}

	if d.InfoString != nil {
		parameters.Encode(encoded, d.InfoString)
	}
	return encoded
}

func (d *DestinationStateAudit) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamNetworkAppearance:
				d.NetworkAppearance = param.(*parameters.NetworkAppearance)
			case parameters.ParamAffectedPointCode:
				d.AffectedPointCode = param.(*parameters.AffectedPointCode)
			case parameters.ParamInfoString:
				d.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (d *DestinationStateAudit) GetHeader() *messages.MessageHeader {
	return d.MessageHeader
}

func (d *DestinationStateAudit) SetHeader(header *messages.MessageHeader) {
	d.MessageHeader = header
}

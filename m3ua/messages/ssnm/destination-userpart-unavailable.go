package ssnm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type DestinationUserPartUnavailable struct {
	*messages.MessageHeader
	*parameters.NetworkAppearance
	*parameters.AffectedPointCode
	*parameters.UserPartUnavailableCause
	*parameters.InfoString
}

func (d *DestinationUserPartUnavailable) EncodeMessage() []byte {
	var encoded []byte

	if d.NetworkAppearance != nil {
		parameters.Encode(encoded, d.NetworkAppearance)
	}

	if d.AffectedPointCode != nil {
		parameters.Encode(encoded, d.AffectedPointCode)
	}

	if d.UserPartUnavailableCause != nil {
		parameters.Encode(encoded, d.UserPartUnavailableCause)
	}

	if d.InfoString != nil {
		parameters.Encode(encoded, d.InfoString)
	}

	return encoded
}

func (d *DestinationUserPartUnavailable) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamNetworkAppearance:
				d.NetworkAppearance = param.(*parameters.NetworkAppearance)
			case parameters.ParamAffectedPointCode:
				d.AffectedPointCode = param.(*parameters.AffectedPointCode)
			case parameters.ParamUserCause:
				d.UserPartUnavailableCause = param.(*parameters.UserPartUnavailableCause)
			case parameters.ParamInfoString:
				d.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (d *DestinationUserPartUnavailable) GetHeader() *messages.MessageHeader {
	return d.MessageHeader
}

func (d *DestinationUserPartUnavailable) SetHeader(header *messages.MessageHeader) {
	d.MessageHeader = header
}

func NewDestinationUserPartUnavailable(networkAppearance *parameters.NetworkAppearance, affectedPointCode *parameters.AffectedPointCode, userPartUnavailableCause *parameters.UserPartUnavailableCause, infoString *parameters.InfoString) *DestinationUserPartUnavailable {
	return &DestinationUserPartUnavailable{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassSsnm,
			MessageType:  messages.MessageTypeDupu,
		},
		NetworkAppearance:        networkAppearance,
		AffectedPointCode:        affectedPointCode,
		UserPartUnavailableCause: userPartUnavailableCause,
		InfoString:               infoString,
	}
}

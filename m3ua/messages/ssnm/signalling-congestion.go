package ssnm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type SignallingCongestion struct {
	*messages.MessageHeader
	*parameters.NetworkAppearance
	*parameters.RoutingContext
	*parameters.AffectedPointCode
	*parameters.ConcernedDpc
	*parameters.CongestionIndication
	*parameters.InfoString
}

func (s *SignallingCongestion) EncodeMessage() []byte {
	var encoded []byte
	if s.NetworkAppearance != nil {
		parameters.Encode(encoded, s.NetworkAppearance)
	}

	if s.RoutingContext != nil {
		parameters.Encode(encoded, s.RoutingContext)
	}

	if s.AffectedPointCode != nil {
		parameters.Encode(encoded, s.AffectedPointCode)
	}

	if s.ConcernedDpc != nil {
		parameters.Encode(encoded, s.ConcernedDpc)
	}

	if s.CongestionIndication != nil {
		parameters.Encode(encoded, s.CongestionIndication)
	}

	if s.InfoString != nil {
		parameters.Encode(encoded, s.InfoString)
	}
	return encoded
}

func (s *SignallingCongestion) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamNetworkAppearance:
				s.NetworkAppearance = param.(*parameters.NetworkAppearance)
			case parameters.ParamRoutingContext:
				s.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamAffectedPointCode:
				s.AffectedPointCode = param.(*parameters.AffectedPointCode)
			case parameters.ParamConcernedDestinations:
				s.ConcernedDpc = param.(*parameters.ConcernedDpc)
			case parameters.ParamCongestionIndications:
				s.CongestionIndication = param.(*parameters.CongestionIndication)
			case parameters.ParamInfoString:
				s.InfoString = param.(*parameters.InfoString)
			}
		}
	}
}

func (s *SignallingCongestion) GetHeader() *messages.MessageHeader {
	return s.MessageHeader
}

func (s *SignallingCongestion) SetHeader(header *messages.MessageHeader) {
	s.MessageHeader = header
}

func NewSignallingCongestion(networkAppearance *parameters.NetworkAppearance, routingContext *parameters.RoutingContext,
	affectedPointCode *parameters.AffectedPointCode, concernedDpc *parameters.ConcernedDpc,
	congestionIndication *parameters.CongestionIndication, infoString *parameters.InfoString) *SignallingCongestion {
	return &SignallingCongestion{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassSsnm,
			MessageType:  messages.MessageTypeScon,
		},
		NetworkAppearance:    networkAppearance,
		RoutingContext:       routingContext,
		AffectedPointCode:    affectedPointCode,
		ConcernedDpc:         concernedDpc,
		CongestionIndication: congestionIndication,
		InfoString:           infoString}
}

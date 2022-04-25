package mgmt

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
)

type Error struct {
	*messages.MessageHeader
	*parameters.RoutingContext
	*parameters.AffectedPointCode
	*parameters.NetworkAppearance
	*parameters.DiagnosticInformation
}

func NewError(routingContext *parameters.RoutingContext, affectedPointCode *parameters.AffectedPointCode, networkAppearance *parameters.NetworkAppearance, diagnosticInformation *parameters.DiagnosticInformation) *Error {
	return &Error{
		MessageHeader: &messages.MessageHeader{
			MessageClass: messages.MessageClassMgmt,
			MessageType:  messages.MessageTypeErr,
		},
		RoutingContext:        routingContext,
		AffectedPointCode:     affectedPointCode,
		NetworkAppearance:     networkAppearance,
		DiagnosticInformation: diagnosticInformation,
	}
}

func (e *Error) EncodeMessage() []byte {
	var encoded []byte
	if e.RoutingContext != nil {
		parameters.Encode(encoded, e.RoutingContext)
	}

	if e.AffectedPointCode != nil {
		parameters.Encode(encoded, e.AffectedPointCode)
	}

	if e.NetworkAppearance != nil {
		parameters.Encode(encoded, e.NetworkAppearance)
	}

	if e.DiagnosticInformation != nil {
		parameters.Encode(encoded, e.DiagnosticInformation)
	}
	return encoded
}

func (e *Error) DecodeMessage(b []byte) {
	if len(b) > 0 {
		var params = parameters.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			var param = params[idx]
			switch param.GetHeader().Tag {
			case parameters.ParamRoutingContext:
				e.RoutingContext = param.(*parameters.RoutingContext)
			case parameters.ParamAffectedPointCode:
				e.AffectedPointCode = param.(*parameters.AffectedPointCode)
			case parameters.ParamNetworkAppearance:
				e.NetworkAppearance = param.(*parameters.NetworkAppearance)
			case parameters.ParamDiagnosticInformation:
				e.DiagnosticInformation = param.(*parameters.DiagnosticInformation)
			}
		}
	}
}

func (e *Error) GetHeader() *messages.MessageHeader {
	return e.MessageHeader
}

func (e *Error) SetHeader(header *messages.MessageHeader) {
	e.MessageHeader = header
}

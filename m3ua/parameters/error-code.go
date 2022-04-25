package parameters

import "encoding/binary"

const (
	InvalidVersion            = 0x01
	UnsupportedMessageClass   = 0x03
	UnsupportedMessageType    = 0x04
	UnsupportedTrafficMode    = 0x05
	UnexpectedMessage         = 0x06
	ProtocolError             = 0x07
	InvalidStreamIdentifier   = 0x09
	RefusedManagementBlocking = 0x0d
	AspIdentifierRequired     = 0x0e
	InvalidAspIdentifier      = 0x0f
	InvalidParameterValue     = 0x11
	ParameterFieldError       = 0x12
	UnexpectedParameter       = 0x13
	DestinationStatusUnknown  = 0x14
	InvalidNetworkAppearance  = 0x15
	MissingParameter          = 0x16
	InvalidRoutingContext     = 0x19
	NoConfiguredAsForAsp      = 0x1a
)

type ErrorCode struct {
	*ParameterHeader
	Code uint32
}

func NewErrorCode(code uint32) *ErrorCode {
	return &ErrorCode{
		Code: code,
		ParameterHeader: &ParameterHeader{
			Tag: ParamErrorCode,
		},
	}
}

func (e *ErrorCode) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, e.Code)
	return encoded
}

func (e *ErrorCode) DecodeParameter(p []byte) {
	e.Code = binary.BigEndian.Uint32(p)
}

func (e *ErrorCode) GetHeader() *ParameterHeader {
	return e.ParameterHeader
}

func (e *ErrorCode) SetHeader(header *ParameterHeader) {
	e.ParameterHeader = header
}

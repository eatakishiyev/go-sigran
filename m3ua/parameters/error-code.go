package parameters

import "encoding/binary"

type ErrorCode struct {
	*ParameterHeader
	Code M3UAError
}

func NewErrorCode(code M3UAError) *ErrorCode {
	return &ErrorCode{
		Code: code,
		ParameterHeader: &ParameterHeader{
			Tag: ParamErrorCode,
		},
	}
}

func (e *ErrorCode) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, uint32(e.Code))
	return encoded
}

func (e *ErrorCode) DecodeParameter(p []byte) {
	e.Code = M3UAError(binary.BigEndian.Uint32(p))
}

func (e *ErrorCode) GetHeader() *ParameterHeader {
	return e.ParameterHeader
}

func (e *ErrorCode) SetHeader(header *ParameterHeader) {
	e.ParameterHeader = header
}

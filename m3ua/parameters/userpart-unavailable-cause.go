package parameters

import "encoding/binary"

type UserPartUnavailableCause struct {
	*ParameterHeader
	Cause            Cause
	ServiceIndicator ServiceIndicator
}

func (u *UserPartUnavailableCause) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint16(encoded, uint16(u.Cause))
	binary.BigEndian.PutUint16(encoded, uint16(u.ServiceIndicator))
	return encoded
}

func (u *UserPartUnavailableCause) DecodeParameter(p []byte) {
	u.Cause = Cause(binary.BigEndian.Uint16(p[0:2]))
	u.ServiceIndicator = ServiceIndicator(binary.BigEndian.Uint16(p[2:4]))
}

func (u *UserPartUnavailableCause) GetHeader() *ParameterHeader {
	return u.ParameterHeader
}

func (u *UserPartUnavailableCause) SetHeader(header *ParameterHeader) {
	u.ParameterHeader = header
}

func NewUserPartUnavailableCause(cause Cause, serviceIndicator ServiceIndicator) *UserPartUnavailableCause {
	return &UserPartUnavailableCause{
		ParameterHeader: &ParameterHeader{
			Tag: ParamUserCause,
		},
		Cause:            cause,
		ServiceIndicator: serviceIndicator,
	}
}

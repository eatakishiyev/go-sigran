package parameters

import "encoding/binary"

const (
	CauseUnknown = iota
	CauseUnequippedRemoteUser
	CauseInaccessibleRemoteUser
)

type UserPartUnavailableCause struct {
	*ParameterHeader
	Cause                uint16
	ServiceIdentificator uint16
}

func (u *UserPartUnavailableCause) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint16(encoded, u.Cause)
	binary.BigEndian.PutUint16(encoded, u.ServiceIdentificator)
	return encoded
}

func (u *UserPartUnavailableCause) DecodeParameter(p []byte) {
	u.Cause = binary.BigEndian.Uint16(p[0:2])
	u.ServiceIdentificator = binary.BigEndian.Uint16(p[2:4])
}

func (u *UserPartUnavailableCause) GetHeader() *ParameterHeader {
	return u.ParameterHeader
}

func (u *UserPartUnavailableCause) SetHeader(header *ParameterHeader) {
	u.ParameterHeader = header
}

func NewUserPartUnavailableCause(cause uint16, serviceIdentificator uint16) *UserPartUnavailableCause {
	return &UserPartUnavailableCause{
		ParameterHeader: &ParameterHeader{
			Tag: ParamUserCause,
		},
		Cause:                cause,
		ServiceIdentificator: serviceIdentificator,
	}
}

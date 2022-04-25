package parameters

import "encoding/binary"

type AspIdentifier struct {
	*ParameterHeader
	Id uint32
}

func NewAspIdentifier(id uint32) *AspIdentifier {
	return &AspIdentifier{
		Id:              id,
		ParameterHeader: &ParameterHeader{Tag: ParamAspIdentifier},
	}
}

func (a *AspIdentifier) EncodeParameter() []byte {
	var data []byte
	binary.BigEndian.PutUint32(data, a.Id)
	return data
}

func (a *AspIdentifier) DecodeParameter(b []byte) {
	a.Id = binary.BigEndian.Uint32(b)
}

func (a *AspIdentifier) GetHeader() *ParameterHeader {
	return a.ParameterHeader
}

func (a *AspIdentifier) SetHeader(header *ParameterHeader) {
	a.ParameterHeader = header
}

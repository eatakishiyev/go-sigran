package parameters

import "encoding/binary"

type TrafficMode struct {
	*ParameterHeader
	Mode uint32
}

func NewTrafficMode(mode uint32) *TrafficMode {
	return &TrafficMode{
		ParameterHeader: &ParameterHeader{
			Tag: ParamTrafficMode,
		},
		Mode: mode,
	}
}

func (t *TrafficMode) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, t.Mode)
	return encoded
}

func (t *TrafficMode) DecodeParameter(p []byte) {
	t.Mode = binary.BigEndian.Uint32(p)
}

func (t *TrafficMode) GetHeader() *ParameterHeader {
	return t.ParameterHeader
}

func (t *TrafficMode) SetHeader(header *ParameterHeader) {
	t.ParameterHeader = header
}

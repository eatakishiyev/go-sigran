package parameters

import "encoding/binary"

type TrafficMode struct {
	*ParameterHeader
	Mode Mode
}

func NewTrafficMode(mode Mode) *TrafficMode {
	return &TrafficMode{
		ParameterHeader: &ParameterHeader{
			Tag: ParamTrafficMode,
		},
		Mode: mode,
	}
}

func (t *TrafficMode) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, uint32(t.Mode))
	return encoded
}

func (t *TrafficMode) DecodeParameter(p []byte) {
	t.Mode = Mode(binary.BigEndian.Uint32(p))
}

func (t *TrafficMode) GetHeader() *ParameterHeader {
	return t.ParameterHeader
}

func (t *TrafficMode) SetHeader(header *ParameterHeader) {
	t.ParameterHeader = header
}

package parameters

import "encoding/binary"

type CongestionIndication struct {
	*ParameterHeader
	CongestionLevel Level
}

func (c *CongestionIndication) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, uint32(c.CongestionLevel))
	return encoded
}

func (c *CongestionIndication) DecodeParameter(p []byte) {
	c.CongestionLevel = Level(binary.BigEndian.Uint32(p))
}

func (c *CongestionIndication) GetHeader() *ParameterHeader {
	return c.ParameterHeader
}

func (c *CongestionIndication) SetHeader(header *ParameterHeader) {
	c.ParameterHeader = header
}

func NewCongestionIndication(congestionLevel Level) *CongestionIndication {
	return &CongestionIndication{
		ParameterHeader: &ParameterHeader{
			Tag: ParamCongestionIndications,
		},
		CongestionLevel: congestionLevel,
	}
}

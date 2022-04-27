package parameters

import "encoding/binary"

const (
	CongestionLevelNoCongestion = iota
	CongestionLevel1
	CongestionLevel2
	CongestionLevel3
)

type CongestionIndication struct {
	*ParameterHeader
	CongestionLevel uint32
}

func (c *CongestionIndication) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, c.CongestionLevel)
	return encoded
}

func (c *CongestionIndication) DecodeParameter(p []byte) {
	c.CongestionLevel = binary.BigEndian.Uint32(p)
}

func (c *CongestionIndication) GetHeader() *ParameterHeader {
	return c.ParameterHeader
}

func (c *CongestionIndication) SetHeader(header *ParameterHeader) {
	c.ParameterHeader = header
}

func NewCongestionIndication(congestionLevel uint32) *CongestionIndication {
	return &CongestionIndication{
		ParameterHeader: &ParameterHeader{
			Tag: ParamCongestionIndications,
		},
		CongestionLevel: congestionLevel,
	}
}

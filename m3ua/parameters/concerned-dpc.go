package parameters

import "encoding/binary"

type ConcernedDpc struct {
	*ParameterHeader
	PointCode uint32
}

func (c *ConcernedDpc) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, c.PointCode)
	return encoded
}

func (c *ConcernedDpc) DecodeParameter(p []byte) {
	c.PointCode = binary.BigEndian.Uint32(p)
}

func (c *ConcernedDpc) GetHeader() *ParameterHeader {
	return c.ParameterHeader
}

func (c *ConcernedDpc) SetHeader(header *ParameterHeader) {
	c.ParameterHeader = header
}

func NewConcernedDpc(pointCode uint32) *ConcernedDpc {
	return &ConcernedDpc{
		ParameterHeader: &ParameterHeader{
			Tag: ParamConcernedDestinations,
		},
		PointCode: pointCode,
	}
}

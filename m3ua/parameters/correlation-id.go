package parameters

import "encoding/binary"

type CorrelationId struct {
	*ParameterHeader
	CorrelationId uint32
}

func (c *CorrelationId) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, c.CorrelationId)
	return encoded
}

func (c *CorrelationId) DecodeParameter(p []byte) {
	c.CorrelationId = binary.BigEndian.Uint32(p)
}

func (c *CorrelationId) GetHeader() *ParameterHeader {
	return c.ParameterHeader
}

func (c *CorrelationId) SetHeader(header *ParameterHeader) {
	c.ParameterHeader = header
}

func NewCorrelationId(correlationId uint32) *CorrelationId {
	return &CorrelationId{ParameterHeader: &ParameterHeader{
		Tag: ParamCorrelationId,
	},
		CorrelationId: correlationId,
	}
}

package parameters

import (
	"bytes"
	"encoding/binary"
)

type RoutingContext struct {
	*ParameterHeader
	RoutingContexts []uint32
}

func (rcontext *RoutingContext) EncodeParameter() []byte {
	var encoded []byte
	for idx := 0; idx < len(rcontext.RoutingContexts); idx++ {
		rc := rcontext.RoutingContexts[idx]
		binary.BigEndian.PutUint32(encoded, rc)
	}
	return encoded
}

func (rcontext *RoutingContext) DecodeParameter(b []byte) {
	if len(b) > 0 {
		r := bytes.NewReader(b)
		for {
			rcData := make([]byte, 4)
			r.Read(rcData)
			rc := binary.BigEndian.Uint32(rcData)
			rcontext.RoutingContexts = append(rcontext.RoutingContexts, rc)
			if r.Len() <= 0 {
				break
			}
		}
	}
}

func (rcontext *RoutingContext) GetHeader() *ParameterHeader {
	return rcontext.ParameterHeader
}

func (rcontext *RoutingContext) SetHeader(header *ParameterHeader) {
	rcontext.ParameterHeader = header
}

func NewRoutingContext(routingContexts []uint32) *RoutingContext {
	return &RoutingContext{
		ParameterHeader: &ParameterHeader{
			Tag: ParamRoutingContext,
		},
		RoutingContexts: routingContexts,
	}
}

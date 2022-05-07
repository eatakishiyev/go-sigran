package parameters

import (
	"bytes"
	"encoding/binary"
)

type RoutingContext struct {
	*ParameterHeader
	RoutingContexts []uint32
}

func (rc *RoutingContext) EncodeParameter() []byte {
	var encoded []byte
	for idx := 0; idx < len(rc.RoutingContexts); idx++ {
		rc := rc.RoutingContexts[idx]
		binary.BigEndian.PutUint32(encoded, rc)
	}
	return encoded
}

func (rc *RoutingContext) DecodeParameter(b []byte) {
	if len(b) > 0 {
		r := bytes.NewReader(b)
		for {
			rcData := make([]byte, 4)
			r.Read(rcData)
			rContext := binary.BigEndian.Uint32(rcData)
			rc.RoutingContexts = append(rc.RoutingContexts, rContext)
			if r.Len() <= 0 {
				break
			}
		}
	}
}

func (rc *RoutingContext) GetHeader() *ParameterHeader {
	return rc.ParameterHeader
}

func (rc *RoutingContext) SetHeader(header *ParameterHeader) {
	rc.ParameterHeader = header
}

func NewRoutingContext(routingContexts []uint32) *RoutingContext {
	return &RoutingContext{
		ParameterHeader: &ParameterHeader{
			Tag: ParamRoutingContext,
		},
		RoutingContexts: routingContexts,
	}
}

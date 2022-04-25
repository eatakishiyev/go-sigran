package parameters

import "encoding/binary"

type NetworkAppearance struct {
	*ParameterHeader
	Appearance uint32
}

func NewNetworkAppearance(appearance uint32) *NetworkAppearance {
	return &NetworkAppearance{
		Appearance:      appearance,
		ParameterHeader: &ParameterHeader{Tag: ParamNetworkAppearance},
	}
}

func (n *NetworkAppearance) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, n.Appearance)
	return encoded
}

func (n *NetworkAppearance) DecodeParameter(p []byte) {
	n.Appearance = binary.BigEndian.Uint32(p)
}

func (n *NetworkAppearance) GetHeader() *ParameterHeader {
	return n.ParameterHeader
}

func (n *NetworkAppearance) SetHeader(header *ParameterHeader) {
	n.ParameterHeader = header
}

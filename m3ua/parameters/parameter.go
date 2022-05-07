package parameters

import (
	"bytes"
	"encoding/binary"
)

type Parameter interface {
	EncodeParameter() []byte
	DecodeParameter(p []byte)
	GetHeader() *ParameterHeader
	SetHeader(*ParameterHeader)
}

func Encode(slice []byte, p Parameter) {
	var encoded []byte
	pData := p.EncodeParameter()
	binary.BigEndian.PutUint16(encoded, uint16(p.GetHeader().Tag))
	binary.BigEndian.PutUint16(encoded, uint16(len(pData)+4)) //parameter data length + 2 tag + 2 length
	encoded = append(encoded, pData...)
	//write padding bytes
	paddingBytesCount := 4 - len(pData)%4
	paddings := make([]byte, paddingBytesCount)
	encoded = append(encoded, paddings...)
	slice = append(slice, encoded...)
}

func DecodeHeader(r *bytes.Reader) *ParameterHeader {
	headerOctets := make([]byte, 4)
	r.Read(headerOctets)

	tag := ParameterTag(binary.BigEndian.Uint16(headerOctets[0:2]))
	length := binary.BigEndian.Uint16(headerOctets[2:4])
	return NewHeader(tag, length)
}

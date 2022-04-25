package parameters

import (
	"bytes"
)

type AffectedPointCode struct {
	*ParameterHeader
	PointCodes []PointCode
}

func NewAffectedPointCode(pointCodes []PointCode) *AffectedPointCode {
	return &AffectedPointCode{
		PointCodes: pointCodes,
		ParameterHeader: &ParameterHeader{
			Tag: ParamAffectedPointCode,
		},
	}
}

func (a *AffectedPointCode) EncodeParameter() []byte {
	var encoded []byte
	for _, pc := range a.PointCodes {
		encoded = append(encoded, pc.Mask)
		encoded = append(encoded, pc.Mask>>16)
		encoded = append(encoded, pc.Mask>>8)
		encoded = append(encoded, pc.Mask)
	}
	return encoded
}

func (a *AffectedPointCode) DecodeParameter(p []byte) {
	r := bytes.NewReader(p)
	for {
		octets := make([]byte, 4)
		r.Read(octets)

		mask := octets[0] & 0xff

		pointCode := octets[1] & 0xff
		pointCode = (pointCode << 8) | (octets[2] & 0xff)
		pointCode = (pointCode << 8) | (octets[3] & 0xff)

		p := PointCode{
			PointCode: uint32(pointCode),
			Mask:      mask,
		}
		a.PointCodes = append(a.PointCodes, p)

		if r.Len() <= 0 {
			break
		}
	}
}

func (a *AffectedPointCode) GetHeader() *ParameterHeader {
	return a.ParameterHeader
}

func (a *AffectedPointCode) SetHeader(header *ParameterHeader) {
	a.ParameterHeader = header
}

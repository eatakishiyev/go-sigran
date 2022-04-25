package parameters

import (
	"bytes"
	"math"
)

func DecodeParameters(b []byte) []Parameter {
	var params []Parameter
	r := bytes.NewReader(b)
	for {
		header := DecodeHeader(r)

		parameterPayloadOctetCount := header.Length - 4
		parameterPayload := make([]byte, parameterPayloadOctetCount)
		r.Read(parameterPayload)
		paddedOctetsCount := math.Max(0, float64(parameterPayloadOctetCount%4))

		if paddedOctetsCount > 0 {
			paddedOctets := make([]byte, parameterPayloadOctetCount)
			r.Read(paddedOctets)
		}

		var p Parameter

		switch header.Tag {
		case ParamAspIdentifier:
			p = &AspIdentifier{}
		case ParamAffectedPointCode:
			p = &AffectedPointCode{}
		case ParamInfoString:
			p = &InfoString{}
		case ParamHeartBeadData:
			p = &HeartbeatData{}
		case ParamRoutingContext:
			p = &RoutingContext{}
		case ParamTrafficMode:
			p = &TrafficMode{}
		case ParamErrorCode:
			p = &ErrorCode{}
		case ParamNetworkAppearance:
			p = &NetworkAppearance{}
		case ParamDiagnosticInformation:
			p = &DiagnosticInformation{}
		}

		if p != nil {
			p.SetHeader(header)
			p.DecodeParameter(parameterPayload)
			params = append(params, p)
		}

		if r.Len() <= 0 {
			break
		}
	}
	return params
}

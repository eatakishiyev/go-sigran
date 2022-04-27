package factory

import (
	"bytes"
	"go-sigtran/m3ua/parameters"
	"math"
)

func DecodeParameters(b []byte) []parameters.Parameter {
	var params []parameters.Parameter
	r := bytes.NewReader(b)
	for {
		header := parameters.DecodeHeader(r)

		parameterPayloadOctetCount := header.Length - 4
		parameterPayload := make([]byte, parameterPayloadOctetCount)
		r.Read(parameterPayload)
		paddedOctetsCount := math.Max(0, float64(parameterPayloadOctetCount%4))

		if paddedOctetsCount > 0 {
			paddedOctets := make([]byte, parameterPayloadOctetCount)
			r.Read(paddedOctets)
		}

		var p parameters.Parameter

		switch header.Tag {
		case parameters.ParamAspIdentifier:
			p = &parameters.AspIdentifier{}
		case parameters.ParamAffectedPointCode:
			p = &parameters.AffectedPointCode{}
		case parameters.ParamInfoString:
			p = &parameters.InfoString{}
		case parameters.ParamHeartBeadData:
			p = &parameters.HeartbeatData{}
		case parameters.ParamRoutingContext:
			p = &parameters.RoutingContext{}
		case parameters.ParamTrafficMode:
			p = &parameters.TrafficMode{}
		case parameters.ParamErrorCode:
			p = &parameters.ErrorCode{}
		case parameters.ParamNetworkAppearance:
			p = &parameters.NetworkAppearance{}
		case parameters.ParamDiagnosticInformation:
			p = &parameters.DiagnosticInformation{}
		case parameters.ParamUserCause:
			p = &parameters.UserPartUnavailableCause{}
		case parameters.ParamConcernedDestinations:
			p = &parameters.ConcernedDpc{}
		case parameters.ParamCongestionIndications:
			p = &parameters.CongestionIndication{}
		case parameters.ParamProtocolData:
			p = &parameters.ProtocolData{}
		case parameters.ParamCorrelationId:
			p = &parameters.CorrelationId{}
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

package parameters

const (
	ParamInfoString                = 0x0004
	ParamRoutingContext            = 0x0006
	ParamDiagnosticInformation     = 0x0007
	ParamHeartBeadData             = 0x0009
	ParamTrafficMode               = 0x000b
	ParamErrorCode                 = 0x000c
	ParamStatus                    = 0x000d
	ParamAspIdentifier             = 0x0011
	ParamAffectedPointCode         = 0x0012
	ParamCorrelationId             = 0x0013
	ParamNetworkAppearance         = 0x0200
	ParamUserCause                 = 0x0204
	ParamCongestionIndications     = 0x0205
	ParamConcernedDestinations     = 0x0206
	ParamRoutingKey                = 0x0207
	ParamRegistrationResult        = 0x0208
	ParamDeregistrationResult      = 0x0209
	ParamLocalRoutingKeyIdentifier = 0x020a
	ParamDestinationPointCode      = 0x020b
	ParamServiceIndicators         = 0x020c
	ParamOriginationPointCodeList  = 0x020e
	ParamProtocolData              = 0x0210
	ParamRegistrationStatus        = 0x0212
	ParamDeregistrationStatus      = 0x0213
	ParamReserved                  = 0xffff
)

const (
	TrafficModeOverride = iota + 1
	TrafficModeLoadShare
	TrafficModeBroadcast
)

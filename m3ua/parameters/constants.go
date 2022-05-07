package parameters

type ParameterTag uint16

const (
	ParamInfoString                ParameterTag = 0x0004
	ParamRoutingContext            ParameterTag = 0x0006
	ParamDiagnosticInformation     ParameterTag = 0x0007
	ParamHeartBeadData             ParameterTag = 0x0009
	ParamTrafficMode               ParameterTag = 0x000b
	ParamErrorCode                 ParameterTag = 0x000c
	ParamStatus                    ParameterTag = 0x000d
	ParamAspIdentifier             ParameterTag = 0x0011
	ParamAffectedPointCode         ParameterTag = 0x0012
	ParamCorrelationId             ParameterTag = 0x0013
	ParamNetworkAppearance         ParameterTag = 0x0200
	ParamUserCause                 ParameterTag = 0x0204
	ParamCongestionIndications     ParameterTag = 0x0205
	ParamConcernedDestinations     ParameterTag = 0x0206
	ParamRoutingKey                ParameterTag = 0x0207
	ParamRegistrationResult        ParameterTag = 0x0208
	ParamDeregistrationResult      ParameterTag = 0x0209
	ParamLocalRoutingKeyIdentifier ParameterTag = 0x020a
	ParamDestinationPointCode      ParameterTag = 0x020b
	ParamServiceIndicators         ParameterTag = 0x020c
	ParamOriginationPointCodeList  ParameterTag = 0x020e
	ParamProtocolData              ParameterTag = 0x0210
	ParamRegistrationStatus        ParameterTag = 0x0212
	ParamDeregistrationStatus      ParameterTag = 0x0213
	ParamReserved                  ParameterTag = 0xffff
)

type Cause uint16

const (
	UnknownCause Cause = iota
	UnequippedRemoteUser
	InaccessibleRemoteUser
)

type Mode uint32

const (
	Override Mode = iota + 1
	LoadShare
	Broadcast
)

type ServiceIndicator uint16

const (
	Sccp                   ServiceIndicator = 3
	Tup                                     = 4
	Isup                                    = 5
	BroadbandIsup                           = 9
	StelliteIsup                            = 10
	AllType2Signalling                      = 12
	Bicc                                    = 13
	GatewayControlProtocol                  = 14
	Unknown                                 = 1
)

type Level uint32

const (
	NoCongestion Level = iota
	Level1
	Level2
	Level3
)

type NetworkIndicator uint16

const (
	International NetworkIndicator = iota
	InternationalReserve
	National
	NationalReserve
)

type M3UAError uint32

const (
	InvalidVersion            M3UAError = 0x01
	UnsupportedMessageClass   M3UAError = 0x03
	UnsupportedMessageType    M3UAError = 0x04
	UnsupportedTrafficMode    M3UAError = 0x05
	UnexpectedMessage         M3UAError = 0x06
	ProtocolError             M3UAError = 0x07
	InvalidStreamIdentifier   M3UAError = 0x09
	RefusedManagementBlocking M3UAError = 0x0d
	AspIdentifierRequired     M3UAError = 0x0e
	InvalidAspIdentifier      M3UAError = 0x0f
	InvalidParameterValue     M3UAError = 0x11
	ParameterFieldError       M3UAError = 0x12
	UnexpectedParameter       M3UAError = 0x13
	DestinationStatusUnknown  M3UAError = 0x14
	InvalidNetworkAppearance  M3UAError = 0x15
	MissingParameter          M3UAError = 0x16
	InvalidRoutingContext     M3UAError = 0x19
	NoConfiguredAsForAsp      M3UAError = 0x1a
)

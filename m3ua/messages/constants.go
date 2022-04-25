package messages

const (
	MessageClassMgmt = iota
	MessageClassTransferMessage
	MessageClassSsnm
	MessageClassAspsm
	MessageClassAsptm
	MessageClassRkm = 9
)

//Management Messages
const (
	MessageTypeErr = iota
	MessageTypeNtfy
)

//Transfer Messages
const (
	MessageTypeReserved = iota
	MessageTypeData
)

//SSNM Messages
const (
	MessageTypeDuna = iota + 1
	MessageTypeDava
	MessageTypeDaud
	MessageTypeScon
	MessageTypeDupu
	MessageTypeDrst
)

//ASPSM Messages
const (
	MessageTypeAspUp = iota + 1
	MessageTypeAspDn
	MessageTypeBeat
	MessageTypeAspUpAck
	MessageTypeAspDnAck
	MessageTypeBeatAck
)

//ASPTM Messages
const (
	MessageTypeAspAc = iota + 1
	MessageTypeAspIa
	MessageTypeAspAcAck
	MessageTypeAspIaAck
)

//RKM Messages
const (
	MessageTypeRegReq = iota + 1
	MessageTypeRegRsp
	MessageTypeDeregReq
	MessageTypeDeregRsp
)

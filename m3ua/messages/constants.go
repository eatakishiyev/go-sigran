package messages

type MessageClass int

const (
	Mgmt MessageClass = iota
	TransferMessage
	Ssnm
	Aspsm
	Asptm
	Rkm = 9
)

//Management Messages
type MessageType int

const (
	Err MessageType = iota
	Ntfy
)

//Transfer Messages
const (
	Reserved MessageType = iota
	Data
)

//SSNM Messages
const (
	Duna MessageType = iota + 1
	Dava
	Daud
	Scon
	Dupu
	Drst
)

//ASPSM Messages
const (
	AspUp MessageType = iota + 1
	AspDn
	Beat
	AspUpAck
	AspDnAck
	BeatAck
)

//ASPTM Messages
const (
	AspAc MessageType = iota + 1
	AspIa
	AspAcAck
	AspIaAck
)

//RKM Messages
const (
	RegReq MessageType = iota + 1
	RegRsp
	DeregReq
	DeregRsp
)

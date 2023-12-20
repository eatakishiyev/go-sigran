package tcap

import (
	"go-sigtran/tcap/parameters"
)

type TcMessage interface {
	GetMessageType() MessageType
}

type MessageType uint

const (
	Unidirectional MessageType = 1
	Begin          MessageType = 2
	End            MessageType = 4
	Continue       MessageType = 5
	Abort          MessageType = 7
)

// BeginMessage

type BeginMessage struct {
	OriginatingTransactionId int64                    `asn1:"tag:8,application"`
	DialogPortion            parameters.DialogPortion `asn1:"tag:11,application,optional"`
	Components               ComponentPortion         `asn1:"tag:12,application,optional"`
}

func (beginMessage *BeginMessage) GetMessageType() MessageType {
	return Begin
}

// ContinueMessage

type ContinueMessage struct {
	OriginatingTransactionId int64                    `asn1:"tag:8,application"`
	DestinationTransactionId int64                    `asn1:"tag:9,application"`
	DialogPortion            parameters.DialogPortion `asn1:"tag:11,application,optional"`
	Components               ComponentPortion         `asn1:"tag:12,application,optional,universal,set"`
}

func (continueMessage *ContinueMessage) GetMessageType() MessageType {
	return Continue
}

// EndMessage

type EndMessage struct {
	DestinationTransactionId int64                    `asn1:"tag:9,application"`
	DialogPortion            parameters.DialogPortion `asn1:"tag:11,application,optional"`
	Components               ComponentPortion         `asn1:"tag:12,application,optional"`
}

func (endMessage *EndMessage) GetMessageType() MessageType {
	return End
}

// AbortMessage

type AbortMessage struct {
	DestinationTransactionId int64                  `asn1:"tag:9,application"`
	PAbortCause              parameters.PAbortCause `asn1:"tag:10,application,optional"`
	UAbortCauseBytes         []byte                 `asn1:"tag:11,application,optional"`
}

func (abortMessage *AbortMessage) GetMessageType() MessageType {
	return Abort
}

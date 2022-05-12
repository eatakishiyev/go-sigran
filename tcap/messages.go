package tcap

import (
	"encoding/asn1"
)

type MessageType uint

const (
	Unidirectional MessageType = 1
	Begin          MessageType = 2
	End            MessageType = 4
	Continue       MessageType = 5
	Abort          MessageType = 7
)

type TCAPMessage interface {
	DecodeMessage(data []byte) ([]byte, error)
	GetMessageType() MessageType
}

type BeginMessage struct {
	OriginatingDialogId int64         `asn1:"tag:8,application"`
	DialogPortion       asn1.RawValue `asn1:"tag:11,application,optional"`
	Components          asn1.RawValue `asn1:"tag:12,application,optional"`
}

type ContinueMessage struct {
	OriginatingDialogId int           `asn1:"tag:8,application"`
	DestinationDialogId int           `asn1:"tag:9,application"`
	DialogPortion       asn1.RawValue `asn1:"tag:11,application,optional,explicit"`
	Components          asn1.RawValue `asn1:"tag:12,application,optional,explicit"`
}

type EndMessage struct {
	DestinationDialogId int           `asn1:"tag:9,application"`
	DialogPortion       asn1.RawValue `asn1:"tag:11,application,optional,explicit"`
	Components          asn1.RawValue `asn1:"tag:12,application,optional,explicit"`
}

type AbortMessage struct {
	DestinationTransactionId int           `asn1:"tag:9,application"`
	PAbortCause              int           `asn1:"tag:10,application,optional,explicit"`
	UAbortCause              asn1.RawValue `asn1:"tag:11,application,optional,explicit"`
}

func (b *BeginMessage) DecodeMessage(data []byte) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, b, "tag:2,application")
}

func (b *BeginMessage) GetMessageType() MessageType {
	return Begin
}

func (c *ContinueMessage) DecodeMessage(data []byte) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, c, "tag:5,application")
}

func (c *ContinueMessage) GetMessageType() MessageType {
	return Continue
}

func (e *EndMessage) DecodeMessage(data []byte) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, e, "tag:4,application")
}

func (e *EndMessage) GetMessageType() MessageType {
	return End
}

func (a *AbortMessage) DecodeMessage(data []byte) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, a, "tag:7,application")
}

func (a *AbortMessage) GetMessageType() MessageType {
	return Abort
}

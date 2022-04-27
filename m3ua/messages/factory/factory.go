package factory

import (
	"bytes"
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/messages/aspsm"
	"go-sigtran/m3ua/messages/asptm"
	"go-sigtran/m3ua/messages/mgmt"
	"go-sigtran/m3ua/messages/transfer"
)

func CreateMessage(r *bytes.Reader) *messages.Message {
	header := messages.DecodeHeader(r)
	data := make([]byte, header.MessageLength-8)
	r.Read(data)

	var m messages.Message

	switch header.MessageClass {
	case messages.MessageClassAspsm:
		switch header.MessageType {
		case messages.MessageTypeAspDn:
			m = &aspsm.AspDown{}
		case messages.MessageTypeAspDnAck:
			m = &aspsm.AspDownAck{}
		case messages.MessageTypeAspUp:
			m = &aspsm.AspUp{}
		case messages.MessageTypeAspUpAck:
			m = &aspsm.AspUpAck{}
		case messages.MessageTypeBeat:
			m = &aspsm.Heartbeat{}
		case messages.MessageTypeBeatAck:
			m = &aspsm.HeartbeatAck{}
		}
	case messages.MessageClassAsptm:
		switch header.MessageType {
		case messages.MessageTypeAspAc:
			m = &asptm.AspActive{}
		case messages.MessageTypeAspAcAck:
			m = &asptm.AspActiveAck{}
		case messages.MessageTypeAspIa:
			m = &asptm.AspInactive{}
		case messages.MessageTypeAspIaAck:
			m = &asptm.AspInactiveAck{}
		}
	case messages.MessageClassMgmt:
		switch header.MessageType {
		case messages.MessageTypeErr:
			m = &mgmt.Error{}
		}

	case messages.MessageClassTransferMessage:
		switch header.MessageType {
		case messages.MessageTypeData:
			m = &transfer.PayloadData{}
		}
	}
	if m != nil {
		m.SetHeader(header)
		m.DecodeMessage(data)
	}
	return &m
}

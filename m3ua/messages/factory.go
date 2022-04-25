package messages

import (
	"bytes"
	"go-sigtran/m3ua/messages/aspsm"
	"go-sigtran/m3ua/messages/asptm"
	"go-sigtran/m3ua/messages/mgmt"
)

func CreateMessage(r *bytes.Reader) *Message {
	header := DecodeHeader(r)
	data := make([]byte, header.MessageLength-8)
	r.Read(data)

	var m Message

	switch header.MessageClass {
	case MessageClassAspsm:
		switch header.MessageType {
		case MessageTypeAspDn:
			m = &aspsm.AspDown{}
		case MessageTypeAspDnAck:
			m = &aspsm.AspDownAck{}
		case MessageTypeAspUp:
			m = &aspsm.AspUp{}
		case MessageTypeAspUpAck:
			m = &aspsm.AspUpAck{}
		case MessageTypeBeat:
			m = &aspsm.Heartbeat{}
		case MessageTypeBeatAck:
			m = &aspsm.HeartbeatAck{}
		}
	case MessageClassAsptm:
		switch header.MessageType {
		case MessageTypeAspAc:
			m = &asptm.AspActive{}
		case MessageTypeAspAcAck:
			m = &asptm.AspActiveAck{}
		case MessageTypeAspIa:
			m = &asptm.AspInactive{}
		case MessageTypeAspIaAck:
			m = &asptm.AspInactiveAck{}
		}
	case MessageClassMgmt:
		switch header.MessageType {
		case MessageTypeErr:
			m = &mgmt.Error{}
		}
	}
	if m != nil {
		m.SetHeader(header)
		m.DecodeMessage(data)
	}
	return &m
}

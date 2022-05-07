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
	case messages.Aspsm:
		switch header.MessageType {
		case messages.AspDn:
			m = &aspsm.AspDown{}
		case messages.AspDnAck:
			m = &aspsm.AspDownAck{}
		case messages.AspUp:
			m = &aspsm.AspUp{}
		case messages.AspUpAck:
			m = &aspsm.AspUpAck{}
		case messages.Beat:
			m = &aspsm.Heartbeat{}
		case messages.BeatAck:
			m = &aspsm.HeartbeatAck{}
		}
	case messages.Asptm:
		switch header.MessageType {
		case messages.AspAc:
			m = &asptm.AspActive{}
		case messages.AspAcAck:
			m = &asptm.AspActiveAck{}
		case messages.AspIa:
			m = &asptm.AspInactive{}
		case messages.AspIaAck:
			m = &asptm.AspInactiveAck{}
		}
	case messages.Mgmt:
		switch header.MessageType {
		case messages.Err:
			m = &mgmt.Error{}
		}

	case messages.TransferMessage:
		switch header.MessageType {
		case messages.Data:
			m = &transfer.PayloadData{}
		}
	}
	if m != nil {
		m.SetHeader(header)
		m.DecodeMessage(data)
	}
	return &m
}

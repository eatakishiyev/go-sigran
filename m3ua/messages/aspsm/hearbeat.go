package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type Heartbeat struct {
	*messages.MessageHeader
	*parameters.HeartbeatData
}

func NewHeartbeat(data *parameters.HeartbeatData) *Heartbeat {
	return &Heartbeat{
		MessageHeader: messages.NewHeader(messages.Aspsm, messages.Beat),
		HeartbeatData: data,
	}
}

func (h *Heartbeat) EncodeMessage() []byte {
	var encoded []byte
	if h.HeartbeatData != nil {
		parameters.Encode(encoded, h.HeartbeatData)
	}
	return encoded
}

func (h *Heartbeat) DecodeMessage(b []byte) {
	if len(b) > 0 {
		params := factory.DecodeParameters(b)
		for idx := 0; idx < len(params); idx++ {
			p := params[idx]
			switch p.GetHeader().Tag {
			case parameters.ParamHeartBeadData:
				h.HeartbeatData = p.(*parameters.HeartbeatData)
			}
		}
	}
}

func (h *Heartbeat) GetHeader() *messages.MessageHeader {
	return h.MessageHeader
}

func (h *Heartbeat) SetHeader(header *messages.MessageHeader) {
	h.MessageHeader = header
}

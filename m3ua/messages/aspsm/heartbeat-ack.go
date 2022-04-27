package aspsm

import (
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/parameters"
	"go-sigtran/m3ua/parameters/factory"
)

type HeartbeatAck struct {
	*messages.MessageHeader
	*parameters.HeartbeatData
}

func (h *HeartbeatAck) EncodeMessage() []byte {
	var encoded []byte
	if h.HeartbeatData != nil {
		parameters.Encode(encoded, h.HeartbeatData)
	}
	return encoded
}

func (h *HeartbeatAck) DecodeMessage(b []byte) {
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

func (h *HeartbeatAck) GetHeader() *messages.MessageHeader {
	return h.MessageHeader
}

func (h *HeartbeatAck) SetHeader(header *messages.MessageHeader) {
	h.MessageHeader = header
}

func NewHeartbeatAck(heartbeatData *parameters.HeartbeatData) *HeartbeatAck {
	return &HeartbeatAck{
		MessageHeader: messages.NewHeader(messages.MessageClassAspsm, messages.MessageTypeBeatAck),
		HeartbeatData: heartbeatData,
	}
}

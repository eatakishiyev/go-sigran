package messages

import (
	"bytes"
	"encoding/binary"
)

type MessageHeader struct {
	Version       byte         // 1 byte
	Reserved      byte         // 1 byte
	MessageClass  MessageClass // 1 byte
	MessageType   MessageType  // 1 byte
	MessageLength uint32       // 4 byte
}

type Message interface {
	EncodeMessage() []byte
	DecodeMessage(b []byte)
	GetHeader() *MessageHeader
	SetHeader(header *MessageHeader)
}

func NewHeader(messageClass MessageClass, messageType MessageType) *MessageHeader {
	return &MessageHeader{
		Version:       1,
		Reserved:      0,
		MessageClass:  messageClass,
		MessageType:   messageType,
		MessageLength: 0,
	}
}

func EncodeMessage(m Message) []byte {
	var encoded []byte

	encoded = append(encoded, m.GetHeader().Version)
	encoded = append(encoded, m.GetHeader().Reserved)
	encoded = append(encoded, byte(m.GetHeader().MessageClass))
	encoded = append(encoded, byte(m.GetHeader().MessageType))
	encodedMessage := m.EncodeMessage()
	messageLength := len(encodedMessage) + 8 // version(1 byte) + reserved(1 byte) + class(1 byte) + type(1 byte) + length(4 byte)
	binary.BigEndian.PutUint32(encoded, uint32(messageLength))
	encoded = append(encoded, encodedMessage...)

	return encoded
}

func DecodeHeader(r *bytes.Reader) *MessageHeader {
	//TODO validate version of message and in case incorrect version raise error
	r.ReadByte() //Version
	r.ReadByte() //Reserved
	messageClass, _ := r.ReadByte()
	messageType, _ := r.ReadByte()
	remainingBytes := make([]byte, 4)
	r.Read(remainingBytes)
	length := binary.BigEndian.Uint32(remainingBytes)

	header := NewHeader(MessageClass(messageClass), MessageType(messageType))
	header.MessageLength = length
	return header
}

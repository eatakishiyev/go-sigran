package main

import (
	"fmt"
	"go-sigtran/m3ua/messages/aspsm"
)

func main() {
	//b := []byte{0x01, 0x00, 0x03, 0x01, 0x00, 0x00, 0x00, 0x10, 0x00, 0x11, 0x00, 0x08, 0x00, 0x00, 0x00, 0x0d}
	//aspUp := aspsm.CreateMessage(bytes.NewReader(b))
	ack := &aspsm.HeartbeatAck{}
	fmt.Printf("%p\n", &ack)
	ack.GetHeader()
	//fmt.Printf("%s", aspUp)
}

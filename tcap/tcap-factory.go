package tcap

import (
	"encoding/asn1"
	"fmt"
)

type TcapMessageFactory struct {
}

func (tf TcapMessageFactory) DecodeTCAPMessage(data []byte) TCAPMessage {
	var rawValue asn1.RawValue
	_, err := asn1.Unmarshal(data, &rawValue)

	if err != nil {
		fmt.Printf("error occurred during decode TCAP message %s\n", err)
		return nil
	}

	var tcapMessage TCAPMessage
	switch rawValue.Tag {
	case 2: //BeginMessage Message
		tcapMessage = new(BeginMessage)
		_, err := tcapMessage.DecodeMessage(data)

		if err != nil {
			fmt.Printf("error occurred during decode dialogPortion %s\n", err)
		}
	case 4: //End Message
		tcapMessage = new(EndMessage)
		tcapMessage.DecodeMessage(data)
	case 5: //Continue Message
		tcapMessage = new(ContinueMessage)
		tcapMessage.DecodeMessage(data)
	case 7: //Abort Message
		tcapMessage = new(AbortMessage)
		tcapMessage.DecodeMessage(data)
	}

	return tcapMessage
}

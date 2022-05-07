package tcap

import (
	"encoding/asn1"
	"fmt"
)

type TcapMessageFactory struct {
}

func (tf TcapMessageFactory) DecodeTCAPMessage(data []byte) {
	var rawValue asn1.RawValue
	_, err := asn1.Unmarshal(data, &rawValue)

	if err != nil {
		fmt.Printf("error occurred during decode TCAP message %s\n", err)
		return
	}

	switch rawValue.Tag {
	case 2: //BeginMessage Message
		beginMessage := new(BeginMessage)
		_, err := beginMessage.Decode(data)
		if err != nil {
			if err != nil {
				fmt.Printf("error occurred during decode dialogPortion %s\n", err)
			}
		}
	case 4: //End Message
		endMessage := new(EndMessage)
		endMessage.Decode(data)
		fmt.Printf("%v", endMessage)
	case 5: //Continue Message
		continueMessage := new(ContinueMessage)
		continueMessage.Decode(data)
		fmt.Printf("%v", continueMessage)
	case 7: //Abort Message
		abortMessage := new(AbortMessage)
		abortMessage.Decode(data)
		fmt.Printf("%v", abortMessage)
	}
}

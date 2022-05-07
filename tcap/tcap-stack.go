package tcap

import "go-sigtran/sccp"

type TcapStack struct {
}

func (ts *TcapStack) OnMessage(calledParty sccp.SccpAddress, callingParty sccp.SccpAddress, userData []byte,
	sequenceControl bool, sequenceNumber uint, messageHandling sccp.MessageHandling) {

}

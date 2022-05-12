package tcap

import (
	"go-sigtran/sccp"
	"go-sigtran/sccp/parameters"
	"math/rand"
	"sync"
	"time"
)

type TcapStack struct {
	*TcapMessageFactory
	Dialogs           map[int]*TcapDialog
	TransactionIdPool []int
	Name              string
	KeepAliveTime     time.Duration
	mutex             sync.Mutex
}

func NewTcapStack(name string, minTransactionId int, maxTransactionId int, keepAliveTimer time.Duration) *TcapStack {
	var transactionIdPool []int
	for trId := minTransactionId; trId < maxTransactionId; trId++ {
		transactionIdPool = append(transactionIdPool, trId)
	}
	rand.Shuffle(len(transactionIdPool), func(i, j int) {
		transactionIdPool[i], transactionIdPool[j] = transactionIdPool[j], transactionIdPool[i]
	})

	tcapStack := TcapStack{
		Name:              name,
		TransactionIdPool: transactionIdPool,
		KeepAliveTime:     keepAliveTimer,
	}

	return &tcapStack
}

func (ts *TcapStack) borrowTransactionId() int {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	if len(ts.TransactionIdPool) <= 0 {
		return -1
	}
	trId := ts.TransactionIdPool[0]
	ts.TransactionIdPool = ts.TransactionIdPool[1:]
	return trId
}

func (ts *TcapStack) releaseTransactionId(transactionId int) {
	ts.mutex.Lock()
	defer ts.mutex.Unlock()
	ts.TransactionIdPool = append(ts.TransactionIdPool, transactionId)
}

func (ts *TcapStack) OnMessage(calledParty *parameters.SccpAddress, callingParty *parameters.SccpAddress, userData []byte,
	sequenceControl bool, sequenceNumber uint, messageHandling sccp.MessageHandling) {
	message := TcapMessageFactory{}.DecodeTCAPMessage(userData)
	messageType := message.GetMessageType()
	switch messageType {
	case Begin:
		trId := ts.borrowTransactionId()
		if trId < 0 {
			//TODO implement free
		}
		tcapDialog := new(TcapDialog)
		tcapDialog.beginMessage(message.(*BeginMessage), callingParty, calledParty)
		ts.Dialogs[trId] = tcapDialog
	case Continue:
		continueMessage := message.(*ContinueMessage)
		tcapDialog, ok := ts.Dialogs[continueMessage.DestinationDialogId]
		if ok {
			tcapDialog.continueMessage(continueMessage, callingParty, calledParty)
		} else {
			//TODO implement abort
		}
	case End:
	case Abort:
		abortMessage := message.(*AbortMessage)
		tcapDialog := new(TcapDialog)
		tcapDialog.abortMessage(abortMessage, callingParty, calledParty)

	}
}

func (ts *TcapStack) OnNotice(calledParty *parameters.SccpAddress, callingParty *parameters.SccpAddress, userData []byte,
	errorReason sccp.ErrorReason, importance uint) {

}

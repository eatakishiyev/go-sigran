package tcap

import (
	"go-sigtran/sccp"
	sccpparameters "go-sigtran/sccp/parameters"
	"math/rand"
	"sync"
	"time"
)

type TcapStack struct {
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
		Dialogs:           make(map[int]*TcapDialog),
	}

	return &tcapStack
}

func (tcapStack *TcapStack) BorrowTransactionId() int {
	tcapStack.mutex.Lock()
	defer tcapStack.mutex.Unlock()
	if len(tcapStack.TransactionIdPool) <= 0 {
		return -1
	}
	trId := tcapStack.TransactionIdPool[0]
	tcapStack.TransactionIdPool = tcapStack.TransactionIdPool[1:]
	return trId
}

func (tcapStack *TcapStack) releaseTransactionId(transactionId int) {
	tcapStack.mutex.Lock()
	defer tcapStack.mutex.Unlock()
	tcapStack.TransactionIdPool = append(tcapStack.TransactionIdPool, transactionId)
}

func (tcapStack *TcapStack) OnMessage(calledParty *sccpparameters.SccpAddress, callingParty *sccpparameters.SccpAddress, userData []byte,
	sequenceControl bool, sequenceNumber uint, messageHandling sccp.MessageHandling) {
	var tcapMessage interface{}
	Decode(userData, &tcapMessage)

	//qoS := &tcap_parameters.QoS{
	//	MessageHandling: messageHandling,
	//	SequenceControl: sequenceControl,
	//	SequenceNumber:  sequenceNumber}

	switch tcapMessage.(type) {
	case BeginMessage:
		println("BeginMessage")
	case EndMessage:
		println("EndMessage")
	case ContinueMessage:
		println("ContinueMessage")
	case AbortMessage:
		println("AbortMessage")
	}
	//switch MessageType(raw.Tag) {
	//case Begin:
	//	beginMessage := BeginMessage{}
	//	_, err := asn1.UnmarshalWithParams(raw.FullBytes, &beginMessage, "tag:2,application,class:1")
	//	fmt.Printf("%d", beginMessage.GetOriginatingTransactionId())
	//	if err != nil {
	//		fmt.Println(err)
	//		return
	//	}
	//	trId := tcapStack.BorrowTransactionId()
	//	if trId < 0 {
	//		//TODO implement send abort message, out of resource
	//	}
	//	tcapDialog := new(TcapDialog)
	//	tcapDialog.TcapStack = tcapStack
	//	tcapDialog.QoS = qoS
	//
	//	tcapStack.Dialogs[trId] = tcapDialog
	//	//tcapDialog.beginMessage(&beginMessage, callingParty, calledParty)
	//	//case Continue:
	//	//	continueMessage := tcapMessage.(ContinueMessage)
	//	//	tcapDialog, ok := ts.Dialogs[continueMessage.DestinationTransactionId.Id]
	//	//	if ok {
	//	//		tcapDialog.TcapStack = ts
	//	//		tcapDialog.QoS = qoS
	//	//		tcapDialog.continueMessage(&continueMessage, callingParty, calledParty)
	//	//	} else {
	//	//		if continueMessage.OriginatingTransactionId != nil {
	//	//			ts.sendTrPortionAbortMessage(continueMessage.OriginatingTransactionId, calledParty, callingParty,
	//	//				tcapparameters.UnrecognizedTransactionId, qoS)
	//	//		}
	//	//	}
	//	//case End:
	//	//case Abort:
	//	//	abortMessage := tcapMessage.(AbortMessage)
	//	//	tcapDialog, ok := ts.Dialogs[abortMessage.DestinationTransactionId.Id]
	//	//	if ok {
	//	//		tcapDialog.abortMessage(&abortMessage, callingParty, calledParty)
	//	//	}
	//}
}

func (tcapStack *TcapStack) OnNotice(calledParty *sccpparameters.SccpAddress, callingParty *sccpparameters.SccpAddress, userData []byte,
	errorReason sccp.ErrorReason, importance uint) {

}

//func (ts *TcapStack) sendTrPortionAbortMessage(transactionId *tcapparameters.TransactionId, calledParty *sccpparameters.SccpAddress,
//	callingParty *sccpparameters.SccpAddress, cause tcapparameters.PAbortCause, qoS *tcapparameters.QoS) {
//	abortMessage := AbortMessage{DestinationTransactionId: transactionId,
//		PAbortCause: tcapparameters.UnrecognizedTransactionId}
//	ts.Send(callingParty, calledParty, qoS, abortMessage.Encode())
//}

package main

import (
	"encoding/hex"
	"fmt"
	"go-sigtran/m3ua"
	"go-sigtran/sccp"
	"go-sigtran/tcap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type Queue struct {
	q []int
}

func (q *Queue) Dequeue() int {
	var e = q.q[0]
	q.q = q.q[1:]
	return e
}

func (q *Queue) Enqueue(e int) {
	q.q = append(q.q, e)
}

//func testEncodeDialogPortion() {
//
//	var h, _ = hex.DecodeString("9b139192")
//	var dialogId = binary.BigEndian.Uint32(h)
//
//	abortPdu := tcap_parameters.DialogAbortPDU{
//		AbortSource: tcap_parameters.DialogServiceProvider,
//	}
//
//	dialogPortion := tcap_parameters.DialogPortion{
//		DialogPDU: abortPdu,
//	}
//	var b = dialogPortion.Encode()
//	fmt.Printf("encoded dialog-portion %x\n", b)
//
//	var abortMessage = tcap.AbortMessage{
//		DestinationTransactionId: int(dialogId),
//	}
//
//	var abortData, err = asn1.EncodeWithOptions(abortMessage, "tag:7,application")
//	if err != nil {
//		fmt.Printf("error occurred during encode abort message %s", err)
//	} else {
//		fmt.Printf("abort message data %x", abortData)
//	}
//
//}

func main() {
	//testEncodeDialogPortion()

	//End
	var messageBytes = "6481844904a11003266b2a2828060700118605010101a01d611b80020780a109060704000001000503a203020100a305a1030201006c80a24c0201013047020116a380890804007131929364f9040791995480316763a1030401f8820791995490040871a019a017301506092a863a0089613a0100a508300681010285010a8f0204c000000000"
	//var messageBytes = "65819a480482f403034904431001566c818ba12c02010202012e3024a01da01ba119810101a214a012800101810101820103830101840101850103a103800101a144020103020117303ca03a30068001048101003006800105810100300680010781010030068001fe810100300b800109810100a203800102300b800106810100be03810141a115020104020114300da00b0409029000995430477725"
	//var messageBytes = "64144904bf1c06d16c0ca10a020104020116040282e6"
	//var messageBytes = "654a480443100156490482f403036c3ca11b02010202012404138001018106327082010201a20381010183012ba11d0201030201183015800109a206a70480028090a303810101a403800100"
	//Begin
	//var messageBytes = "6581b24802047b4904070004006b2a2828060700118605010101a01d611b80020780a109060704000001003201a203020100a305a1030201006c7aa165020101020117305da05b300b800104810100a203800102300b800105810100a203800102300b800106810100a203800102300b800107810101a203800102300b800109810100a203800101300b800109810100a203800102300b80010a810101a203800101a1110201020201143009a00704050210792210"
	data, err := hex.DecodeString(messageBytes)
	if err != nil {
		fmt.Printf("error occurred while decode hex-string %s", err)
	}

	tcapStack := tcap.NewTcapStack("TestTcapStack", 1, 10, time.Duration(10)*time.Second)

	tcapStack.OnMessage(nil, nil, data, true, 1, sccp.ReturnMessageOnError)

	asp := &m3ua.Asp{
		BindPort: 2905,
		BindIp:   "127.0.0.1",
		IsServer: true,
	}

	asp.Start()

	quitChannel := make(chan os.Signal, 1)
	signal.Notify(quitChannel, syscall.SIGINT, syscall.SIGTERM)
	<-quitChannel

	stack := m3ua.GetM3UAStackInstance()
	fmt.Printf("Stack = %s\n", stack)
	stack = m3ua.GetM3UAStackInstance()
	fmt.Printf("Stack = %s\n", stack)
}

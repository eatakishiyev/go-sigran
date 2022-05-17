package main

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"github.com/PromonLogicalis/asn1"
	"go-sigtran/m3ua"
	"go-sigtran/sccp"
	"go-sigtran/tcap"
	tcap_parameters "go-sigtran/tcap/parameters"
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

func testEncodeDialogPortion() {

	var h, _ = hex.DecodeString("9b139192")
	var dialogId = binary.BigEndian.Uint64(h)

	abortPdu := tcap_parameters.DialogAbortPDU{
		AbortSource: tcap_parameters.DialogServiceProvider,
	}

	dialogPortion := tcap_parameters.DialogPortion{
		DialogPDU: abortPdu,
	}
	var b = dialogPortion.Encode()
	fmt.Printf("encoded dialog-portion %x\n", b)

	var abortMessage = tcap.AbortMessage{
		DestinationTransactionId: int(dialogId),
	}

	var abortData, err = asn1.EncodeWithOptions(abortMessage, "tag:7,application")
	if err != nil {
		fmt.Printf("error occurred during encode abort message %s", err)
	} else {
		fmt.Printf("abort message data %x", abortData)
	}

}

func main() {
	testEncodeDialogPortion()

	var tcapMessageData = "4b10280e060700118605010101a003800101"

	data, err := hex.DecodeString(tcapMessageData)
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

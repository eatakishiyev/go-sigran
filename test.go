package main

import (
	"encoding/hex"
	"fmt"
	"go-sigtran/m3ua"
	"go-sigtran/tcap"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	var tcapMessageData = "6281a6480403fb41066b1a2818060700118605010101a00d600ba1090607040000010032016c8181a17f0201010201003077800200c98308041399549513798885010a8a080413995405001012bb0580038090a39c01029f320804007210661068f2bf3417020100810791995405001012a309800704f020060cd554bf35038301119f360513984a3a229f3707919954050010129f3807919954255179829f39080222308241034223"

	data, err := hex.DecodeString(tcapMessageData)
	if err != nil {
		fmt.Printf("error occurred while decode hex-string %s", err)
	}

	tcap.TcapMessageFactory{}.DecodeTCAPMessage(data)

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

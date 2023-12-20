package m3ua

import (
	"fmt"
	_ "github.com/ishidawataru/sctp"
	sctp "github.com/ishidawataru/sctp"
	"github.com/looplab/fsm"
	"go-sigtran/m3ua/messages"
	"go-sigtran/m3ua/messages/aspsm"
	"go-sigtran/m3ua/messages/asptm"
	"go-sigtran/m3ua/messages/ssnm"
	"go-sigtran/m3ua/parameters"
	"math/rand"
	"net"
	"strings"
	"time"
)

type Asp struct {
	aspStateMachines map[uint32]*fsm.FSM
	association      *sctp.SCTPConn
	Name             string
	sls2Stream       []uint16
	*parameters.TrafficMode
	*parameters.InfoString
	*parameters.AspIdentifier
	IsServer   bool
	BindPort   int
	BindIp     string
	RemotePort int
	RemoteIp   string
}

func (a *Asp) Start() {
	if a.IsServer {
		var ips []net.IPAddr
		var bindAddresses = strings.Split(a.BindIp, ";")
		for _, bindAddress := range bindAddresses {
			addr, err := net.ResolveIPAddr("ip", bindAddress)
			if err == nil {
				ips = append(ips, *addr)
			} else {
				fmt.Printf("error occurred while resolve ip address %s", addr)
			}
		}

		listenSCTP, err := sctp.ListenSCTP("sctp", &sctp.SCTPAddr{
			IPAddrs: ips,
			Port:    a.BindPort,
		})

		if err != nil {
			fmt.Printf("error occurred while starting sctp server : %s", err)
			return
		}
		go func() {
			for {
				a.association, err = listenSCTP.AcceptSCTP()
				var remoteAddr = a.association.RemoteAddr()

				if remoteAddr.String() != strings.Join([]string{a.RemoteIp, string(a.RemotePort)}, ":") {
					a.association.Close()
				}

				if err != nil {
					fmt.Printf("error occurred during accept incoming connection %s", a.Name)
				}

				fmt.Printf("incoming connection accepted")

				time.Sleep(time.Duration(5) * time.Millisecond)
			}
		}()

	} else {

	}
}

func (a *Asp) mapSlsToStream(outboundStreams int) {
	a.sls2Stream = make([]uint16, M3UAConfig.MaxSls)
	var streamId = 1
	a.sls2Stream[0] = 0
	for i := 1; i < len(a.sls2Stream); i++ {
		if streamId >= outboundStreams-1 {
			streamId = 1
		}
		streamId += 1
		a.sls2Stream[i] = uint16(streamId)
	}
}

func (a *Asp) Send(message messages.Message, sls uint) {

	encoded := messages.EncodeMessage(message)
	var streamNumber uint16 = 0

	switch message.GetHeader().MessageClass {
	case messages.TransferMessage:
		streamNumber = a.sls2Stream[sls]
		if streamNumber == 0 {
			streamNumber = 1
		}
	case messages.Mgmt:
	case messages.Rkm:
	case messages.Aspsm:
		switch message.GetHeader().MessageType {
		case messages.Beat:
		case messages.BeatAck:
		case messages.Ntfy:
			streamNumber = a.sls2Stream[rand.Int31n(int32(M3UAConfig.MaxSls))]
		default:
			streamNumber = 0
		}
	}

	a.association.SCTPWrite(encoded, &sctp.SndRcvInfo{
		Stream: streamNumber,
		SSN:    3,
	})
}

func (a *Asp) SendAspActive(rc *uint32) {
	time.AfterFunc(time.Duration(M3UAConfig.DelayBetweenInitiationMessages)*time.Second, func() {
		aspActive := asptm.NewAspActive(a.TrafficMode, parameters.NewRoutingContext([]uint32{*rc}), a.InfoString)
		a.Send(aspActive, 0)
		a.transiteAspStateMachine(rc, "sendActive")
	})
}

func (a *Asp) SendAspActiveAck(rc *uint32) {
	aspActiveAck := asptm.NewAspInactiveAck(parameters.NewRoutingContext([]uint32{*rc}), a.InfoString)
	a.Send(aspActiveAck, 0)
}

func (a *Asp) SendAspInactiveAck(rc *uint32) {
	aspInactiveAck := asptm.NewAspInactiveAck(parameters.NewRoutingContext([]uint32{*rc}), a.InfoString)
	a.Send(aspInactiveAck, 0)
	a.transiteAspStateMachine(rc, "aspInactive")
}

func (a *Asp) SendDava(affectedPointCode *parameters.AffectedPointCode) {
	dava := ssnm.NewDestinationAvailable(nil, affectedPointCode, a.InfoString)
	a.Send(dava, 0)
}

func (a *Asp) SendAspInactive(rc *uint32) {
	time.AfterFunc(time.Duration(M3UAConfig.DelayBetweenInitiationMessages)*time.Second, func() {
		aspInactive := asptm.NewAspInactive(parameters.NewRoutingContext([]uint32{*rc}), a.InfoString)
		a.Send(aspInactive, 0)
		a.transiteAspStateMachine(rc, "aspInactiveSent")
	})
}

func (a *Asp) SendAspUpAck() {
	aspUpAck := aspsm.NewAspUpAck(a.InfoString, a.AspIdentifier)
	a.Send(aspUpAck, 0)
	a.transiteAspStateMachine(nil, "aspUpAck")
}

func (a *Asp) SendAspUp() {
	time.AfterFunc(time.Duration(M3UAConfig.DelayBetweenInitiationMessages)*time.Second, func() {
		aspUp := aspsm.NewAspUp(a.AspIdentifier, a.InfoString)
		a.Send(aspUp, 0)
		a.transiteAspStateMachine(nil, "sendUp")
	})
}

func (a *Asp) SendHeartBeat(data []byte) {
	heartBeat := aspsm.NewHeartbeat(parameters.NewHeartbeatData(data))
	a.Send(heartBeat, 0)
}

func (a *Asp) SendHeartBeatAck(data []byte) {
	heartBeatAck := aspsm.NewHeartbeatAck(parameters.NewHeartbeatData(data))
	a.Send(heartBeatAck, 0)
}

func (a *Asp) transiteAspStateMachine(rc *uint32, event string) {
	if rc == nil {
		for _, fsm := range a.aspStateMachines {
			fsm.Event(event)
		}
	} else {
		fsm, ok := a.aspStateMachines[*rc]
		if ok {
			fsm.Event(event)
		}
	}
}

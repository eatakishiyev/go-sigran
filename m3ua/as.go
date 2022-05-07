package m3ua

import (
	"fmt"
	"github.com/looplab/fsm"
	"go-sigtran/m3ua/messages/transfer"
	"go-sigtran/m3ua/parameters"
)

type As struct {
	fsm *fsm.FSM
	*parameters.RoutingContext
	*parameters.NetworkAppearance
	Name string
	Asps []*Asp
}

func (as *As) Init() {

	var events = fsm.Events{
		{Name: "asInactive", Src: []string{"asDown", "asInactive", "asPending"}, Dst: "asInactive"},
		{Name: "asDown", Src: []string{"asDown", "asInactive", "asActive", "asPending"}, Dst: "asDown"},
		{Name: "asActive", Src: []string{"asInactive", "asActive", "asPending"}, Dst: "asActive"},
		{Name: "asPending", Src: []string{"asInactive"}, Dst: "asInactive"},
		{Name: "asPending", Src: []string{"asActive"}, Dst: "asPending"},
		{Name: "payload", Src: []string{"asActive"}, Dst: "asActive"},
	}
	as.fsm = fsm.NewFSM("asDown", events, fsm.Callbacks{})

	var str, _ = fsm.VisualizeWithType(as.fsm, fsm.GRAPHVIZ)
	fmt.Printf("%s\n", str)
}

func (as *As) AddAsp(asp *Asp) {
	as.Asps = append(as.Asps, asp)
}

func (as *As) Send(payload transfer.PayloadData) {
	err := as.fsm.Event("payload")
	if err != nil {
		fmt.Printf("Cannot send payload: %s", err)
		return
	}

	if as.NetworkAppearance != nil {
		payload.NetworkAppearance = as.NetworkAppearance
	}

	if as.RoutingContext != nil {
		payload.RoutingContext = as.RoutingContext
	}

	var aspIndex = int(payload.ProtocolData.Sls & 0x7F)
	aspIndex = aspIndex % len(as.Asps)

	//var asp = as.Asps[aspIndex]

	//TODO implement payload sending
}

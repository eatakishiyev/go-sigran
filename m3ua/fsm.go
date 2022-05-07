package m3ua

import (
	"fmt"
	"github.com/looplab/fsm"
)

func NewAspFsm() *fsm.FSM {
	var aspFsm *fsm.FSM

	var events = fsm.Events{
		{Name: "sendUp", Src: []string{"aspDown"}, Dst: "upSent"},
		{Name: "commLost", Src: []string{"aspDown", "upSent", "aspInactive", "activeSent", "aspActive", "aspDownSent",
			"aspInactiveSent", "aspActiveSent"}, Dst: "aspDown"},
		{Name: "aspUpAck", Src: []string{"upSent"}, Dst: "aspInactive"},
		{Name: "sendActive", Src: []string{"aspInactive"}, Dst: "aspActiveSent"},
		{Name: "sendDown", Src: []string{"aspInactive", "aspActive"}, Dst: "aspDownSent"},
		{Name: "aspActiveAck", Src: []string{"aspActiveSent"}, Dst: "aspActive"},
		{Name: "alternateAspActive", Src: []string{"aspActive"}, Dst: "aspInactive"},
		{Name: "sendInactive", Src: []string{"aspActive"}, Dst: "aspInactive"},
		{Name: "aspInactive", Src: []string{"aspActive"}, Dst: "aspInactive"},
		{Name: "aspDownAck", Src: []string{"aspDownSent"}, Dst: "aspDown"},
		{Name: "aspInactiveAck", Src: []string{"aspInactiveSent"}, Dst: "aspInactive"},
	}

	var callbacks = fsm.Callbacks{
		"enter_upSent": func(event *fsm.Event) {
			fmt.Printf("Entering %s\n", event.Dst)
		},
	}

	aspFsm = fsm.NewFSM("aspDown", events, callbacks)
	fmt.Printf("Transition Result is %s\n", aspFsm.Event("sendUp"))
	var str, _ = fsm.VisualizeWithType(aspFsm, fsm.GRAPHVIZ)
	fmt.Printf("%s\n", str)

	return aspFsm
}

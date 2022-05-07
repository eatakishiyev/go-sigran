package m3ua

import "sync"

type m3uaStack struct {
	ApplicationServers         []*As
	ApplicationServerProcesses []*Asp
}

var instance *m3uaStack
var m = &sync.Mutex{}

func GetM3UAStackInstance() *m3uaStack {
	m.Lock()
	defer m.Unlock()

	if instance == nil {
		instance = &m3uaStack{}
	}
	return instance
}

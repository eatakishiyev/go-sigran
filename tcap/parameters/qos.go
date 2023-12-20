package parameters

import "go-sigtran/sccp"

type QoS struct {
	MessageHandling sccp.MessageHandling
	SequenceNumber  uint
	SequenceControl bool
}

package sccp

import "go-sigtran/sccp/parameters"

type MessageHandling uint

const (
	NoSpecialOptions     MessageHandling = 0
	ReturnMessageOnError MessageHandling = 8
)

type ErrorReason uint

const (
	NoTranslationForAddressSuchNature ErrorReason = iota
	NoTranslationForSpecificAddress
	SubsystemCongestion
	SubsystemFailure
	UnequippedUser
	MtpFailure
	NetworkCongestion
	Unqualified
	ErrorInMessageTransfer
	ErrorInLocalProcessing
	DestinationCannotPerformReassembly
	SccpFailure
	HopCounterViolation
	SegmentationNotSupported
	SegmentationFailure
)

type SccpUser interface {
	OnMessage(calledParty *parameters.SccpAddress, callingParty *parameters.SccpAddress, userData []byte, sequenceControl bool, sequenceNumber uint,
		messageHandling MessageHandling)
	OnNotice(calledParty *parameters.SccpAddress, callingParty *parameters.SccpAddress, userData []byte, errorReason ErrorReason, importance uint)
}

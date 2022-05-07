package sccp

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
	OnMessage(calledParty SccpAddress, callingParty SccpAddress, userData []byte, sequenceControl bool, sequenceNumber uint,
		messageHandling MessageHandling)
	OnNotice(calledParty SccpAddress, callingParty SccpAddress, userData []byte, errorReason ErrorReason, importance uint)
}

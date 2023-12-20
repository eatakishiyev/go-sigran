package operations

type Operation uint

const (
	TCInvoke Operation = iota
	TCCancel
	TCLReject
	TCRReject
	TCTimerReset
	TCUCancel
	TCUError
	TCUReject
	TCReturnResultLast
	TCReturnResultNotLast
	TCUAbort
)

type Operations interface {
	GetOperationType() Operation
}

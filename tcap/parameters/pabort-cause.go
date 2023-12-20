package parameters

type PAbortCause uint

const (
	UnrecognizedMessageType PAbortCause = iota
	UnrecognizedTransactionId
	BadlyFormattedTransactionPortion
	IncorrectTransactionPortion
	ResourceLimitation
	AbnormalDialog
	NoCommonDialogPortion
)

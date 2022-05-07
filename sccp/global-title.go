package sccp

type GlobalTitleIndicator byte

const (
	NoGlobalTitleIncluded GlobalTitleIndicator = iota
	NatureOfAddressIndOnly
	TranslationTypeOnly
	TranslationTypeNpEnc
	TranslationTypeNpEncNatureOfAddressInd
)

type GlobalTitle interface {
	GetGlobalTitleIndicator() *GlobalTitleIndicator
}

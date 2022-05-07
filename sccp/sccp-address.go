package sccp

type SccpAddress struct {
	*AddressIndicator
	*SubSystemNumber
	GlobalTitle
}

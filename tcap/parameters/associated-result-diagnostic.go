package parameters

type DialogueServiceUserDiagnostic struct {
	Value ServiceUserDiagnostic
}

type DialogueServiceProviderDiagnostic struct {
	Value ServiceProviderDiagnostic
}

type ServiceProviderDiagnostic int

const (
	ProviderDiagnosticNull ServiceProviderDiagnostic = iota
	ProviderDiagnosticNoReasonGiven
	ProviderDiagnosticNoCommonDialogPortion
)

type ServiceUserDiagnostic int

const (
	UserDiagnosticNull ServiceUserDiagnostic = iota
	UserDiagnosticNoReasonGive
	UserDiagnosticAcnNotSupported
)

type AssociateSourceDiagnostic struct {
	Diagnostic interface{} `asn1:"choice:associate-source-diagnostic"`
}

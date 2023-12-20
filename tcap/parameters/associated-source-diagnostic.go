package parameters

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
	ServiceProviderDiagnostic ServiceProviderDiagnostic
	ServiceUserDiagnostic     ServiceUserDiagnostic
}

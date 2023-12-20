package parameters

type ServiceProviderDiagnostic int

const (
	ProviderDiagnosticNull                  ServiceProviderDiagnostic = 0
	ProviderDiagnosticNoReasonGiven         ServiceProviderDiagnostic = 1
	ProviderDiagnosticNoCommonDialogPortion ServiceProviderDiagnostic = 2
)

type ServiceUserDiagnostic int

const (
	UserDiagnosticNull            ServiceUserDiagnostic = 0
	UserDiagnosticNoReasonGive    ServiceUserDiagnostic = 1
	UserDiagnosticAcnNotSupported ServiceUserDiagnostic = 2
)

type AssociateSourceDiagnostic struct {
	Diagnostic DiagnosticChoice `asn1:"tag:1"`
}

type DiagnosticChoice struct {
	DiagnosticChoice interface{} `asn1:"choice:associate-source-diagnostic"`
}

type ServiceProviderAssociateSourceDiagnostic struct {
	ServiceProviderDiagnostic ServiceProviderDiagnostic
}

type ServiceUserAssociateSourceDiagnostic struct {
	ServiceUserDiagnostic ServiceUserDiagnostic
}

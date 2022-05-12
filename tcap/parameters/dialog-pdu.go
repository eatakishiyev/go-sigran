package parameters

type DialogPDUType uint

const (
	DialogRequestApdu  DialogPDUType = 0
	DialogResponseApdu DialogPDUType = 1
	DialogAbortApdu    DialogPDUType = 4
)

type DialogPDU interface {
	GetDialogType() DialogPDUType
	//Decode(data []byte, params string) ([]byte, error)
}

/*=========================================================
  = 				DialogAbort PDU					  =
 =========================================================*/

type DialogAbortPDU struct {
	AbortSource     AbortSource `asn1:"tag:0"`
	UserInformation []byte      `asn1:"tag:30,optional"`
}

func (d *DialogAbortPDU) GetDialogType() DialogPDUType {
	return DialogAbortApdu
}

/*=========================================================
  = 				DialogRequest PDU					  =
 =========================================================*/

type DialogRequestPDU struct {
	ProtocolVersion    []byte             `asn1:"tag:0,optional"`
	ApplicationContext ApplicationContext `asn1:"tag:1"`
	UserInformation    []byte             `asn1:"tag:30,optional"`
}

func (d *DialogRequestPDU) GetDialogType() DialogPDUType {
	return DialogRequestApdu
}

/*=========================================================
  = 				DialogResponse PDU					  =
 =========================================================*/

type DialogResponsePDU struct {
	ProtocolVersion           []byte                    `asn1:"tag:0,optional"`
	ApplicationContext        ApplicationContext        `asn1:"tag:1"`
	AssociateResult           AssociateResult           `asn1:"tag:2"`
	AssociateSourceDiagnostic AssociateSourceDiagnostic `asn1:"tag:3"`
	UserInformation           []byte                    `asn1:"tag:30,optional"`
}

func (d *DialogResponsePDU) GetDialogType() DialogPDUType {
	return DialogResponseApdu
}

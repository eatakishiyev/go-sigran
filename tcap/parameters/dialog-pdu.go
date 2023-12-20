package parameters

type DialogPDUType uint

const (
	RequestPDU  DialogPDUType = 0
	ResponsePDU DialogPDUType = 1
	AbortPDU    DialogPDUType = 4
)

type DialogPDU interface {
	GetDialogPDUType() DialogPDUType
	GetUserInformationBytes() []byte
}

/*=========================================================
  = 				DialogAbort PDU					  =
 =========================================================*/

type DialogAbortPDU struct {
	AbortSource          AbortSource `asn1:"tag:0"`
	UserInformationBytes []byte      `asn1:"tag:30,optional"`
}

func (dialogPdu *DialogAbortPDU) GetDialogPDUType() DialogPDUType {
	return AbortPDU
}

func (dialogPdu *DialogAbortPDU) GetUserInformationBytes() []byte {
	return dialogPdu.UserInformationBytes
}

/*=========================================================
  = 				DialogRequest PDU					  =
 =========================================================*/

type DialogRequestPDU struct {
	ProtocolVersion      []byte             `asn1:"tag:0,optional"`
	ApplicationContext   ApplicationContext `asn1:"tag:1"`
	UserInformationBytes []byte             `asn1:"tag:30,optional"`
}

func (dialogPdu *DialogRequestPDU) GetDialogPDUType() DialogPDUType {
	return RequestPDU
}

func (dialogPdu *DialogRequestPDU) GetUserInformationBytes() []byte {
	return dialogPdu.UserInformationBytes
}

func (dialogPdu *DialogRequestPDU) IsProtocolVersionCorrect() bool {
	if dialogPdu.ProtocolVersion == nil {
		return false
	}

	return dialogPdu.ProtocolVersion[1]&0b10000000 == 0b10000000
}

/*=========================================================
  = 				DialogResponse PDU					  =
 =========================================================*/

type DialogResponsePDU struct {
	ProtocolVersion           []byte                    `asn1:"tag:0,optional"`
	ApplicationContext        ApplicationContext        `asn1:"tag:1"`
	AssociateResult           int                       `asn1:"tag:2"`
	AssociateSourceDiagnostic AssociateSourceDiagnostic `asn1:"tag:3,choice:associate-source-diagnostic"`
	UserInformationBytes      []byte                    `asn1:"tag:30,optional"`
}

func (dialogPdu *DialogResponsePDU) GetDialogPDUType() DialogPDUType {
	return ResponsePDU
}

func (dialogPdu *DialogResponsePDU) GetUserInformationBytes() []byte {
	return dialogPdu.UserInformationBytes
}

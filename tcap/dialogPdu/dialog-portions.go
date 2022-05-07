package dialogPdu

import (
	"encoding/asn1"
	"go-sigtran/tcap/parameters"
)

type DialogPDUType uint

const (
	DialogRequestApdu  DialogPDUType = 0
	DialogResponseApdu DialogPDUType = 1
	DialogAbortApdu    DialogPDUType = 4
)

type DialogPDU interface {
	GetDialogType() DialogPDUType
	GetUserInformation() *parameters.UserInformation
	SetUserInformation(information *parameters.UserInformation)
	Decode(data []byte, params string) ([]byte, error)
}

type DialogRequestPDU struct {
	ProtocolVersion    asn1.BitString             `asn1:"tag:0,optional"`
	ApplicationContext asn1.ObjectIdentifier      `asn1:"tag:1"`
	UserInformation    parameters.UserInformation `asn1:"tag:30,optional"`
}

func (d *DialogRequestPDU) GetDialogType() DialogPDUType {
	return DialogRequestApdu
}

func (d *DialogRequestPDU) GetUserInformation() *parameters.UserInformation {
	return nil
}

func (d *DialogRequestPDU) SetUserInformation(information *parameters.UserInformation) {

}

func (d *DialogRequestPDU) Decode(data []byte, params string) ([]byte, error) {
	return asn1.UnmarshalWithParams(data, d, params)
}

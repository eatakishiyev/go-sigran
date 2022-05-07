package m3ua

import (
	"go-sigtran/m3ua/messages/transfer"
	"go-sigtran/m3ua/parameters"
)

type M3UAUser interface {
	ServiceIdentification() parameters.ServiceIndicator
	OnMtpTransferIndication(data *transfer.PayloadData)
	MtpTransferRequest(data *transfer.PayloadData)
	OnMtpPause(dpc int)
	OnMtpResume(dpc int)
	OnMtpStatusWithCongestionLevel(dpc int, congestionLevel int)
	OnMtpStatusWithCause(dpc int, cause int)
}

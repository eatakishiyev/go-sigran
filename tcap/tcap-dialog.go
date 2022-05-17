package tcap

import (
	"fmt"
	sccp_parameters "go-sigtran/sccp/parameters"
	tcap_parameters "go-sigtran/tcap/parameters"

	"sync"
)

type DialogState byte

const (
	Idle DialogState = iota
	InitiationReceived
	InitiationSent
	Active
)

type TcapDialog struct {
	CallingParty   *sccp_parameters.SccpAddress
	CalledParty    *sccp_parameters.SccpAddress
	State          DialogState
	lock           sync.Mutex
	DialogId       int64
	RemoteDialogId int64
	localInitiated bool
}

func (d *TcapDialog) beginMessage(beginMessage *BeginMessage, callingParty *sccp_parameters.SccpAddress,
	calledParty *sccp_parameters.SccpAddress) {
	/* ITU Q.774 Page 12
	   * The receiving transaction sub-layer stores the received Originating Address as the destination address
	   for this transaction and its own address as the Originating Address (from memory or from the
	   destination address of the received N-UNITDATA indication primitive). The TC-user receives a
	   TC-BEGIN indication primitive containing the received Destination Address and Originating
	   Address.
	*/
	d.lock.Lock()
	defer d.lock.Unlock()

	d.RemoteDialogId = beginMessage.OriginatingDialogId

	d.CallingParty = callingParty
	d.CalledParty = calledParty

	if beginMessage.DialogPortion.FullBytes != nil &&
		len(beginMessage.DialogPortion.FullBytes) > 0 {
		dp := tcap_parameters.DialogPortion{}
		dp.Decode(beginMessage.DialogPortion.Bytes)
		if dp.DialogPDU.GetDialogType() != tcap_parameters.DialogRequestApdu {
			fmt.Printf("begin message unexpected dialog pdu received. expecting DialogRequestApdu but received %s",
				dp.DialogPDU.GetDialogType())

			abortPdu := tcap_parameters.DialogAbortPDU{
				AbortSource: tcap_parameters.DialogServiceProvider,
			}

			dialogPortion := tcap_parameters.DialogPortion{
				DialogPDU: &abortPdu,
			}
			dialogPortion.Encode()

			return
		}

		d.State = InitiationReceived
	}
}

func (d *TcapDialog) continueMessage(continueMessage *ContinueMessage, callingParty *sccp_parameters.SccpAddress,
	calledParty *sccp_parameters.SccpAddress) {
	d.lock.Lock()
	defer d.lock.Unlock()

	if continueMessage.DialogPortion.FullBytes != nil &&
		len(continueMessage.DialogPortion.FullBytes) > 0 {
		dp := tcap_parameters.DialogPortion{}
		dp.Decode(continueMessage.DialogPortion.Bytes)

		if dp.DialogPDU.GetDialogType() != tcap_parameters.DialogResponseApdu {

		}
	}
}

func (d *TcapDialog) abortMessage(abortMessage *AbortMessage, party *sccp_parameters.SccpAddress, party2 *sccp_parameters.SccpAddress) {
	if abortMessage.UAbortCause.FullBytes != nil &&
		len(abortMessage.UAbortCause.FullBytes) > 0 {
		dp := tcap_parameters.DialogPortion{}
		dp.Decode(abortMessage.UAbortCause.Bytes)

		if dp.DialogPDU.GetDialogType() != tcap_parameters.DialogAbortApdu {

		}
	}
}

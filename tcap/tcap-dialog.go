package tcap

import (
	sccp_parameters "go-sigtran/sccp/parameters"
	"go-sigtran/tcap/operations"
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
	CallingParty       *sccp_parameters.SccpAddress
	CalledParty        *sccp_parameters.SccpAddress
	State              DialogState
	lock               sync.Mutex
	DialogId           int
	RemoteDialogId     int
	localInitiated     bool
	ApplicationContext *tcap_parameters.ApplicationContext
	QoS                *tcap_parameters.QoS
	TcapStack          *TcapStack
}

//func (dialog *TcapDialog) beginMessage(beginMessage *BeginMessage, callingParty *sccp_parameters.SccpAddress,
//	calledParty *sccp_parameters.SccpAddress) {
//	/* ITU Q.774 Page 12
//	   * The receiving transaction sub-layer stores the received Originating Address as the destination address
//	   for this transaction and its own address as the Originating Address (from memory or from the
//	   destination address of the received N-UNITDATA indication primitive). The TC-user receives a
//	   TC-BEGIN indication primitive containing the received Destination Address and Originating
//	   Address.
//	*/
//	dialog.lock.Lock()
//	defer dialog.lock.Unlock()
//
//	dialog.RemoteDialogId = beginMessage.OriginatingTransactionId
//
//	dialog.CallingParty = callingParty
//	dialog.CalledParty = calledParty
//
//	if beginMessage.DialogPortionBytes != nil &&
//		len(beginMessage.DialogPortionBytes) > 0 {
//		dialogPortion := tcap_parameters.DialogPortion{}
//		ok := dialogPortion.Decode(beginMessage.DialogPortionBytes)
//
//		if !ok || (ok && dialogPortion.DialogPDU.GetDialogPDUType() != tcap_parameters.RequestPDU) {
//			//Send abort message
//			abort := tcap_parameters.DialogAbortPDU{
//				AbortSource: tcap_parameters.DialogServiceProvider,
//			}
//			dialogPortion = tcap_parameters.DialogPortion{
//				DialogPDU: &abort,
//				Oid:       DialogAsId(),
//			}
//			return
//		}
//
//		dialogRequestPdu := dialogPortion.DialogPDU.(*tcap_parameters.DialogRequestPDU)
//		if dialogRequestPdu.IsProtocolVersionCorrect() {
//			dialog.ApplicationContext = dialogRequestPdu.ApplicationContext
//			tcBeginIndication := primitives.TCBegin{
//				ApplicationContext: dialog.ApplicationContext,
//				TCAPDialog:         dialog,
//				UserInformation:    dialogRequestPdu.GetUserInformationBytes(),
//				Components:         processComponents(beginMessage.ComponentsBytes),
//			}
//			dialog.State = InitiationReceived
//
//			//dialog.TcapStack.OnBegin(tcBeginIndication)
//		} else {
//			//Protocol version is not correct. Abort message
//			diagnostic := tcap_parameters.AssociateSourceDiagnostic{
//				ServiceProviderDiagnostic: tcap_parameters.ProviderDiagnosticNoCommonDialogPortion,
//			}
//
//			associateResult := tcap_parameters.AssociateResult{Result: tcap_parameters.RejectPermanent}
//			dialogResponse := tcap_parameters.DialogResponsePDU{
//				ApplicationContext:        dialogRequestPdu.ApplicationContext,
//				AssociateResult:           &associateResult,
//				AssociateSourceDiagnostic: &diagnostic,
//			}
//
//			dialogPortion = tcap_parameters.DialogPortion{
//				Oid:       DialogAsId(),
//				DialogPDU: &dialogResponse,
//			}
//
//			dialogPortionBytes := dialogPortion.Encode()
//			abortMessage := AbortMessage{
//				UAbortCauseBytes:         dialogPortionBytes,
//				DestinationTransactionId: beginMessage.OriginatingTransactionId,
//			}
//
//			//abortMessageBytes := abortMessage.Encode()
//			//dialog.TcapStack.Send(callingParty, calledParty, dialog.QoS, abortMessageBytes)
//		}
//	} else {
//		//tcBegin := primitives.TCBegin{
//		//	TCAPDialog: dialog,
//		//	Components: processComponents(beginMessage.ComponentsBytes),
//		}
//		dialog.State = InitiationReceived
//		//dialog.TcapStack.OnBegin(tcBegin)
//	}
//}

func (dialog *TcapDialog) continueMessage(continueMessage *ContinueMessage, callingParty *sccp_parameters.SccpAddress,
	calledParty *sccp_parameters.SccpAddress) {
	dialog.lock.Lock()
	defer dialog.lock.Unlock()

	//if continueMessage.DialogPortion != nil &&
	//	len(continueMessage.DialogPortion) > 0 {
	//	external, ok := tcap_parameters.DecodeExternal(continueMessage.DialogPortion)
	//	dialogPortion, ok := tcap_parameters.DecodeDialogPortionFromExternal(external)
	//
	//	fmt.Printf("DialogPortion %s", dialogPortion)
	//}
}

func (dialog *TcapDialog) abortMessage(abortMessage *AbortMessage, callingParty *sccp_parameters.SccpAddress,
	calledParty *sccp_parameters.SccpAddress) {
	//if abortMessage.UAbortCause != nil &&
	//	len(abortMessage.UAbortCause) > 0 {
	//	dp := tcap_parameters.DialogPortion{}
	//	//dp.Decode(abortMessage.UAbortCause)
	//
	//	if dp.DialogPDU.GetDialogType() != tcap_parameters.DialogAbortApdu {
	//
	//	}
	//}
}

func processComponents(componentBytes []byte) []operations.Operations {
	return nil
}

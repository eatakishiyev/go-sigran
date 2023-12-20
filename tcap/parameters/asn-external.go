package parameters

type ExternalPDU struct {
	Oid       []byte      `asn1:"tag:6,universal,optional"`
	DialogPdu interface{} `asn1:"tag:0,explicit,choice:dialogPortion"`
}

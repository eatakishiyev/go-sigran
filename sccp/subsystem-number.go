package sccp

type SubSystemNumber uint

const (
	SsnNotUsed     SubSystemNumber = 0
	SccpManagement SubSystemNumber = 1
	Hlr            SubSystemNumber = 6
	Vlr            SubSystemNumber = 7
	Msc            SubSystemNumber = 8
	Eir            SubSystemNumber = 9
	Auc            SubSystemNumber = 10
	Reserved108    SubSystemNumber = 108
	Ranap          SubSystemNumber = 142
	Rnsap          SubSystemNumber = 143
	Gmlc           SubSystemNumber = 145
	Cap            SubSystemNumber = 146
	GsmScf         SubSystemNumber = 147
	Siwf           SubSystemNumber = 148
	Sgsn           SubSystemNumber = 149
	Ggsn           SubSystemNumber = 150
	Pcap           SubSystemNumber = 249
	BssapBsc       SubSystemNumber = 250
	BssApMsc       SubSystemNumber = 251
	BssapSmlc      SubSystemNumber = 252
	BssOm          SubSystemNumber = 253
	AInterface     SubSystemNumber = 254
)

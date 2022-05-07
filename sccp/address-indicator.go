package sccp

/*
 * Encoding Address Indicator
 * |----8----|----7----|----6----|----5----|----4----|----3----|----2----|----1----|
 * |Reserved | Routing | 		 |  SSN    |PointCode|		   |for nat. |Indicator| Global
 * Title Indicator |Indicator|Indicator| |use | | | | |
 * |---------|---------|---------------------------------------|---------|---------|
 */

type RoutingIndicator byte

const (
	RouteOnGt RoutingIndicator = iota
	RouteOnSsn
)

type AddressIndicator struct {
	RoutingIndicator     RoutingIndicator
	GlobalTitleIndicator GlobalTitleIndicator
	SsnIndicator         bool
	PointCodeIndicator   bool
}

func (ai *AddressIndicator) Encode() []byte {
	var addressIndicator = byte(0)
	addressIndicator |= 0 //Reserved
	addressIndicator <<= 1
	addressIndicator |= byte(ai.RoutingIndicator) //Routing Indicator
	addressIndicator <<= 4
	addressIndicator |= byte(ai.GlobalTitleIndicator) //Global Title Indicator
	addressIndicator <<= 1
	addressIndicator |= bool2Byte(ai.SsnIndicator)
	addressIndicator <<= 1
	addressIndicator |= bool2Byte(ai.PointCodeIndicator)
	return []byte{addressIndicator}
}

func (ai *AddressIndicator) Decode(data []byte) {
	var addressIndicator = data[0]
	var _ = addressIndicator >> 7                                                  //Reserved
	ai.RoutingIndicator = RoutingIndicator((addressIndicator & 0x40) >> 6)         //Routing Indicator
	ai.GlobalTitleIndicator = GlobalTitleIndicator((addressIndicator & 0x3C) >> 2) //Global Title Indicator
	ai.SsnIndicator = ((addressIndicator & 0x02) >> 1) == 1                        //SSN Indicator
	ai.PointCodeIndicator = (addressIndicator & 0x01) == 1                         //Point Code Indicator
}

func bool2Byte(b bool) byte {
	if b {
		return 1
	} else {
		return 0
	}
}

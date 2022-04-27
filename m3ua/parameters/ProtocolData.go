package parameters

import "encoding/binary"

type ProtocolData struct {
	*ParameterHeader
	UserProtocolData []byte
	Opc              uint32
	Dpc              uint32
	Sls              byte
	Ni               byte
	Mp               byte
	Si               byte
}

func (pd *ProtocolData) EncodeParameter() []byte {
	var encoded []byte
	binary.BigEndian.PutUint32(encoded, pd.Opc)
	binary.BigEndian.PutUint32(encoded, pd.Dpc)
	encoded = append(encoded, pd.Si)
	encoded = append(encoded, pd.Ni)
	encoded = append(encoded, pd.Mp)
	encoded = append(encoded, pd.Sls)
	encoded = append(encoded, pd.UserProtocolData...)
	return encoded
}

func (pd *ProtocolData) DecodeParameter(p []byte) {
	pd.Opc = binary.BigEndian.Uint32(p[0:4])
	pd.Dpc = binary.BigEndian.Uint32(p[4:8])
	pd.Si = p[8]
	pd.Ni = p[9]
	pd.Mp = p[10]
	pd.Sls = p[11]
	pd.UserProtocolData = p[12:]
}

func (pd *ProtocolData) GetHeader() *ParameterHeader {
	return pd.ParameterHeader
}

func (pd *ProtocolData) SetHeader(header *ParameterHeader) {
	pd.ParameterHeader = header
}

func NewProtocolData(userProtocolData []byte, opc uint32, dpc uint32, sls byte, ni byte, mp byte, si byte) *ProtocolData {
	return &ProtocolData{
		ParameterHeader: &ParameterHeader{
			Tag: ParamProtocolData,
		},
		UserProtocolData: userProtocolData,
		Opc:              opc,
		Dpc:              dpc,
		Sls:              sls,
		Ni:               ni,
		Mp:               mp,
		Si:               si,
	}
}

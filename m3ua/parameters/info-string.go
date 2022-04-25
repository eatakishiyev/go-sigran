package parameters

type InfoString struct {
	*ParameterHeader
	InfoString string
}

func NewInfoString(infoString string) *InfoString {
	is := &InfoString{
		ParameterHeader: &ParameterHeader{
			Tag: ParamInfoString,
		},
		InfoString: infoString,
	}

	return is
}

func (i *InfoString) EncodeParameter() []byte {
	infoStringBytes := []byte(i.InfoString)
	data := make([]byte, len(infoStringBytes))
	return append(data)
}

func (i *InfoString) DecodeParameter(p []byte) {
	i.InfoString = string(p)
}

func (i *InfoString) GetHeader() *ParameterHeader {
	return i.ParameterHeader
}

func (i *InfoString) SetHeader(header *ParameterHeader) {
	i.ParameterHeader = header
}

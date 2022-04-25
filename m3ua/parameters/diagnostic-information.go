package parameters

type DiagnosticInformation struct {
	*ParameterHeader
	Information []byte
}

func NewDiagnosticInformation(information []byte) *DiagnosticInformation {
	return &DiagnosticInformation{Information: information,
		ParameterHeader: &ParameterHeader{
			Tag: ParamDiagnosticInformation,
		},
	}
}

func (d *DiagnosticInformation) EncodeParameter() []byte {
	var encoded []byte
	encoded = append(encoded, d.Information...)
	return encoded
}

func (d *DiagnosticInformation) DecodeParameter(p []byte) {
	d.Information = p
}

func (d *DiagnosticInformation) GetHeader() *ParameterHeader {
	return d.ParameterHeader
}

func (d *DiagnosticInformation) SetHeader(header *ParameterHeader) {
	d.ParameterHeader = header
}

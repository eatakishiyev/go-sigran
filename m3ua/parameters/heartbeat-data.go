package parameters

type HeartbeatData struct {
	*ParameterHeader
	Payload []byte
}

func NewHeartbeatData(paylaod []byte) *HeartbeatData {
	return &HeartbeatData{
		ParameterHeader: &ParameterHeader{
			Tag: ParamHeartBeadData,
		},
		Payload: paylaod,
	}
}

func (h *HeartbeatData) EncodeParameter() []byte {
	var encoded []byte
	return append(encoded, h.Payload...)
}

func (h *HeartbeatData) DecodeParameter(p []byte) {
	h.Payload = p
}

func (h *HeartbeatData) GetHeader() *ParameterHeader {
	return h.ParameterHeader
}

func (h *HeartbeatData) SetHeader(header *ParameterHeader) {
	h.ParameterHeader = header
}

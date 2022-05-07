package parameters

type ParameterHeader struct {
	Tag    ParameterTag //2 byte
	Length uint16       //2 byte
}

func NewHeader(tag ParameterTag, length uint16) *ParameterHeader {
	return &ParameterHeader{
		Tag:    tag,
		Length: length,
	}
}

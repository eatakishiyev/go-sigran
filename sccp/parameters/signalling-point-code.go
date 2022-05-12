package parameters

type SignallingPointCode struct {
	PointCode uint
}

func (spc *SignallingPointCode) Encode() []byte {
	return nil
}

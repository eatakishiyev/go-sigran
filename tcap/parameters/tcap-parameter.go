package parameters

type TCAPParameter interface {
	Decode(data []byte) ([]byte, error)
}

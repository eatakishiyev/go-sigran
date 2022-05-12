package parameters

type Result int

const (
	Accepted Result = iota
	RejectPermanent
)

type AssociateResult struct {
	Result Result
}

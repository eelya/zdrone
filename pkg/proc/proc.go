package proc

// Proc is
type Proc interface {
	Add(int) int
	Mul(int) int
}

type procImpl struct {
	factor int
}

func (p procImpl) Add(x int) int {
	return x + p.factor
}

func (p procImpl) Mul(x int) int {
	return x * p.factor
}

// NewProc creates
func NewProc(f int) Proc {
	return &procImpl{
		factor: f,
	}
}

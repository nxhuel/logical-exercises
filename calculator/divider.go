package calculator

type Divider interface {
	Divide(a, b float64) (float64, error)
}

type divider struct{}

func NewDivider() Divider {
	return &divider{}
}

func (d *divider) Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, ErrorDivideByZero
	}
	return a / b, nil
}

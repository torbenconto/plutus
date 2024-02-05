package indicators

type Indicator interface {
	Calculate() (float64, error)
}

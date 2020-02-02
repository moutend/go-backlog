package types

type Hours float64

func NewHours(f float64) *Hours {
	h := Hours(f)

	return &h
}

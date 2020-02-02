package types

import "fmt"

type Hours float64

func (h *Hours) String() string {
	if h == nil {
		return ""
	}

	return fmt.Sprint(*h)
}

func NewHours(f float64) *Hours {
	h := Hours(f)

	return &h
}

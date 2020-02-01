package types

import (
	"time"
)

type Date string

func (d Date) Time() time.Time {
	t, _ := time.Parse(time.RFC3339, string(d))

	return t
}

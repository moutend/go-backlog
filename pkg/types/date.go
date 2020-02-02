package types

import (
	"time"
)

type Date string

func (d *Date) Time() time.Time {
	if d == nil {
		return time.Time{}
	}
	if t, err := time.Parse("2006-01-02", string(*d)); err == nil {
		return t
	}
	if t, err := time.Parse(time.RFC3339, string(*d)); err == nil {
		return t
	}

	return time.Time{}
}

func NewDate(s string) *Date {
	d := Date(s)

	return &d
}

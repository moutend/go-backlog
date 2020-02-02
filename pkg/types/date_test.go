package types

import (
	"testing"
	"time"
)

func TestDate_empty(t *testing.T) {
	d1 := NewDate("")

	if !d1.Time().IsZero() {
		t.Fatalf("Must return zero value")
	}
}

func TestDate_invalid(t *testing.T) {
	d1 := NewDate("abcdef")

	if !d1.Time().IsZero() {
		t.Fatalf("Must return zero value")
	}
}
func TestDate_pattern1(t *testing.T) {
	t1, _ := time.Parse("2006-01-02", "2006-01-02")
	d1 := NewDate("2006-01-02")

	if !d1.Time().Equal(t1) {
		t.Fatalf("Failed to parse")
	}
}

func TestDate_rfc3339(t *testing.T) {
	t1, _ := time.Parse(time.RFC3339, "2006-01-02T15:04:05Z")
	d1 := NewDate("2006-01-02T15:04:05Z")

	if !d1.Time().Equal(t1) {
		t.Fatalf("Failed to parse")
	}
}

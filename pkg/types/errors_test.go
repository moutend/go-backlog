package types

import "testing"

func TestErrorsHas(t *testing.T) {
	e1 := Error{
		Code:     1,
		Message:  "message1",
		MoreInfo: "more info1",
	}
	e2 := Error{
		Code:     2,
		Message:  "message2",
		MoreInfo: "more info2",
	}
	es := Errors{
		Errors: []Error{e1, e2},
	}

	if es.Has(ErrorUnexpected) {
		t.Fatalf("must not contain ErrorUnexpected")
	}
	if !es.Has(ErrorInternal) {
		t.Fatalf("must contain ErrorInternal")
	}
	if !es.Has(ErrorLicense) {
		t.Fatalf("must contain ErrorLicense")
	}
}

func TestErrorsError(t *testing.T) {
	e1 := Error{
		Code:     1,
		Message:  "message1",
		MoreInfo: "more info1",
	}
	e2 := Error{
		Code:     2,
		Message:  "message2",
		MoreInfo: "more info2",
	}
	es := Errors{
		Errors: []Error{e1, e2},
	}

	expected := "internal error: message1 (more info1); license error: message2 (more info2)"
	actual := es.Error()

	if actual != expected {
		t.Fatalf("actual: %q\nexpected: %q", actual, expected)
	}
}

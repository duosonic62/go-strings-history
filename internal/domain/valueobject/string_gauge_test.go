package valueobject

import "testing"

func TestNewStringGauge_Positive(t *testing.T) {
	actual, err := NewStringGauge(9, 42)
	if err != nil {
		t.Error("expected err is nil but actual ", err)
	}
	if actual.ThinGauge() != 9 {
		t.Error("expected 9 but actual: " + string(actual.ThinGauge()))
	}
	if actual.ThickGauge() != 42 {
		t.Error("expected 9 but actual: " + string(actual.ThickGauge()))
	}
}

func TestNewStringGauge_NegativeReverseStringGauge(t *testing.T) {
	_, err := NewStringGauge(42, 9)
	if err == nil {
		t.Error("expected err is not nil but actual ", err)
	}
}
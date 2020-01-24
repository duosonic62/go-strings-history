package factoryimple

import "testing"

func TestStringFactoryImpl_NewString(t *testing.T) {
	factory := NewStringFactory(NewIdFactory())
	actual, _ := factory.NewString(9, 42, "test_name", "test_description", "test_maker", "test_url")
	if actual.GetName() != "test_name" {
		t.Error("Expected: test_name, but " + actual.GetName())
	}
	if actual.GetDescription() != "test_description" {
		t.Error("Expected: test_description, but " + actual.GetDescription())
	}
	if actual.GetMaker() != "test_maker" {
		t.Error("Expected: test_maker, but " + actual.GetMaker())
	}
	if actual.GetGauge().ThinGauge() != 9 {
		t.Error("Expected: 9, but " + string(actual.GetGauge().ThinGauge()))
	}
	if actual.GetGauge().ThickGauge() != 42 {
		t.Error("Expected: 42, but " + string(actual.GetGauge().ThickGauge()))
	}
}
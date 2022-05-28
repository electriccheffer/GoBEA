package ITA

import "testing"

func TestITA_RequestType(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null.")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNullCase should be JSON")
	}

}

func TestITA_Year(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "2012",
		requestType:   "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed nonNullCase year should be set to 2012")
	}

}

func TestITA_Frequency(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be null.")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "moderate",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed nonNullcase should be moderate")
	}
}

func TestITA_AreaOrCountry(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "usa",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.AreaOrCountry() != "usa" {
		t.Error("Test failed nonNullCase should be set to usa")
	}

}

func TestITA_Indicator(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Indicator() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "moderate",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed nonNullCase should be set to moderate")
	}

}

func TestITA_SetAreaOrCountry(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetAreaOrCountry("usa")
	if nullCase.AreaOrCountry() != "usa" {
		t.Error("Test failed areaOrCountry should be set to usa")
	}
}

func TestITA_SetFrequency(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be null.")
	}
	nullCase.SetFrequency("moderate")
	if nullCase.Frequency() != "moderate" {
		t.Error("Test failed nullCase frequency should be moderate")
	}
}

func TestITA_SetIndicator(t *testing.T) {

	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Indicator() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nullCase.SetIndicator("positive")
	if nullCase.Indicator() != "positive" {
		t.Error("Test failed nullCase indicator should be set to positive.")
	}
}

func TestITA_SetRequestType(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be set to null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("nullCase should be set to JSON")
	}
}

func TestITA_SetYear(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be set to null.")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed nullCase should be set to 2012")
	}
}

func TestITA_String(t *testing.T) {
	firstCase := ITA{
		indicator:     "positive",
		areaOrCountry: "usa",
		frequency:     "moderate",
		year:          "2012",
		requestType:   "JSON",
	}
	if firstCase.String() != "positive usa moderate 2012 JSON" {
		t.Error("Test failed string form is not correct")
	}
}

func TestNewITA(t *testing.T) {

	constructorCase := NewITA("positive", "usa", "moderate", "2012", "JSON")
	if constructorCase.Indicator() != "positive" && constructorCase.AreaOrCountry() != "usa" && constructorCase.Frequency() != "moderate" && constructorCase.Year() != "2012" && constructorCase.RequestType() != "JSON" {
		t.Error("Test failed for constructor case.")
	}
}

package InternationalInvestmentPosition

import "testing"

func TestIIPRequest_Year(t *testing.T) {

	nullCase := IIPRequest{
		typeOfInvestment: "static",
		component:        "none",
		frequency:        "none",
		year:             "",
		returnFormat:     "none",
	}
	if nullCase.Year() != "" {
		t.Error("Year test 1 failed.  Year should be empty.")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "2012",
		returnFormat:     "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Year test 2 failed. Year should be 2012")
	}
}

func TestIIPRequest_Frequency(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.Frequency() != "" {
		t.Error("Frequency test null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "non-frequent",
		year:             "",
		returnFormat:     "",
	}
	if nonNullCase.Frequency() != "non-frequent" {
		t.Error("Frequency test non-null case false")
	}

}

func TestIIPRequest_Component(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.Component() != "" {
		t.Error("Component test null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "tons",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nonNullCase.Component() != "tons" {
		t.Error("Component test nonNull case should be tons")
	}
}

func TestIIPRequest_TypeOfInvestment(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.TypeOfInvestment() != "" {
		t.Error("Type of investment null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "bad",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	if nonNullCase.TypeOfInvestment() != "bad" {
		t.Error("Test Failed type of investment should be bad")
	}

}

func TestIIPRequest_ReturnFormat(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.ReturnFormat() != "" {
		t.Error("Test Failed returnFormat should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "XML",
	}

	if nonNullCase.ReturnFormat() != "XML" {
		t.Error("Test Failed Return Format should be set to XML")
	}

}

func TestIIPRequest_SetComponent(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetComponent("JSON")
	if nullCase.Component() != "JSON" {
		t.Error("Test Failed Component should be set to JSON")
	}
}

func TestIIPRequest_SetFrequency(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetFrequency("infrequent")
	if nullCase.Frequency() != "infrequent" {
		t.Error("Test Failed Frequency should be set to infrequent")
	}

}

func TestIIPRequest_SetReturnFormat(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetReturnFormat("JSON")
	if nullCase.ReturnFormat() != "JSON" {
		t.Error("Test Failed return format should be set to JSON")
	}
}

func TestIIPRequest_SetYear(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test Failed Year should be set to 2012")
	}
}

func TestIIPRequest_SetTypeOfInvestment(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetTypeOfInvestment("bad")
	if nullCase.TypeOfInvestment() != "bad" {
		t.Error("Test Failed type of investment should be bad")
	}
}

func TestIIPRequest_String(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "bad",
		component:        "multiple",
		frequency:        "not",
		year:             "2012",
		returnFormat:     "JSON",
	}
	if nullCase.String() != "multiple bad not 2012 JSON" {

	}
}

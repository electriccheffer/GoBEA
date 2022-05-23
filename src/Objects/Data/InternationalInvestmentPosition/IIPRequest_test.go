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

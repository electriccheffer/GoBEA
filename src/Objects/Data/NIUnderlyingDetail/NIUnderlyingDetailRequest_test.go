package NIUnderlyingDetailRequest

import "testing"

func TestNIUnderlyingDetailRequest_Frequency(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}

	if nullCase.Frequency() != "" {
		t.Error("Test failed Frequency should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "moderate",
		year:      "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed Frequency should be moderate")
	}

}

func TestNIUnderlyingDetailRequest_Year(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed year field should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "2012",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed year field should be 2012")
	}

}

func TestNIUnderlyingDetailRequest_TableName(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}

	if nullCase.TableName() != "" {
		t.Error("Test failed Null Case should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "Tayne",
		frequency: "",
		year:      "",
	}

	if nonNullCase.TableName() != "Tayne" {
		t.Error("Test Failed nonNull case tableName should be Tayne")
	}
}

func TestNIUnderlyingDetailRequest_SetFrequency(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed null case should start null")
	}
	nullCase.SetFrequency("moderate")
	if nullCase.Frequency() != "moderate" {
		t.Error("Test failed Frequency should be moderate")
	}

}

func TestNIUnderlyingDetailRequest_SetYear(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should start null")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed year should be 2012")
	}

}

func TestNIUnderlyingDetailRequest_SetTableName(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed tableName should start null")
	}
	nullCase.SetTableName("Tayne")
	if nullCase.TableName() != "Tayne" {
		t.Error("Test failed tableName should be Tayne")
	}

}

func TestNIUnderlyingDetailRequest_RequestType(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed request type should be JSON")
	}

}

func TestNIUnderlyingDetailRequest_SetRequestType(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed nullCase should start as null")
	}

	nullCase.SetRequestType("XML")
	if nullCase.RequestType() != "XML" {
		t.Error("Test failed nullCase should be XML.")
	}
}

func TestNIUnderlyingDetailRequest_String(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "Tayne",
		frequency: "frequent",
		year:      "2012",
	}

	if nullCase.String() != "Tayne frequent 2012" {
		t.Error("Test failed string ouput should be: Tayne frequent 2012")
	}
}

func TestNewNIUnderlyingDetailRequest(t *testing.T) {
	constructCase := NewNIUnderlyingDetailRequest("Tayne", "regular", "2012", "JSON")
	if constructCase.Year() != "2012" {
		t.Error("Test failed constructCase year should be 2012")
	}
	if constructCase.RequestType() != "JSON" {
		t.Error("Test failed constructCase request type should be JSON ")
	}
	if constructCase.TableName() != "Tayne" {
		t.Error("Test failed constructCase tableName should be Tayne")
	}
}

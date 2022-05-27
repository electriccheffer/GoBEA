package FixedAssets

import "testing"

func TestFixedAssetsRequest_Year(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}

	if nullCase.Year() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "2012",
		requestType: "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed nonNull case should be set to 2012")
	}
}

func TestFixedAssetsRequest_RequestType(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null.")
	}
	nonNullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNull case should be set to JSON")
	}
}

func TestFixedAssetsRequest_TableName(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}

	if nullCase.TableName() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nonNullCase := FixedAssetsRequest{
		tableName:   "Tayne",
		year:        "",
		requestType: "",
	}

	if nonNullCase.TableName() != "Tayne" {
		t.Error("Test failed nonNull case should be set to Tayne")
	}
}

func TestFixedAssetsRequest_String(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "Tayne",
		year:        "2012",
		requestType: "JSON",
	}
	if nullCase.String() != "Tayne 2012 JSON" {
		t.Error("Test failed string should read Tayne 2012 JSON")
	}
}

func TestFixedAssetsRequest_SetYear(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed nullcase should be set to 2012")
	}
}

func TestFixedAssetsRequest_SetTableName(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed null case should be set to null")
	}

	nullCase.SetTableName("Tayne")
	if nullCase.TableName() != "Tayne" {
		t.Error("Test failed null case should be set to TAyne")
	}
}

func TestFixedAssetsRequest_SetRequestType(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be set to null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed request type should be set to JSON")
	}
}

func TestNewFixedAssetsRequest(t *testing.T) {
	newCase := NewFixedAssetsRequest("Tayne", "2012", "JSON")
	if newCase.TableName() != "Tayne" {
		t.Error("Test failed tableName should be Tayne")
	}
	if newCase.Year() != "2012" {
		t.Error("Test failed year should be 2012")
	}
	if newCase.RequestType() != "JSON" {
		t.Error("Test failed requestType should be set to JSON.")
	}

}

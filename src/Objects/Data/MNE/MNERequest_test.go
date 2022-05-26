package MNE

import "testing"

func TestMNERequest_Classification(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Classification() != "" {
		t.Error("Test failed classification should me empty")
	}
	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "negative",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.Classification() != "negative" {
		t.Error("Test failed classification should be set to negative")
	}
}

func TestMNERequest_Country(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Country() != "" {
		t.Error("Test failed null case should be null.")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "India",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.Country() != "India" {
		t.Error("Test failed nonNull case should be set to India.")
	}
}

func TestMNERequest_DirectionOfInvestment(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.DirectionOfInvestment() != "" {
		t.Error("Test failed null case should be null")
	}
	nonNullCase := MNERequest{
		directionOfInvestment: "reverse",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.DirectionOfInvestment() != "reverse" {
		t.Error("Test failed direction of investment should be reverse")
	}
}

func TestMNERequest_Industry(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.Industry() != "" {
		t.Error("Test failed industry should be set to null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "news",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.Industry() != "news" {
		t.Error("Test failed industry should be set to news")
	}
}

func TestMNERequest_SeriesId(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.SeriesId() != 0 {
		t.Error("Test failed seriesId should be set to 0")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              1,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.SeriesId() != 1 {
		t.Error("Test failed seriesId should be set to 1")
	}

}

func TestMNERequest_GetFootnotes(t *testing.T) {

	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.GetFootnotes() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "yes",
		requestType:           "",
	}

	if nonNullCase.GetFootnotes() != "yes" {
		t.Error("Test failed nonNull case should be yes")
	}
}

func TestMNERequest_RequestType(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.RequestType() != "" {
		t.Error("Test failed nullCase request type should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "JSON",
	}

	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNullCase should be JSON")
	}
}

func TestMNERequest_Year(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "2012",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.Year() != "2012" {
		t.Error("Test failed year should be set to 2012")
	}
}

func TestMNERequest_SetCountry(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Country() != "" {
		t.Error("Test failed nullCase should be set to null")
	}
	nullCase.SetCountry("USA")
	if nullCase.Country() != "USA" {
		t.Error("Test failed nullCAse should be set to USA")
	}
}

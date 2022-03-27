package Regional

import "testing"

func TestRegional_TableName(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	if nullCase.tableName != "" {
		t.Error("null case TableName")
	}

	nonNullCase := Regional{
		tableName: "Table",
		lineCode:  2,
		geoFips:   "392",
		year:      "2005",
	}
	if nonNullCase.tableName != "Table" {
		t.Error("nonNullCase TableName")
	}

}

func TestRegional_LineCode(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	if nullCase.lineCode != 0 {

		t.Error("nullCase LineCode")

	}
	nonNullCase := Regional{
		tableName: "name",
		lineCode:  2,
		geoFips:   "what is a fip",
		year:      "2003",
	}
	if nonNullCase.lineCode != 2 {
		t.Error("nonNullCase LineCode")
	}
}

func TestRegional_GeoFips(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	if nullCase.geoFips != "" {
		t.Error("nullCase Geofips")
	}
	nonNullCase := Regional{
		tableName: "table",
		lineCode:  3,
		geoFips:   "fils",
		year:      "3002",
	}
	if nonNullCase.geoFips != "fils" {
		t.Error("nonNullCase Geofips")
	}

}

func TestRegional_SetTableName(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	nullCase.SetTableName("newName")
	if nullCase.TableName() != "newName" {
		t.Error("nullCase SetTableName")
	}

	nullCase.SetTableName("otherName")
	if nullCase.TableName() != "otherName" {
		t.Error("nullCase2 SetTableName")
	}
}

func TestRegional_SetLineCode(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	nullCase.SetLineCode(20)
	if nullCase.LineCode() != 20 {
		t.Error("nullCase LineCode")
	}
	nullCase.SetLineCode(30)
	if nullCase.LineCode() != 30 {
		t.Error("nullCase2 LineCode")
	}
}

func TestRegional_SetGeoFips(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	nullCase.SetGeoFips("newFip")
	if nullCase.GeoFips() != "newFip" {
		t.Error("nullCase GeoFips")
	}

}

func TestRegional_SetYear(t *testing.T) {
	nullCase := Regional{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	nullCase.SetYear("2003")
	if nullCase.Year() != "2003" {
		t.Error("nullCase SetYear")
	}
	nullCase.SetYear("2004")
	if nullCase.Year() != "2004" {
		t.Error("nullCase2 SetYear")
	}
}

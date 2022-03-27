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

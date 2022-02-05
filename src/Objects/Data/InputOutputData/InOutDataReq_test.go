package InputOutputData

import (
	"testing"
)

func TestInputOutputDataReq_ReturnFormat(t *testing.T) {

	//check a null param case

	nullCase := InputOutputDataReq{
		tableId:      0,
		years:        nil,
		returnFormat: "",
	}

	if nullCase.ReturnFormat() != "" {
		t.Error("ReturnFormat() = false")
	}

	nonNullCase := InputOutputDataReq{
		tableId:      0,
		years:        nil,
		returnFormat: "JSON",
	}
	if nonNullCase.ReturnFormat() != "JSON" {
		t.Error("ReturnFormat() = false")
	}

}

func TestInputOutputDataReq_TableId(t *testing.T) {

	nullCase := InputOutputDataReq{
		tableId:      0,
		years:        nil,
		returnFormat: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("TableId() == 0")
	}

	nonNullCase := InputOutputDataReq{
		tableId:      52,
		years:        nil,
		returnFormat: "",
	}
	if nonNullCase.TableId() != 52 {
		t.Error("TableId() != 52")
	}
	//match only if offered

}

func TestInputOutputDataReq_Years(t *testing.T) {
	nullCase := InputOutputDataReq{
		tableId:      0,
		years:        nil,
		returnFormat: "",
	}

	//match correct year
	if nullCase.Years() != nil {
		t.Error("The Years != empty")
	}

	nonNullCase := InputOutputDataReq{
		tableId:      0,
		years:        []string{"2012", "2013"},
		returnFormat: "",
	}
	//match if year offered
	if nonNullCase.Years()[0] != "2012" {
		t.Error("The year[0] != 2012")
	}

	if nonNullCase.Years()[1] != "2013" {
		t.Error("The year[1] != 2013")
	}

}

func TestInputOutputDataReq_SetReturnFormat(t *testing.T) {

	//match correct type
	nullCase := InputOutputDataReq{
		tableId:      0,
		years:        nil,
		returnFormat: "JSON",
	}
	//match offered type
	nullCase.setReturnFormat("XML")
	if nullCase.ReturnFormat() != "XML" {
		t.Error("returnFormat != XML")
	}
	nullCase.setReturnFormat("JSON")
	if nullCase.ReturnFormat() != "JSON" {
		t.Error("returnFormat != JSON")
	}

}

func TestInputOutputDataReq_SetYears(t *testing.T) {

	//match if correct years

	//match if offered years

}

func TestInputOutputDataReq_AddYear(t *testing.T) {
	//match if correct

	//match if offered
}

func TestInputOutputDataReq_SetTableId(t *testing.T) {
	//match if correct

	//match if offered
}

func TestInputOutputDataReq_ToString(t *testing.T) {
	//match if correct

}

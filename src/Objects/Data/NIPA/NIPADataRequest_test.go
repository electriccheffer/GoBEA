package NIPA

import (
	"testing"
)

func TestNIPADataRequest_Frequency(t *testing.T) {

	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}

	if nullCase.Frequency() != nil {
		t.Error("frequency != nill")
	}

	nonNullCase := NIPADataRequest{
		tableId:   0,
		frequency: []string{"2003", "20"},
		years:     nil,
	}

	if nonNullCase.Frequency()[0] != "2003" {
		t.Error("nonNullCase.frequency[0] != 2003")
	}

	if nonNullCase.Frequency()[1] != "20" {
		t.Error("nonNullCase.frequency[1] != 20")
	}
}

func TestNIPADataRequest_TableId(t *testing.T) {

	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}
	if nullCase.TableId() != 0 {
		t.Error("tableId != 0")
	}

	nonNullCase := NIPADataRequest{
		tableId:   2,
		frequency: nil,
		years:     nil,
	}
	if nonNullCase.TableId() != 2 {
		t.Error("tableId != 2")
	}

}

func TestNIPADataRequest_Years(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}
	if nullCase.Years() != nil {
		t.Error("nullCase.years != nil")
	}
	nonNullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     []string{"2000", "2012"},
	}
	if nonNullCase.Years()[0] != "2000" {
		t.Error("nonNullCase.years[0] != 2000")
	}
	if nonNullCase.Years()[1] != "2012" {
		t.Error("nonNullCase.years[1] != 2012")
	}

}

func TestNIPADataRequest_setYears(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}
	nullCase.setYears([]string{"2000", "2012"})
	if nullCase.Years()[0] != "2000" {
		t.Error("nullCase.Years()[0] != 2000")
	}
	nullCase.setYears([]string{"1992", "1193"})
	if nullCase.Years()[0] != "1992" {
		t.Error("nullCase.Years()[0] != 1992")
	}
	if len(nullCase.Years()) != 2 {
		t.Error("size of the array is incorrect")
	}

}

func TestNIPADataRequest_setFrequency(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}
	nullCase.setFrequency([]string{"2012", "2000"})
	if nullCase.Frequency()[0] != "2012" {
		t.Error("nullCase.Frequency()[0] != 2012")
	}
	if nullCase.Frequency()[1] != "2000" {
		t.Error("nullCase.Frequency()[1] != 2000")
	}
	nullCase.setFrequency([]string{"1999", "1998"})
	if len(nullCase.Frequency()) != 2 {
		t.Error("len(nullCase.Frequency()) != 2")
	}
	if nullCase.Frequency()[0] != "1999" {
		t.Error("nullCase.Frequency()[0] != 1999")
	}
}

func TestNIPADataRequest_setTableId(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}
	nullCase.setTableId(12)
	if nullCase.TableId() != 12 {
		t.Error("nullCase.TableId() != 12")
	}

}

func TestNIPADataRequest_addFrequency(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}

	nullCase.addFrequency("2012")
	if nullCase.Frequency()[0] != "2012" {
		t.Error("nullCase.Frequency()[0] != 2012")
	}
	nullCase.addFrequency("2013")
	if nullCase.Frequency()[1] != "2013" {
		t.Error("nullCase.Frequency()[1] != 2013")
	}
	if len(nullCase.Frequency()) != 2 {
		t.Error("len(nullCase.Frequency()) != 2")
	}
}

func TestNIPADataRequest_addYear(t *testing.T) {
	nullCase := NIPADataRequest{
		tableId:   0,
		frequency: nil,
		years:     nil,
	}

	nullCase.addYear("210")
	if nullCase.Years()[0] != "210" {
		t.Error("nullCase.Years()[0] != 210")
	}
	nullCase.addYear("2000")
	if nullCase.Years()[1] != "2000" {
		t.Error("nullCase.Years()[1] != 2000")
	}
	if len(nullCase.Years()) != 2 {
		t.Error("len(nullCase.Years()) != 2")
	}

}

func TestNIPADataRequest_toString(t *testing.T) {
	nonNullCase := NIPADataRequest{
		tableId:   3,
		frequency: []string{"2003", "2004"},
		years:     []string{"122", "332"},
	}

	if nonNullCase.toString() != "3,2003 2004 ,122 332 " {
		t.Error("nonNullCase toString()")
	}
}

package GDPByIndustry

import "testing"

func TestGDPByIndustry_Frequency(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase frequency should be set to null")
	}
	nonNullCase := GDPByIndustry{
		tableId:   0,
		frequency: "regular",
		year:      0,
		industry:  "",
	}
	if nonNullCase.Frequency() != "regular" {
		t.Error("Test failed nonNullCase frequency should be set to regular")
	}
}

func TestGDPByIndustry_Year(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Year() != 0 {
		t.Error("Test failed null case year should be set to 0")
	}
	nonNullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      2012,
		industry:  "",
	}
	if nonNullCase.Year() != 2012 {
		t.Error("Test failed nonNullCase year should be set to 2012 ")
	}
}

func TestGDPByIndustry_Industry(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Industry() != "" {
		t.Error("Test failed nullCase industry should be set to null")
	}
	nonNullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "manufacturing",
	}
	if nonNullCase.Industry() != "manufacturing" {
		t.Error("Test failed nonNullCase industry should be set to manufacturing")
	}
}

func TestGDPByIndustry_TableId(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed nullCase should be set to 0")
	}
	nonNullCase := GDPByIndustry{
		tableId:   12,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nonNullCase.TableId() != 12 {
		t.Error("Test failed nonNullCase tableId should be set to 12")
	}
}

func TestGDPByIndustry_SetFrequency(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be set to null")
	}
	nullCase.SetFrequency("regular")
	if nullCase.Frequency() != "regular" {
		t.Error("Test Failed frequency should be set to regular")
	}
}

func TestGDPByIndustry_SetIndustry(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Industry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetIndustry("manufacturing")
	if nullCase.Industry() != "manufacturing" {
		t.Error("Test failed nullCase industry field should be set to manufacturing.")
	}
}

func TestGDPByIndustry_SetTableId(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed null case tableId should be set to 0")
	}
	nullCase.SetTableId(2)
	if nullCase.TableId() != 2 {
		t.Error("Test failed null case tableId field should be set to 2")
	}
}

func TestGDPByIndustry_SetYear(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:   0,
		frequency: "",
		year:      0,
		industry:  "",
	}
	if nullCase.Year() != 0 {
		t.Error("Test failed null case year field should be set to 0")
	}
	nullCase.SetYear(2012)
	if nullCase.Year() != 2012 {
		t.Error("Test failed null case year field should be set to 2012")
	}
}

func TestGDPByIndustry_String(t *testing.T) {
	stringCase := GDPByIndustry{
		tableId:   1,
		frequency: "regular",
		year:      2012,
		industry:  "manufacturing",
	}
	if stringCase.String() != "1 regular 2012 manufacturing" {
		t.Error("Test failed string should read 1 regular 2012 manufacturing")
	}
}

func TestNewGDPByIndustry(t *testing.T) {
	newCase := NewGDPByIndustry(1, "regular", 2012, "manufacturing")
	if newCase.TableId() != 1 {
		t.Error("Test failed tableId should be 1")
	}
	if newCase.Frequency() != "regular" {
		t.Error("Test failed frequency should be set to regular")
	}
	if newCase.Year() != 2012 {
		t.Error("Test failed year should be set to 2012")
	}
	if newCase.Industry() != "manufacturing" {
		t.Error("Test failed industry should be set to manufacturing")
	}
}

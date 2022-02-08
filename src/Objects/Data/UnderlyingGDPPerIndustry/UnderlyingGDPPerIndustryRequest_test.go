package UnderlyingGDPPerIndustry

import "testing"

func TestUnderLyingGDPPerIndustry_TableId(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}
	if nullCase.TableId() != 0 {
		t.Error("nullCase.TableId() != 0")
	}
}

func TestUnderLyingGDPPerIndustryRequest_Frequency(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}
	if nullCase.Frequency() != "" {
		t.Error("nullCase.Frequency() != null")
	}

}

func TestUnderLyingGDPPerIndustryRequest_Years(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}
	if nullCase.Years() != nil {
		t.Error("")
	}

}

func TestUnderLyingGDPPerIndustryRequest_Industry(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}

	if nullCase.Industry() != "" {
		t.Error("nullCase.Industry() != null ")
	}

}

func TestUnderLyingGDPPerIndustryRequest_setTableId(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}

	nullCase.setTableId(12)
	if nullCase.TableId() != 12 {
		t.Error("nullCase.TableId() != 12")
	}

	nullCase.setTableId(13)
	if nullCase.TableId() != 13 {
		t.Error("nullCase.TableId() != 13")
	}
}

func TestUnderLyingGDPPerIndustryRequest_setFrequency(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}

	nullCase.setFrequency("another")
	if nullCase.Frequency() != "another" {
		t.Error("nullCase.Frequency() != \"another\"")
	}
}

func TestUnderLyingGDPPerIndustryRequest_setYears(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}

	nullCase.setYears([]string{"2004", "2003", "2007"})
	if nullCase.Years()[0] != "2004" {
		t.Error("nullCase.Years()[0] != \"2004\"")
	}

}

func TestUnderLyingGDPPerIndustryRequest_setIndustry(t *testing.T) {
	nullCase := UnderLyingGDPPerIndustryRequest{
		tableId:   0,
		frequency: "",
		years:     nil,
		industry:  "",
	}

	nullCase.setIndustry("steel")
	if nullCase.Industry() != "steel" {
		t.Error("nullCase.Industry() != \"steel\"")
	}

	nullCase.setIndustry("tech")
	if nullCase.Industry() != "tech" {
		t.Error("nullCase.Industry() != \"tech\"")
	}

}

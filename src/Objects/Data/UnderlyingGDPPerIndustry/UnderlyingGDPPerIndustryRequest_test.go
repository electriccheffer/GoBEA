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

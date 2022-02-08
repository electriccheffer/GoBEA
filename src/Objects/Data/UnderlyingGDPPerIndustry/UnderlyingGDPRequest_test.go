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

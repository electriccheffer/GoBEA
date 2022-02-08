package UnderlyingGDPPerIndustry

type UnderLyingGDPPerIndustryRequest struct {
	tableId   int
	frequency string
	years     []string
	industry  string
}

//TableId returns the set tableId value for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) TableId() int {
	return r.tableId
}

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

//setTableId sets the value of the tableId field of the UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) setTableId(newId int) {
	r.tableId = newId
}

//Frequency returns the frequency value for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Frequency() string {
	return r.frequency
}

//Years returns an array of strings for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Years() []string {
	return r.years
}

//Industry returns the value of the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Industry() string {
	return r.industry
}
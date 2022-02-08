package UnderlyingGDPPerIndustry

import "strconv"

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

//setFrequency sets the frequency field of the UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) setFrequency(newFrequency string) {
	r.frequency = newFrequency
}

//Years returns an array of strings for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Years() []string {
	return r.years
}

//setYears sets the array of the years field of a UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) setYears(newYears []string) {
	r.years = newYears
}

//addYear appends a year to the years field of UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) addYear(newYear string) {
	r.years = append(r.years, newYear)
}

//Industry returns the value of the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Industry() string {
	return r.industry
}

//setIndustry sets the industry field of the UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) setIndustry(newIndustry string) {
	r.industry = newIndustry
}

//toString() a function which returns the current string representation of an UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) toString() string {

	var request = ""
	request += r.Frequency() + " "
	request += r.Industry() + " "
	request += strconv.Itoa(r.TableId())
	for i := 0; i < len(r.Years()); i++ {
		request += " " + r.Years()[i]
	}
	return request
}

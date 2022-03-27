package NIPA

import "strconv"

//NIPADataRequest is a struct used to represent a NIPA request
//to the BEA.  It consists of a tableId, frequency[], years[].
type NIPADataRequest struct {
	tableId   int
	frequency []string
	years     []string
}

//TableId returns the value of the tableId field of the NIPADataRequest.
func (r NIPADataRequest) TableId() int {
	return r.tableId
}

//setTableId sets the value of the tableId field for a NIPADataRequest object.
func (r *NIPADataRequest) setTableId(newTableId int) {
	r.tableId = newTableId
}

//Frequency returns an array of strings for the time series data.
func (r NIPADataRequest) Frequency() []string {
	return r.frequency
}

//setFrequency sets the value of the frequency field fo a NIPADataRequest object.
func (r *NIPADataRequest) setFrequency(newFrequencies []string) {
	r.frequency = newFrequencies
}

//addFrequency adds a new frequency to the end of the list of frequencies in the NIPADataRequest
func (r *NIPADataRequest) addFrequency(newFrequency string) {
	r.frequency = append(r.frequency, newFrequency)
}

//Years returns the array of years in the NIPADataRequest.
func (r NIPADataRequest) Years() []string {
	return r.years
}

//setYears sets the value of years field of a NIPADataRequest object.
func (r *NIPADataRequest) setYears(newYears []string) {
	r.years = newYears
}

func (r *NIPADataRequest) addYear(newYear string) {
	r.years = append(r.years, newYear)
}

func (r *NIPADataRequest) toString() string {

	var returnString string = strconv.Itoa(r.tableId) + ","

	for _, freq := range r.frequency {

		returnString += freq + " "

	}
	returnString += ","

	for _, year := range r.years {
		returnString += year + " "
	}
	return returnString
}

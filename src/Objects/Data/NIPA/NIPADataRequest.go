package NIPA

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
func (r NIPADataRequest) setTableId(newTableId int) {
	r.tableId = newTableId
}

//Frequency returns an array of strings for the time series data.
func (r NIPADataRequest) Frequency() []string {
	return r.frequency
}

//setFrequency sets the value of the frequency field fo a NIPADataRequest object.
func (r NIPADataRequest) setFrequency(newFrequencies []string) {
	r.frequency = newFrequencies
}

//Years returns the array of years in the NIPADataRequest.
func (r NIPADataRequest) Years() []string {
	return r.years
}

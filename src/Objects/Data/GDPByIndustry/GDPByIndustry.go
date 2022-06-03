package GDPByIndustry

import "strconv"

type GDPByIndustry struct {
	tableId   int
	frequency string
	year      int
	industry  string
}

//TableId is a getter for the tableId field of a GDPByIndustry object
func (G *GDPByIndustry) TableId() int {
	return G.tableId
}

//SetTableId is a setter for the tableId field of a GDPByIndustry object
func (G *GDPByIndustry) SetTableId(tableId int) {
	G.tableId = tableId
}

//Frequency is a getter for the frequency field of a GDPByIndustry object
func (G *GDPByIndustry) Frequency() string {
	return G.frequency
}

//SetFrequency is a setter for the frequency field of a GDPByIndustry object
func (G *GDPByIndustry) SetFrequency(frequency string) {
	G.frequency = frequency
}

//Year is a getter for the the year field of a GDPByIndustry object
func (G *GDPByIndustry) Year() int {
	return G.year
}

//SetYear is a setter for a GDPByIndustry object
func (G *GDPByIndustry) SetYear(year int) {
	G.year = year
}

//Industry is a getter for a GDPByIndustry object
func (G *GDPByIndustry) Industry() string {
	return G.industry
}

//SetIndustry is a setter for the industry field of the GDPByIndustry object
func (G *GDPByIndustry) SetIndustry(industry string) {
	G.industry = industry
}

//NewGDPByIndustry is a constructor for a GDPByIndustry object
func NewGDPByIndustry(tableId int, frequency string, year int, industry string) *GDPByIndustry {
	return &GDPByIndustry{tableId: tableId, frequency: frequency, year: year, industry: industry}
}

//String is a method which returns a string representation of a GDPByIndustry object
func (G *GDPByIndustry) String() string {
	return strconv.Itoa(G.tableId) + " " + G.frequency + " " + strconv.Itoa(G.year) + " " + G.industry
}

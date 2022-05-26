package MNE

import "strconv"

type MNERequest struct {
	directionOfInvestment string
	seriesId              int
	classification        string
	year                  string
	country               string
	industry              string
	getFootnotes          string
	requestType           string
}

//NewMNERequest is a constructor for the MNERequest
func NewMNERequest(directionOfInvestment string, seriesId int, classification string, year string, country string, industry string, getFootnotes string, requestType string) *MNERequest {
	return &MNERequest{directionOfInvestment: directionOfInvestment, seriesId: seriesId, classification: classification, year: year, country: country, industry: industry, getFootnotes: getFootnotes, requestType: requestType}
}

//DirectionOfInvestment is a getter for directionOfInvestment MNERequest
func (M *MNERequest) DirectionOfInvestment() string {
	return M.directionOfInvestment
}

//SetDirectionOfInvestment is a setter for directionOfInvestment MNERequest
func (M *MNERequest) SetDirectionOfInvestment(directionOfInvestment string) {
	M.directionOfInvestment = directionOfInvestment
}

//SeriesId is a getter for seriesId MNERequest
func (M *MNERequest) SeriesId() int {
	return M.seriesId
}

//SetSeriesId is a setter for seriesId MNERequest
func (M *MNERequest) SetSeriesId(seriesId int) {
	M.seriesId = seriesId
}

//Classification is a getter for classification MNERequest
func (M *MNERequest) Classification() string {
	return M.classification
}

//SetClassification is a setter for classification MNERequest
func (M *MNERequest) SetClassification(classification string) {
	M.classification = classification
}

//Year is a getter for year MNERequest
func (M *MNERequest) Year() string {
	return M.year
}

//SetYear is a setter for year MNERequest
func (M *MNERequest) SetYear(year string) {
	M.year = year
}

//Country is a getter for country MNERequest
func (M *MNERequest) Country() string {
	return M.country
}

//SetCountry is a setter for country MNERequest
func (M *MNERequest) SetCountry(country string) {
	M.country = country
}

//Industry is a getter for industry MNERequest
func (M *MNERequest) Industry() string {
	return M.industry
}

//SetIndustry is a setter for industry MNERequest
func (M *MNERequest) SetIndustry(industry string) {
	M.industry = industry
}

//GetFootnotes is a getter getFootnotes MNERequest
func (M *MNERequest) GetFootnotes() string {
	return M.getFootnotes
}

//SetGetFootnotes is a setter for getFootnotes MNERequest
func (M *MNERequest) SetGetFootnotes(getFootnotes string) {
	M.getFootnotes = getFootnotes
}

//RequestType is a getter for requestType MNERequest
func (M *MNERequest) RequestType() string {
	return M.requestType
}

//SetRequestType is a setter for requestType MNERequest
func (M *MNERequest) SetRequestType(requestType string) {
	M.requestType = requestType
}

//String is a method which returns a string of the MNERequest
func (M *MNERequest) String() string {
	return M.directionOfInvestment + " " + strconv.Itoa(M.seriesId) + " " + M.classification + " " + M.year + " " + M.country + " " + M.industry + " " + M.getFootnotes + " " + M.requestType + " " + M.requestType
}

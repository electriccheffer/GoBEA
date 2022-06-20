package src

import (
	"BEALIb/src/Objects/Data/RequestInterface"
	"strconv"
)

type Request interface {
	DataRequest() string
}

type FixedAssetsRequest struct {
	tableName   string
	year        string
	requestType string
}

//TableName is a getter for the tableName FixedAssetsRequest
func (f *FixedAssetsRequest) TableName() string {
	return f.tableName
}

//SetTableName is a setter for the tableName FixedAssetsRequest
func (f *FixedAssetsRequest) SetTableName(tableName string) {
	f.tableName = tableName
}

//Year is a getter for the year FixedAssetsRequest
func (f *FixedAssetsRequest) Year() string {
	return f.year
}

//SetYear is a setter for the year of a FixedAssetsRequest
func (f *FixedAssetsRequest) SetYear(year string) {
	f.year = year
}

//RequestType is a getter for the requestType of a FixedAssetsRequest
func (f *FixedAssetsRequest) RequestType() string {
	return f.requestType
}

//SetRequestType is a setter for the requestType of a FixedAssetsRequest
func (f *FixedAssetsRequest) SetRequestType(requestType string) {
	f.requestType = requestType
}

//NewFixedAssetsRequest is a constructor for a FixedAssetsRequest
func NewFixedAssetsRequest(tableName string, year string, requestType string) *FixedAssetsRequest {
	return &FixedAssetsRequest{tableName: tableName, year: year, requestType: requestType}
}

//String is a method which returns a string of the FixedAssetsRequest
func (f *FixedAssetsRequest) String() string {
	return f.tableName + " " + f.year + " " + f.requestType
}

type ITA struct {
	indicator     string
	areaOrCountry string
	frequency     string
	year          string
	requestType   string
}

//Indicator is a getter for the indicator field of an ITA struct
func (I *ITA) Indicator() string {
	return I.indicator
}

//SetIndicator is a setter for the indicator field of an ITA struct
func (I *ITA) SetIndicator(indicator string) {
	I.indicator = indicator
}

//AreaOrCountry is a getter for a the areaOrCountry field of an ITA struct
func (I *ITA) AreaOrCountry() string {
	return I.areaOrCountry
}

//SetAreaOrCountry is a setter for the areaOrCountry field of an ITA struct
func (I *ITA) SetAreaOrCountry(areaOrCountry string) {
	I.areaOrCountry = areaOrCountry
}

//Frequency is a getter for the frequency field of an ITA struct
func (I *ITA) Frequency() string {
	return I.frequency
}

//SetFrequency is a setter for the frequency field of an ITA struct
func (I *ITA) SetFrequency(frequency string) {
	I.frequency = frequency
}

//Year is a getter for the year field of an ITA struct
func (I *ITA) Year() string {
	return I.year
}

//SetYear is a setter for the year field of an ITA struct
func (I *ITA) SetYear(year string) {
	I.year = year
}

//RequestType is a getter for the requestType field of an ITA struct
func (I *ITA) RequestType() string {
	return I.requestType
}

//SetRequestType is a setter for the requestType field of an ITA struct
func (I *ITA) SetRequestType(requestType string) {
	I.requestType = requestType
}

//NewITA creates a new instance of an ITA struct
func NewITA(indicator string, areaOrCountry string, frequency string, year string, requestType string) *ITA {
	return &ITA{indicator: indicator, areaOrCountry: areaOrCountry, frequency: frequency, year: year, requestType: requestType}
}

//String is a string representation of ITA struct
func (I *ITA) String() string {
	return I.indicator + " " + I.areaOrCountry + " " + I.frequency + " " + I.year + " " + I.requestType
}

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

//InputOutputDataReq is a type struct which embodies all the required fields
//	of a request made for Input/Output data
type InputOutputDataReq struct {
	tableId      int
	years        []string
	returnFormat string
	RequestInterface.Request
}

//ReturnFormat provides a method for returning the value of the returnFormat
//	variable of the InputOutputDataReq struct
func (r InputOutputDataReq) ReturnFormat() string {
	return r.returnFormat
}

//setReturnFormat provides a method for setting the returnFormat variable of
//	of the InputOutputDataReq
func (r *InputOutputDataReq) setReturnFormat(format string) {
	r.returnFormat = format
}

//Years provides a method for returning the years[] variable of the InputOutputDataReq
//struct
func (r InputOutputDataReq) Years() []string {
	return r.years
}

//addYear provides a method for appending a year to the years[] variable
func (r *InputOutputDataReq) addYear(year string) {
	if r.years == nil {
		r.years = []string{year}
		return
	}
	r.years = append(r.years, year)
}

//setYears provides a method for changing the entire years[] with a new array of years[]
//for the InputOutputDataReq struct
func (r *InputOutputDataReq) setYears(newYears []string) {
	r.years = newYears
}

//TableId provides a method for getting the tableId of the InputOutputDataReq
//struct
func (r InputOutputDataReq) TableId() int {
	return r.tableId
}

//setTableId provides a method for setting the tableId field of the InputOutputDataReq
//struct
func (r *InputOutputDataReq) setTableId(id int) {

	r.tableId = id

}

//toString provides a method for returning the InputOutputDataReq struct as a
//string in the format of a BEA request
func (r *InputOutputDataReq) toString() (string, int) {

	var returnYears string
	for _, year := range r.years {

		returnYears += " " + year

	}
	return strconv.Itoa(r.tableId) + returnYears + " " + r.returnFormat, 0

}

type IIPRequest struct {
	typeOfInvestment string
	component        string
	frequency        string
	year             string
	returnFormat     string
}

func (I *IIPRequest) ReturnFormat() string {
	return I.returnFormat
}

func (I *IIPRequest) SetReturnFormat(returnFormat string) {
	I.returnFormat = returnFormat
}

func (I *IIPRequest) TypeOfInvestment() string {
	return I.typeOfInvestment
}

func (I *IIPRequest) SetTypeOfInvestment(typeOfInvestment string) {
	I.typeOfInvestment = typeOfInvestment
}

func (I *IIPRequest) Component() string {
	return I.component
}

func (I *IIPRequest) SetComponent(component string) {
	I.component = component
}

func (I *IIPRequest) Frequency() string {
	return I.frequency
}

func (I *IIPRequest) SetFrequency(frequency string) {
	I.frequency = frequency
}

func (I *IIPRequest) Year() string {
	return I.year
}

func (I *IIPRequest) SetYear(year string) {
	I.year = year
}

func (I *IIPRequest) String() string {
	return I.component + " " + I.typeOfInvestment + " " + I.frequency + " " + I.year + " " + I.returnFormat
}

type IntlServTradeRequest struct {
	typeOfService  string
	tradeDirection string
	affiliation    string
	areaOrCountry  string
}

//TypeOfService is a method used to access the typeOfService field of a IntlServTradeRequest object
func (i *IntlServTradeRequest) TypeOfService() string {
	return i.typeOfService
}

//SetTypeOfService is a method used to set the typeOfService field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetTypeOfService(typeOfService string) {
	i.typeOfService = typeOfService
}

//TradeDirection is a method used to access the tradeDirection field of IntlServTradeRequest object
func (i *IntlServTradeRequest) TradeDirection() string {
	return i.tradeDirection
}

//SetTradeDirection is a method used to set the tradeDirection field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetTradeDirection(tradeDirection string) {
	i.tradeDirection = tradeDirection
}

//Affiliation is a method used to access the affiliation field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) Affiliation() string {
	return i.affiliation
}

//SetAffiliation is a method used to set the affiliation field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetAffiliation(affiliation string) {
	i.affiliation = affiliation
}

//AreaOrCountry is a method used to access the areaOrCountry field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) AreaOrCountry() string {
	return i.areaOrCountry
}

//SetAreaOrCountry is a method used to set the areaOrCountry field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetAreaOrCountry(areaOrCountry string) {
	i.areaOrCountry = areaOrCountry
}

//toString returns a string representation of the IntlServTradeRequest object
func (i *IntlServTradeRequest) toString() string {
	return i.TypeOfService() + " " + i.TradeDirection() + " " + i.Affiliation() + " " + i.AreaOrCountry()
}

type IntlServTradeRequest struct {
	typeOfService  string
	tradeDirection string
	affiliation    string
	areaOrCountry  string
}

//TypeOfService is a method used to access the typeOfService field of a IntlServTradeRequest object
func (i *IntlServTradeRequest) TypeOfService() string {
	return i.typeOfService
}

//SetTypeOfService is a method used to set the typeOfService field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetTypeOfService(typeOfService string) {
	i.typeOfService = typeOfService
}

//TradeDirection is a method used to access the tradeDirection field of IntlServTradeRequest object
func (i *IntlServTradeRequest) TradeDirection() string {
	return i.tradeDirection
}

//SetTradeDirection is a method used to set the tradeDirection field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetTradeDirection(tradeDirection string) {
	i.tradeDirection = tradeDirection
}

//Affiliation is a method used to access the affiliation field of the IntlServTradeRequest object
func (i *IntlServTradeRequest) Affiliation() string {
	return i.affiliation
}

//SetAffiliation is a method used to set the affiliation field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetAffiliation(affiliation string) {
	i.affiliation = affiliation
}

//AreaOrCountry is a method used to access the areaOrCountry field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) AreaOrCountry() string {
	return i.areaOrCountry
}

//SetAreaOrCountry is a method used to set the areaOrCountry field in the IntlServTradeRequest object
func (i *IntlServTradeRequest) SetAreaOrCountry(areaOrCountry string) {
	i.areaOrCountry = areaOrCountry
}

//toString returns a string representation of the IntlServTradeRequest object
func (i *IntlServTradeRequest) toString() string {
	return i.TypeOfService() + " " + i.TradeDirection() + " " + i.Affiliation() + " " + i.AreaOrCountry()
}

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
	return M.directionOfInvestment + " " + strconv.Itoa(M.seriesId) + " " + M.classification + " " + M.year + " " + M.country + " " + M.industry + " " + M.getFootnotes + " " + M.requestType
}

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

type NIUnderlyingDetailRequest struct {
	tableName   string
	frequency   string
	year        string
	requestType string
}

//NewNIUnderlyingDetailRequest a constuctor for anNIUnderlyingDetailRequest
func NewNIUnderlyingDetailRequest(tableName string, frequency string, year string, requestType string) *NIUnderlyingDetailRequest {
	return &NIUnderlyingDetailRequest{tableName: tableName, frequency: frequency, year: year, requestType: requestType}
}

//RequestType is a setter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) RequestType() string {
	return N.requestType
}

//SetRequestType is a getter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) SetRequestType(requestType string) {
	N.requestType = requestType
}

//TableName is a getter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) TableName() string {
	return N.tableName
}

//SetTableName is a setter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) SetTableName(tableName string) {
	N.tableName = tableName
}

//Frequency is a getter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) Frequency() string {
	return N.frequency
}

//SetFrequency is a setter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) SetFrequency(frequency string) {
	N.frequency = frequency
}

//Year is a getter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) Year() string {
	return N.year
}

//SetYear is a setter for an NIUnderlyingDetailRequest
func (N *NIUnderlyingDetailRequest) SetYear(year string) {
	N.year = year
}

//String is a method to return a string version of an NewNIUnderlyingDetailRequest
func (N NIUnderlyingDetailRequest) String() string {
	return N.tableName + " " + N.frequency + " " + N.year
}

type Regional struct {
	tableName string
	lineCode  int
	geoFips   string
	year      string
}

func (r *Regional) TableName() string {
	return r.tableName
}

func (r *Regional) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *Regional) LineCode() int {
	return r.lineCode
}

func (r *Regional) SetLineCode(lineCode int) {
	r.lineCode = lineCode
}

func (r *Regional) GeoFips() string {
	return r.geoFips
}

func (r *Regional) SetGeoFips(geoFips string) {
	r.geoFips = geoFips
}

func (r *Regional) Year() string {
	return r.year
}

func (r *Regional) SetYear(year string) {
	r.year = year
}

func NewRegional(tableName string, lineCode int, geoFips string, year string) *Regional {
	return &Regional{tableName: tableName, lineCode: lineCode, geoFips: geoFips, year: year}
}

func (r *Regional) toString() string {

	return r.tableName + ", " + strconv.Itoa(r.lineCode) + ", " + r.geoFips + ", " + r.year
}

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

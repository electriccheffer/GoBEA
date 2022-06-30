package GoBea

import (
	"strconv"
)

const apiAddress = "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName="

type Request interface {
	Request() string
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

func (f *FixedAssetsRequest) Request() string {
	return apiAddress + "FixedAssets&" + "TableName=" + f.tableName + "&Year=" + f.year + "&ResultFormat=" + f.requestType
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

func (I *ITA) Request() string {
	return apiAddress + "ITA&Indicator=" + I.indicator + "&AreaOrCountry=" + I.areaOrCountry + "&Frequency=" + I.frequency + "&Year=" + I.year + "&ResultFormat=" + I.requestType
}

type GDPByIndustry struct {
	tableId     int
	frequency   string
	year        string
	industry    string
	requestType string
}

func NewGDPByIndustry(tableId int, frequency string, year string, industry string, requestType string) *GDPByIndustry {
	return &GDPByIndustry{tableId: tableId, frequency: frequency, year: year, industry: industry, requestType: requestType}
}

func (G *GDPByIndustry) RequestType() string {
	return G.requestType
}

func (G *GDPByIndustry) SetRequestType(requestType string) {
	G.requestType = requestType
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
func (G *GDPByIndustry) Year() string {
	return G.year
}

//SetYear is a setter for a GDPByIndustry object
func (G *GDPByIndustry) SetYear(year string) {
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

//String is a method which returns a string representation of a GDPByIndustry object
func (G *GDPByIndustry) String() string {
	return strconv.Itoa(G.tableId) + " " + G.frequency + " " + G.year + " " + G.industry + " " + G.requestType
}

func (G *GDPByIndustry) Request() string {
	return apiAddress + "&DataSetName=GDPbyIndustry&Year=" + G.year + "&Industry=" + G.industry + "&tableID=" + strconv.Itoa(G.tableId) + "&Frequency=" + G.frequency + "&ResultFormat=" + G.requestType
}

//InputOutputDataReq is a type struct which embodies all the required fields
//	of a request made for Input/Output data
type InputOutputDataReq struct {
	tableId      int
	years        string
	returnFormat string
}

func NewInputOutputDataReq(tableId int, years string, returnFormat string) *InputOutputDataReq {
	return &InputOutputDataReq{tableId: tableId, years: years, returnFormat: returnFormat}
}

//ReturnFormat provides a method for returning the value of the returnFormat
//	variable of the InputOutputDataReq struct
func (r InputOutputDataReq) ReturnFormat() string {
	return r.returnFormat
}

//SetReturnFormat provides a method for setting the returnFormat variable of
//	of the InputOutputDataReq
func (r *InputOutputDataReq) SetReturnFormat(format string) {
	r.returnFormat = format
}

//Years provides a method for returning the years[] variable of the InputOutputDataReq
//struct
func (r InputOutputDataReq) Years() string {
	return r.years
}

//AddYear provides a method for appending a year to the years[] variable
func (r *InputOutputDataReq) AddYear(year string) {
	if r.years != "" {
		r.years += "," + year
		return
	}
	r.SetYears(year)
}

//SetYears provides a method for changing the entire years[] with a new array of years[]
//for the InputOutputDataReq struct
func (r *InputOutputDataReq) SetYears(newYears string) {
	r.years = newYears
}

//TableId provides a method for getting the tableId of the InputOutputDataReq
//struct
func (r InputOutputDataReq) TableId() int {
	return r.tableId
}

//SetTableId provides a method for setting the tableId field of the InputOutputDataReq
//struct
func (r *InputOutputDataReq) SetTableId(id int) {

	r.tableId = id

}

//toString provides a method for returning the InputOutputDataReq struct as a
//string in the format of a BEA request
func (r *InputOutputDataReq) String() string {

	return strconv.Itoa(r.tableId) + " " + r.years + " " + r.returnFormat

}

func (r *InputOutputDataReq) Request() string {

	return apiAddress + "InputOutput" + "&Year=" + r.years + "&tableID=" + strconv.Itoa(r.tableId) + "&ResultFormat=" + r.returnFormat
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

func (I *IIPRequest) Request() string {
	return apiAddress + "IIP" + "&TypeOfInvestment=" + I.typeOfInvestment + "&Component=" + I.component + "&Frequency=" + I.frequency + "&Year=" + I.year + "&RequestFormat=" + I.returnFormat
}

type IntlServTradeRequest struct {
	typeOfService  string
	tradeDirection string
	affiliation    string
	areaOrCountry  string
	requestType    string
}

func NewIntlServTradeRequest(typeOfService string, tradeDirection string, affiliation string, areaOrCountry string, requestType string) *IntlServTradeRequest {
	return &IntlServTradeRequest{typeOfService: typeOfService, tradeDirection: tradeDirection, affiliation: affiliation, areaOrCountry: areaOrCountry, requestType: requestType}
}

func (i *IntlServTradeRequest) RequestType() string {
	return i.requestType
}

func (i *IntlServTradeRequest) SetRequestType(requestType string) {
	i.requestType = requestType
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
func (i *IntlServTradeRequest) String() string {
	return i.TypeOfService() + " " + i.TradeDirection() + " " + i.Affiliation() + " " + i.AreaOrCountry() + " " + i.RequestType()
}

func (i *IntlServTradeRequest) Request() string {
	return apiAddress + "ItlServTrade" + "&TypeOfService=" + i.typeOfService + "&TradeDirection=" + i.tradeDirection + "&Affiliation=" + i.affiliation + "&AreaOrCountry=" + i.areaOrCountry + "&ResultFormat=" + i.requestType
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

func (M *MNERequest) Request() string {
	return apiAddress + "MNE" + "&Year=" + M.directionOfInvestment + "&Country=" + M.country + "&DirectonOfInvestment=" + M.directionOfInvestment + "&Classification=" + M.classification + "&ResultFormat=" + M.requestType
}

//NIPADataRequest is a struct used to represent a NIPA request
//to the BEA.  It consists of a tableId, frequency[], years[].
type NIPADataRequest struct {
	tableId     int
	frequency   string
	years       string
	requestType string
}

func NewNIPADataRequest(tableId int, frequency string, years string, requestType string) *NIPADataRequest {
	return &NIPADataRequest{tableId: tableId, frequency: frequency, years: years, requestType: requestType}
}

func (r *NIPADataRequest) RequestType() string {
	return r.requestType
}

func (r *NIPADataRequest) SetRequestType(requestType string) {
	r.requestType = requestType
}

//TableId returns the value of the tableId field of the NIPADataRequest.
func (r NIPADataRequest) TableId() int {
	return r.tableId
}

//SetTableId sets the value of the tableId field for a NIPADataRequest object.
func (r *NIPADataRequest) SetTableId(newTableId int) {
	r.tableId = newTableId
}

//Frequency returns an array of strings for the time series data.
func (r NIPADataRequest) Frequency() string {
	return r.frequency
}

//SetFrequency sets the value of the frequency field fo a NIPADataRequest object.
func (r *NIPADataRequest) SetFrequency(newFrequencies string) {
	r.frequency = newFrequencies
}

//AddFrequency adds a new frequency to the end of the list of frequencies in the NIPADataRequest
func (r *NIPADataRequest) AddFrequency(newFrequency string) {
	if r.frequency == "" {
		r.frequency = newFrequency
		return
	}
	r.frequency += "," + newFrequency
}

//Years returns the array of years in the NIPADataRequest.
func (r NIPADataRequest) Years() string {
	return r.years
}

//SetYears sets the value of years field of a NIPADataRequest object.
func (r *NIPADataRequest) SetYears(newYears string) {
	r.years = newYears
}

//AddYear appends a new year to the NIPADataRequest
func (r *NIPADataRequest) AddYear(newYear string) {
	if r.years == "" {
		r.years = newYear
		return
	}
	r.years += "," + newYear
}

//String creates a string form of the NIPADataRequest
func (r *NIPADataRequest) String() string {

	var returnString string = strconv.Itoa(r.tableId) + " "
	returnString += r.frequency + " "
	returnString += r.years + " "
	returnString += r.requestType
	return returnString
}

func (r *NIPADataRequest) Request() string {
	return apiAddress + "NIPA" + "&TableName=" + strconv.Itoa(r.tableId) + "&Frequency=" + r.frequency + "&Year=" + r.years + "&RequestType=" + r.requestType
}

type NIUnderlyingDetailRequest struct {
	tableName   string
	frequency   string
	year        string
	requestType string
}

//NewNIUnderlyingDetailRequest a constructor for anNIUnderlyingDetailRequest
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

func (N *NIUnderlyingDetailRequest) Request() string {
	return apiAddress + "NIUnderlyingDetail" + "&TableName=" + N.tableName + "&Frequency=" + N.frequency + "&Year=" + N.year + "&RequestType=" + N.requestType
}

type RegionalRequest struct {
	tableName     string
	lineCode      int
	geoFips       string
	year          string
	requestFormat string
}

//RequestFormat is a getter for the return format of the RegionalRequest
func (r *RegionalRequest) RequestFormat() string {
	return r.requestFormat
}

//SetRequestFormat is a setter for the return format of the RegionalRequest
func (r *RegionalRequest) SetRequestFormat(requestFormat string) {
	r.requestFormat = requestFormat
}

//TableName is a getter for the table name of the RegionalRequest
func (r *RegionalRequest) TableName() string {
	return r.tableName
}

//SetTableName is a setter for the table name of the RegionalRequest
func (r *RegionalRequest) SetTableName(tableName string) {
	r.tableName = tableName
}

//LineCode is a getter for the line code of the RegionalRequest
func (r *RegionalRequest) LineCode() int {
	return r.lineCode
}

//SetLineCode is a setter for the line code of the RegionalRequest
func (r *RegionalRequest) SetLineCode(lineCode int) {
	r.lineCode = lineCode
}

//GeoFips is a getter for the geo fips of the RegionalRequest
func (r *RegionalRequest) GeoFips() string {
	return r.geoFips
}

//SetGeoFips is a setter for the geo fips of the RegionalRequest
func (r *RegionalRequest) SetGeoFips(geoFips string) {
	r.geoFips = geoFips
}

//Year is a getter for the year of the RegionalRequest
func (r *RegionalRequest) Year() string {
	return r.year
}

//SetYear is a setter for the year of the RegionalRequest
func (r *RegionalRequest) SetYear(year string) {
	r.year = year
}

//NewRegionalRequest is a constructor for a RegionalRequest
func NewRegionalRequest(tableName string, lineCode int, geoFips string, year string) *RegionalRequest {
	return &RegionalRequest{tableName: tableName, lineCode: lineCode, geoFips: geoFips, year: year}
}

//String returns a string representation of a RegionalRequest
func (r *RegionalRequest) String() string {

	return r.tableName + ", " + strconv.Itoa(r.lineCode) + ", " + r.geoFips + ", " + r.year
}

func (r *RegionalRequest) Request() string {
	return apiAddress + "Regional" + "&TableName=" + r.tableName + "&LineCode=" + strconv.Itoa(r.lineCode) + "&Year" + r.year + "&GeoFips=" + r.geoFips + "&RequestFormat=" + r.requestFormat
}

type UnderLyingGDPPerIndustryRequest struct {
	tableId      int
	frequency    string
	years        string
	industry     string
	resultFormat string
}

func NewUnderLyingGDPPerIndustryRequest(tableId int, frequency string, years string, industry string, resultFormat string) *UnderLyingGDPPerIndustryRequest {
	return &UnderLyingGDPPerIndustryRequest{tableId: tableId, frequency: frequency, years: years, industry: industry, resultFormat: resultFormat}
}

func (r *UnderLyingGDPPerIndustryRequest) SetTableId(tableId int) {
	r.tableId = tableId
}

func (r *UnderLyingGDPPerIndustryRequest) SetFrequency(frequency string) {
	r.frequency = frequency
}

func (r *UnderLyingGDPPerIndustryRequest) SetYears(years string) {
	r.years = years
}

func (r *UnderLyingGDPPerIndustryRequest) ResultFormat() string {
	return r.resultFormat
}

func (r *UnderLyingGDPPerIndustryRequest) SetResultFormat(resultFormat string) {
	r.resultFormat = resultFormat
}

//TableId returns the set tableId value for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) TableId() int {
	return r.tableId
}

//Frequency returns the frequency value for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Frequency() string {
	return r.frequency
}

//Years returns an array of strings for the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Years() string {
	return r.years
}

//AddYear appends a year to the years field of UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) AddYear(newYear string) {
	if r.years == "" {
		r.years = newYear
		return
	}
	r.years += "," + newYear
}

//Industry returns the value of the UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) Industry() string {
	return r.industry
}

//setIndustry sets the industry field of the UnderLyingGDPPerIndustryRequest object
func (r *UnderLyingGDPPerIndustryRequest) SetIndustry(newIndustry string) {
	r.industry = newIndustry
}

//String() a function which returns the current string representation of an UnderLyingGDPPerIndustryRequest object
func (r UnderLyingGDPPerIndustryRequest) String() string {

	var request = ""
	request += r.Frequency() + " "
	request += r.Industry() + " "
	request += strconv.Itoa(r.TableId()) + " "
	request += r.Years() + " "
	request += r.resultFormat
	return request
}

func (r *UnderLyingGDPPerIndustryRequest) Request() string {
	return apiAddress + "underlyingGDPbyIndustry" + "&Year=" + r.years + "&Industry=" + r.industry + "&tableId=" + strconv.Itoa(r.tableId) + "&Frequency=" + r.frequency + "&ResultFormat=" + r.resultFormat
}

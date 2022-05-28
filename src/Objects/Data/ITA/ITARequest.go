package ITA

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

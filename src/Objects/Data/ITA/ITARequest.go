package ITA

type ITA struct {
	indicator     string
	areaOrCountry string
	frequency     string
	year          string
	requestType   string
}

func (I *ITA) Indicator() string {
	return I.indicator
}

func (I *ITA) SetIndicator(indicator string) {
	I.indicator = indicator
}

func (I *ITA) AreaOrCountry() string {
	return I.areaOrCountry
}

func (I *ITA) SetAreaOrCountry(areaOrCountry string) {
	I.areaOrCountry = areaOrCountry
}

func (I *ITA) Frequency() string {
	return I.frequency
}

func (I *ITA) SetFrequency(frequency string) {
	I.frequency = frequency
}

func (I *ITA) Year() string {
	return I.year
}

func (I *ITA) SetYear(year string) {
	I.year = year
}

func (I *ITA) RequestType() string {
	return I.requestType
}

func (I *ITA) SetRequestType(requestType string) {
	I.requestType = requestType
}

func NewITA(indicator string, areaOrCountry string, frequency string, year string, requestType string) *ITA {
	return &ITA{indicator: indicator, areaOrCountry: areaOrCountry, frequency: frequency, year: year, requestType: requestType}
}

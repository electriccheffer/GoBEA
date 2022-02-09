package IntlServTrade

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

func (i *IntlServTradeRequest) SetTypeOfService(typeOfService string) {
	i.typeOfService = typeOfService
}

//TradeDirection is a method used to access the tradeDirection field of IntlServTradeRequest object
func (i *IntlServTradeRequest) TradeDirection() string {
	return i.tradeDirection
}

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

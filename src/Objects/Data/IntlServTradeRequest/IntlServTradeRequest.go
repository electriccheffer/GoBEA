package IntlServTradeRequest

type IntlServTradeRequest struct {
	typeOfService  string
	tradeDirection string
	affiliation    string
	areaOrCountry  string
}

func (i *IntlServTradeRequest) TypeOfService() string {
	return i.typeOfService
}

func (i *IntlServTradeRequest) SetTypeOfService(typeOfService string) {
	i.typeOfService = typeOfService
}

func (i *IntlServTradeRequest) TradeDirection() string {
	return i.tradeDirection
}

func (i *IntlServTradeRequest) SetTradeDirection(tradeDirection string) {
	i.tradeDirection = tradeDirection
}

func (i *IntlServTradeRequest) Affiliation() string {
	return i.affiliation
}

func (i *IntlServTradeRequest) SetAffiliation(affiliation string) {
	i.affiliation = affiliation
}

func (i *IntlServTradeRequest) AreaOrCountry() string {
	return i.areaOrCountry
}

func (i *IntlServTradeRequest) SetAreaOrCountry(areaOrCountry string) {
	i.areaOrCountry = areaOrCountry
}

package IntlServTrade

import "testing"

func TestIntlServTradeRequest_Affiliation(t *testing.T) {

	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nullCase.Affiliation() != "" {
		t.Error("nullCase.Affiliation() != ")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "Regional",
		areaOrCountry:  "",
	}
	if nonNullCase.Affiliation() != "Regional" {
		t.Error("nullCase.Affiliation() != Regional")
	}

}

func TestIntlServTradeRequest_SetAffiliation(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	nullCase.SetAffiliation("notGood")
	if nullCase.Affiliation() != "notGood" {
		t.Error("nullCase.Affiliation() != \"notGood\"")
	}

}

func TestIntlServTradeRequest_AreaOrCountry(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("nullCase.AreaOrCountry() != \"\"")
	}

	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "south-east",
	}

	if nonNullCase.AreaOrCountry() != "south-east" {
		t.Error("nonNullCase.AreaOrCountry() != \"south-east\"")
	}

}

func TestIntlServTradeRequest_SetAreaOrCountry(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	nullCase.SetAreaOrCountry("United States")
	if nullCase.AreaOrCountry() != "United States" {
		t.Error("nullCase.AreaOrCountry() != \"United States\"")
	}

}

func TestIntlServTradeRequest_TradeDirection(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nullCase.TradeDirection() != "" {
		t.Error("nullCase.TradeDirection() != \"\"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "backwards",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nonNullCase.TradeDirection() != "backwards" {
		t.Error("nonNullCase.TradeDirection() != \"backwards\"")
	}

}

func TestIntlServTradeRequest_SetTradeDirection(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	nullCase.SetTradeDirection("thisway")
	if nullCase.TradeDirection() != "thisway" {
		t.Error("nullCase.TradeDirection() != \"thisway\"")
	}

}

func TestIntlServTradeRequest_TypeOfService(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nullCase.TypeOfService() != "" {
		t.Error("nullCase.TypeOfService() != \"\"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "big",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}

	if nonNullCase.TypeOfService() != "big" {
		t.Error("nullCase.TypeOfService() != \"big\"")
	}

}

func TestIntlServTradeRequest_SetTypeOfService(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}

	nullCase.SetTypeOfService("good")
	if nullCase.TypeOfService() != "good" {
		t.Error("nullCase.TypeOfService() != \"good\"")
	}

}

func TestIntlServTradeRequest_toString(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
	}
	if nullCase.toString() != "   " {
		t.Error("nullCase.toString() != \"   \"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "good",
		tradeDirection: "theOtherWay",
		affiliation:    "notgood",
		areaOrCountry:  "usa",
	}

	if nonNullCase.toString() != "good theOtherWay notgood usa" {
		t.Error("nonNullCase.toString() != \"good theOtherWay notgood usa\"")
	}

}

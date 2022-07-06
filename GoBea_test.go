package GoBea

import (
	"testing"
)

func TestInputOutputDataRequest_Years(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.Years() != "" {
		t.Error("Test failed years != nil")
	}
	nonNullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "2012",
		returnFormat: "",
	}
	if nonNullCase.Years() != "2012" {
		t.Error("Test failed years != 2012")
	}

}

func TestInputOutputDataRequest_TableId(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed tableId != 0")
	}
	nonNullCase := InputOutputDataRequest{
		tableId:      1,
		years:        "",
		returnFormat: "",
	}
	if nonNullCase.TableId() != 1 {
		t.Error("Test failed tableId != 1")
	}
}

func TestInputOutputDataRequest_ReturnFormat(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.ReturnFormat() != "" {
		t.Error("Test failed returnFormat != nil")
	}
	nonNullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "JSON",
	}
	if nonNullCase.ReturnFormat() != "JSON" {
		t.Error("Test failed returnFormat != JSON")
	}
}

func TestInputOutputDataRequest_AddYear(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.Years() != "" {
		t.Error("Test failed years != nil")
	}
	nullCase.AddYear("2012")
	if nullCase.Years() != "2012" {
		t.Error("Test failed years != 2012")
	}
	nullCase.AddYear("2013")
	if nullCase.Years() != "2012,2013" {
		t.Error("Test failed years != 2012,2013")
	}
}

func TestInputOutputDataRequest_SetReturnFormat(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.ReturnFormat() != "" {
		t.Error("Test failed returnFormat != nil")
	}
	nullCase.SetReturnFormat("XML")
	if nullCase.ReturnFormat() != "XML" {
		t.Error("Test failed returnFormat != XML")
	}
}

func TestInputOutputDataRequest_SetYears(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.Years() != "" {
		t.Error("Test failed years != nil")
	}
	nullCase.SetYears("2012")
	if nullCase.Years() != "2012" {
		t.Error("Test failed years != 2012")
	}

}

func TestInputOutputDataRequest_SetTableId(t *testing.T) {
	nullCase := InputOutputDataRequest{
		tableId:      0,
		years:        "",
		returnFormat: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed tableId != 0")
	}
	nullCase.SetTableId(1)
	if nullCase.TableId() != 1 {
		t.Error("Test failed tableId != 1")
	}
}

func TestInputOutputDataRequest_String(t *testing.T) {
	nonNullCase := InputOutputDataRequest{
		tableId:      1,
		years:        "2012",
		returnFormat: "XML",
	}
	if nonNullCase.String() != "1 2012 XML" {
		t.Error("Test failed String != 1 2012 XML")
	}
}

func TestNewInputOutputDataRequest(t *testing.T) {
	newCase := NewInputOutputDataRequest(1, "2012,2013", "JSON")
	if newCase.TableId() != 1 {
		t.Error("Test failed tableId != 1")
	}
	if newCase.Years() != "2012,2013" {
		t.Error("Test failed years != 2012,2013")
	}
	if newCase.ReturnFormat() != "JSON" {
		t.Error("Test failed returnFormat != JSON")
	}
}

func TestInputOutputDataRequest_Request(t *testing.T) {
	testRequest := NewInputOutputDataRequest(56, "2010,2011,2012,2013", "json")
	if testRequest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=InputOutput&Year=2010,2011,2012,2013&tableID=56&ResultFormat=json" {
		t.Error("Test failed for InputOutputRequest")
	}
}

func TestFixedAssetsRequest_Year(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}

	if nullCase.Year() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "2012",
		requestType: "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed nonNull case should be set to 2012")
	}
}

func TestFixedAssetsRequest_RequestType(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null.")
	}
	nonNullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNull case should be set to JSON")
	}
}

func TestFixedAssetsRequest_TableName(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}

	if nullCase.TableName() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nonNullCase := FixedAssetsRequest{
		tableName:   "Tayne",
		year:        "",
		requestType: "",
	}

	if nonNullCase.TableName() != "Tayne" {
		t.Error("Test failed nonNull case should be set to Tayne")
	}
}

func TestFixedAssetsRequest_String(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "Tayne",
		year:        "2012",
		requestType: "JSON",
	}
	if nullCase.String() != "Tayne 2012 JSON" {
		t.Error("Test failed string should read Tayne 2012 JSON")
	}
}

func TestFixedAssetsRequest_SetYear(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed nullcase should be set to 2012")
	}
}

func TestFixedAssetsRequest_SetTableName(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed null case should be set to null")
	}

	nullCase.SetTableName("Tayne")
	if nullCase.TableName() != "Tayne" {
		t.Error("Test failed null case should be set to TAyne")
	}
}

func TestFixedAssetsRequest_SetRequestType(t *testing.T) {
	nullCase := FixedAssetsRequest{
		tableName:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be set to null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed request type should be set to JSON")
	}
}

func TestNewFixedAssetsRequest(t *testing.T) {
	newCase := NewFixedAssetsRequest("Tayne", "2012", "JSON")
	if newCase.TableName() != "Tayne" {
		t.Error("Test failed tableName should be Tayne")
	}
	if newCase.Year() != "2012" {
		t.Error("Test failed year should be 2012")
	}
	if newCase.RequestType() != "JSON" {
		t.Error("Test failed requestType should be set to JSON.")
	}

}

func TestFixedAssetsRequest_Request(t *testing.T) {
	testRequest := NewFixedAssetsRequest("FAAt405", "2015,2016", "json")
	if testRequest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=FixedAssets&TableName=FAAt405&Year=2015,2016&ResultFormat=json" {
		t.Error("Test failed FixedAssets request failed")
	}
}

func TestGDPByIndustry_Frequency(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase frequency should be set to null")
	}
	nonNullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "regular",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nonNullCase.Frequency() != "regular" {
		t.Error("Test failed nonNullCase frequency should be set to regular")
	}
}

func TestGDPByIndustry_Year(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed null case year should be set to 0")
	}
	nonNullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "2012",
		industry:    "",
		requestType: "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed nonNullCase year should be set to 2012 ")
	}
}

func TestGDPByIndustry_RequestType(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.requestType != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "XML",
	}
	if nonNullCase.requestType != "XML" {
		t.Error("Test failed nonNullCase should be XML.")
	}

}

func TestGDPByIndustry_SetRequestType(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.requestType != "" {
		t.Error("Test failed null case should be null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed requestType should be JSON")
	}
}

func TestGDPByIndustry_Industry(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Industry() != "" {
		t.Error("Test failed nullCase industry should be set to null")
	}
	nonNullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "manufacturing",
		requestType: "",
	}
	if nonNullCase.Industry() != "manufacturing" {
		t.Error("Test failed nonNullCase industry should be set to manufacturing")
	}
}

func TestGDPByIndustry_TableId(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed nullCase should be set to 0")
	}
	nonNullCase := GDPByIndustry{
		tableId:     12,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nonNullCase.TableId() != 12 {
		t.Error("Test failed nonNullCase tableId should be set to 12")
	}
}

func TestGDPByIndustry_SetFrequency(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be set to null")
	}
	nullCase.SetFrequency("regular")
	if nullCase.Frequency() != "regular" {
		t.Error("Test Failed frequency should be set to regular")
	}
}

func TestGDPByIndustry_SetIndustry(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Industry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetIndustry("manufacturing")
	if nullCase.Industry() != "manufacturing" {
		t.Error("Test failed nullCase industry field should be set to manufacturing.")
	}
}

func TestGDPByIndustry_SetTableId(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed null case tableId should be set to 0")
	}
	nullCase.SetTableId(2)
	if nullCase.TableId() != 2 {
		t.Error("Test failed null case tableId field should be set to 2")
	}
}

func TestGDPByIndustry_SetYear(t *testing.T) {
	nullCase := GDPByIndustry{
		tableId:     0,
		frequency:   "",
		year:        "",
		industry:    "",
		requestType: "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed null case year field should be set to 0")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed null case year field should be set to 2012")
	}
}

func TestGDPByIndustry_String(t *testing.T) {
	stringCase := GDPByIndustry{
		tableId:     1,
		frequency:   "regular",
		year:        "2012",
		industry:    "manufacturing",
		requestType: "XML",
	}
	if stringCase.String() != "1 regular 2012 manufacturing XML" {
		t.Error("Test failed string should read 1 regular 2012 manufacturing")
	}
}

func TestNewGDPByIndustry(t *testing.T) {
	newCase := NewGDPByIndustry(1, "regular", "2012", "manufacturing", "XML")
	if newCase.TableId() != 1 {
		t.Error("Test failed tableId should be 1")
	}
	if newCase.Frequency() != "regular" {
		t.Error("Test failed frequency should be set to regular")
	}
	if newCase.Year() != "2012" {
		t.Error("Test failed year should be set to 2012")
	}
	if newCase.Industry() != "manufacturing" {
		t.Error("Test failed industry should be set to manufacturing")
	}
	if newCase.RequestType() != "XML" {
		t.Error("Test failed requestType should be set to XML")
	}
}

func TestGDPByIndustry_Request(t *testing.T) {
	testRequest := NewGDPByIndustry(1, "A", "2012,2011", "ALL", "xml")
	if testRequest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=GDPByIndustry&Year=2012,2011&Industry=ALL&tableID=1&Frequency=A&ResultFormat=xml" {
		t.Error("Test failed GDPByIndustry Request")
	}
}

func TestIIPRequest_Year(t *testing.T) {

	nullCase := IIPRequest{
		typeOfInvestment: "static",
		component:        "none",
		frequency:        "none",
		year:             "",
		returnFormat:     "none",
	}
	if nullCase.Year() != "" {
		t.Error("Year test 1 failed.  Year should be empty.")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "2012",
		returnFormat:     "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Year test 2 failed. Year should be 2012")
	}
}

func TestIIPRequest_Frequency(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.Frequency() != "" {
		t.Error("Frequency test null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "non-frequent",
		year:             "",
		returnFormat:     "",
	}
	if nonNullCase.Frequency() != "non-frequent" {
		t.Error("Frequency test non-null case false")
	}

}

func TestIIPRequest_Component(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.Component() != "" {
		t.Error("Component test null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "tons",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nonNullCase.Component() != "tons" {
		t.Error("Component test nonNull case should be tons")
	}
}

func TestIIPRequest_TypeOfInvestment(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.TypeOfInvestment() != "" {
		t.Error("Type of investment null case should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "bad",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	if nonNullCase.TypeOfInvestment() != "bad" {
		t.Error("Test Failed type of investment should be bad")
	}

}

func TestIIPRequest_ReturnFormat(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}

	if nullCase.ReturnFormat() != "" {
		t.Error("Test Failed returnFormat should be null")
	}

	nonNullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "XML",
	}

	if nonNullCase.ReturnFormat() != "XML" {
		t.Error("Test Failed Return Format should be set to XML")
	}

}

func TestIIPRequest_SetComponent(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetComponent("JSON")
	if nullCase.Component() != "JSON" {
		t.Error("Test Failed Component should be set to JSON")
	}
}

func TestIIPRequest_SetFrequency(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetFrequency("infrequent")
	if nullCase.Frequency() != "infrequent" {
		t.Error("Test Failed Frequency should be set to infrequent")
	}

}

func TestIIPRequest_SetReturnFormat(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetReturnFormat("JSON")
	if nullCase.ReturnFormat() != "JSON" {
		t.Error("Test Failed return format should be set to JSON")
	}
}

func TestIIPRequest_SetYear(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test Failed Year should be set to 2012")
	}
}

func TestIIPRequest_SetTypeOfInvestment(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "",
		component:        "",
		frequency:        "",
		year:             "",
		returnFormat:     "",
	}
	nullCase.SetTypeOfInvestment("bad")
	if nullCase.TypeOfInvestment() != "bad" {
		t.Error("Test Failed type of investment should be bad")
	}
}

func TestIIPRequest_String(t *testing.T) {
	nullCase := IIPRequest{
		typeOfInvestment: "bad",
		component:        "multiple",
		frequency:        "not",
		year:             "2012",
		returnFormat:     "JSON",
	}
	if nullCase.String() != "multiple bad not 2012 JSON" {

	}
}

func TestNewIIPRequest(t *testing.T) {
	newCase := NewIIPRequest("FinAssets", "many", "not", "2001", "json")
	if newCase.TypeOfInvestment() != "FinAssets" {
		t.Error("Test failed typeOfInvestment != FinAssets")
	}
	if newCase.Component() != "many" {
		t.Error("Test failed component != many")
	}
	if newCase.Frequency() != "not" {
		t.Error("Test failed frequency != not")
	}
	if newCase.Year() != "2001" {
		t.Error("Test failed year != 2001")
	}
	if newCase.ReturnFormat() != "json" {
		t.Error("Test failed request != json")
	}
}

func TestIIPRequest_Request(t *testing.T) {
	testRequest := NewIIPRequest("FinAssetsExclFinDeriv", "ChgPosPrice", "A", "ALL", "json")
	if testRequest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=IIP&TypeOfInvestment=FinAssetsExclFinDeriv&Component=ChgPosPrice&Frequency=A&Year=ALL&ResultFormat=json" {
		t.Error("Test failed request does not match.")
	}
}

func TestIntlServTradeRequest_Affiliation(t *testing.T) {

	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.Affiliation() != "" {
		t.Error("nullCase.Affiliation() != ")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "Regional",
		areaOrCountry:  "",
		requestType:    "",
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
		requestType:    "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("nullCase.AreaOrCountry() != \"\"")
	}

	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "south-east",
		requestType:    "",
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
		requestType:    "",
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
		requestType:    "",
	}
	if nullCase.TradeDirection() != "" {
		t.Error("nullCase.TradeDirection() != \"\"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "backwards",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
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
		requestType:    "",
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
		requestType:    "",
	}
	if nullCase.TypeOfService() != "" {
		t.Error("nullCase.TypeOfService() != \"\"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "big",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
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
		requestType:    "",
	}

	nullCase.SetTypeOfService("good")
	if nullCase.TypeOfService() != "good" {
		t.Error("nullCase.TypeOfService() != \"good\"")
	}

}

func TestIntlServTradeRequest_String(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.String() != "    " {
		t.Error("nullCase.String() != \"   \"")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "good",
		tradeDirection: "theOtherWay",
		affiliation:    "notgood",
		areaOrCountry:  "usa",
		requestType:    "XML",
	}

	if nonNullCase.String() != "good theOtherWay notgood usa XML" {
		t.Error("nonNullCase.String() != \"good theOtherWay notgood usa XML\"")
	}

}

func TestNewIntlServTradeRequest(t *testing.T) {
	newCase := NewIntlServTradeRequest("good", "2012", "backwards", "non", "USA", "XML")
	if newCase.TypeOfService() != "good" {
		t.Error("Test failed typeOfService != good")
	}
	if newCase.TradeDirection() != "backwards" {
		t.Error("Test failed tradeDirection != backwards")
	}
	if newCase.Affiliation() != "non" {
		t.Error("Test failed affiliation != non")
	}
	if newCase.AreaOrCountry() != "USA" {
		t.Error("Test failed areaOrCountry != USA")
	}
	if newCase.RequestType() != "XML" {
		t.Error("Test failed requestType != USA")
	}
	if newCase.Year() != "2012" {
		t.Error("Test failed year != 2012")
	}
}

func TestIntlServTradeRequest_RequestType(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed request type != null")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "XML",
	}
	if nonNullCase.RequestType() != "XML" {
		t.Error("Test failed requestType != XML")
	}
}

func TestIntlServTradeRequest_SetRequestType(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed requestType should be nil")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed requestType != JSON")
	}
}

func TestIntlServTradeRequest_Year(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		year:           "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed year != nil")
	}
	nonNullCase := IntlServTradeRequest{
		typeOfService:  "",
		year:           "2012",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed year != 2012")
	}
}

func TestIntlServTradeRequest_SetYear(t *testing.T) {
	nullCase := IntlServTradeRequest{
		typeOfService:  "",
		year:           "",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed year is not empty")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed year != 2012")
	}
}

func TestIntlServTradeRequest_AddYear(t *testing.T) {
	addCase := IntlServTradeRequest{
		typeOfService:  "",
		year:           "2012",
		tradeDirection: "",
		affiliation:    "",
		areaOrCountry:  "",
		requestType:    "",
	}
	if addCase.Year() != "2012" {
		t.Error("Test failed year != 2012")
	}
	addCase.AddYear("2013")
	if addCase.Year() != "2012,2013" {
		t.Error("Test failed year != 2012,2013")
	}
}

func TestIntlServTradeRequest_Request(t *testing.T) {
	requestCase := NewIntlServTradeRequest("AllServiceTypes", "2014,2015", "Imports", "AllAffiliations", "Germany", "xml")
	if requestCase.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=IntlServTrade&TypeOfService=AllServiceTypes&TradeDirection=Imports&Affiliation=AllAffiliations&AreaOrCountry=Germany&Year=2014,2015&ResultFormat=xml" {
		t.Log(requestCase.Request())
		t.Error("Test failed Request is not correct")
	}
}

func TestITA_RequestType(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null.")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNullCase should be JSON")
	}

}

func TestITA_Year(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "2012",
		requestType:   "",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed nonNullCase year should be set to 2012")
	}

}

func TestITA_Frequency(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be null.")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "moderate",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed nonNullcase should be moderate")
	}
}

func TestITA_AreaOrCountry(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "usa",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.AreaOrCountry() != "usa" {
		t.Error("Test failed nonNullCase should be set to usa")
	}

}

func TestITA_Indicator(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Indicator() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nonNullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "moderate",
		year:          "",
		requestType:   "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed nonNullCase should be set to moderate")
	}

}

func TestITA_SetAreaOrCountry(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.AreaOrCountry() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetAreaOrCountry("usa")
	if nullCase.AreaOrCountry() != "usa" {
		t.Error("Test failed areaOrCountry should be set to usa")
	}
}

func TestITA_SetFrequency(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed nullCase should be null.")
	}
	nullCase.SetFrequency("moderate")
	if nullCase.Frequency() != "moderate" {
		t.Error("Test failed nullCase frequency should be moderate")
	}
}

func TestITA_SetIndicator(t *testing.T) {

	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Indicator() != "" {
		t.Error("Test failed nullCase should be set to null")
	}

	nullCase.SetIndicator("positive")
	if nullCase.Indicator() != "positive" {
		t.Error("Test failed nullCase indicator should be set to positive.")
	}
}

func TestITA_SetRequestType(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be set to null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("nullCase should be set to JSON")
	}
}

func TestITA_SetYear(t *testing.T) {
	nullCase := ITA{
		indicator:     "",
		areaOrCountry: "",
		frequency:     "",
		year:          "",
		requestType:   "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be set to null.")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed nullCase should be set to 2012")
	}
}

func TestITA_String(t *testing.T) {
	firstCase := ITA{
		indicator:     "positive",
		areaOrCountry: "usa",
		frequency:     "moderate",
		year:          "2012",
		requestType:   "JSON",
	}
	if firstCase.String() != "positive usa moderate 2012 JSON" {
		t.Error("Test failed string form is not correct")
	}
}

func TestITA_Request(t *testing.T) {
	requestCase := NewITA("BalGds", "China", "A", "2011,2012", "xml")
	if requestCase.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=ITA&Indicator=BalGds&AreaOrCountry=China&Frequency=A&Year=2011,2012&ResultFormat=xml" {
		t.Error("Test failed request string is invalid")
	}
}

func TestNewITA(t *testing.T) {

	constructorCase := NewITA("positive", "usa", "moderate", "2012", "JSON")
	if constructorCase.Indicator() != "positive" && constructorCase.AreaOrCountry() != "usa" && constructorCase.Frequency() != "moderate" && constructorCase.Year() != "2012" && constructorCase.RequestType() != "JSON" {
		t.Error("Test failed for constructor case.")
	}
}

func TestMNERequest_Classification(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Classification() != "" {
		t.Error("Test failed classification should me empty")
	}
	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "negative",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.Classification() != "negative" {
		t.Error("Test failed classification should be set to negative")
	}
}

func TestMNERequest_Country(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Country() != "" {
		t.Error("Test failed null case should be null.")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "India",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.Country() != "India" {
		t.Error("Test failed nonNull case should be set to India.")
	}
}

func TestMNERequest_DirectionOfInvestment(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.DirectionOfInvestment() != "" {
		t.Error("Test failed null case should be null")
	}
	nonNullCase := MNERequest{
		directionOfInvestment: "reverse",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nonNullCase.DirectionOfInvestment() != "reverse" {
		t.Error("Test failed direction of investment should be reverse")
	}
}

func TestMNERequest_Industry(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.Industry() != "" {
		t.Error("Test failed industry should be set to null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "news",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.Industry() != "news" {
		t.Error("Test failed industry should be set to news")
	}
}

func TestMNERequest_SeriesId(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.SeriesId() != 0 {
		t.Error("Test failed seriesId should be set to 0")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              1,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.SeriesId() != 1 {
		t.Error("Test failed seriesId should be set to 1")
	}

}

func TestMNERequest_GetFootnotes(t *testing.T) {

	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.GetFootnotes() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "yes",
		requestType:           "",
	}

	if nonNullCase.GetFootnotes() != "yes" {
		t.Error("Test failed nonNull case should be yes")
	}
}

func TestMNERequest_RequestType(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.RequestType() != "" {
		t.Error("Test failed nullCase request type should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "JSON",
	}

	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed nonNullCase should be JSON")
	}
}

func TestMNERequest_Year(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be null")
	}

	nonNullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "2012",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}

	if nonNullCase.Year() != "2012" {
		t.Error("Test failed year should be set to 2012")
	}
}

func TestMNERequest_SetCountry(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Country() != "" {
		t.Error("Test failed nullCase should be set to null")
	}
	nullCase.SetCountry("USA")
	if nullCase.Country() != "USA" {
		t.Error("Test failed nullCAse should be set to USA")
	}
}

func TestMNERequest_SetClassification(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Country() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetClassification("non")
	if nullCase.Classification() != "non" {
		t.Error("Test failed nullCase classificaiton should be set to non.")
	}
}

func TestMNERequest_SetYear(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should be null")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed nullCase should be 2012")
	}
}

func TestMNERequest_SetDirectionOfInvestment(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.DirectionOfInvestment() != "" {
		t.Error("Test failed nullCase should be null.")
	}
	nullCase.SetDirectionOfInvestment("reverse")
	if nullCase.DirectionOfInvestment() != "reverse" {
		t.Error("Test failed nullCase directionOfInvestment should be set to reverse")
	}
}

func TestMNERequest_SetGetFootnotes(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.GetFootnotes() != "" {
		t.Error("Test failed null case should be null")
	}
	nullCase.SetGetFootnotes("yes")
	if nullCase.GetFootnotes() != "yes" {
		t.Error("Test failed Get Footnotes should be yes")
	}
}

func TestMNERequest_SetIndustry(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.Industry() != "" {
		t.Error("Test failed null case should be null")
	}

	nullCase.SetIndustry("news")
	if nullCase.Industry() != "news" {
		t.Error("Test failed industry should be set to news")
	}
}

func TestMNERequest_SetRequestType(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed null case should be set to JSON")
	}
}

func TestMNERequest_SetSeriesId(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "",
		seriesId:              0,
		classification:        "",
		year:                  "",
		country:               "",
		industry:              "",
		getFootnotes:          "",
		requestType:           "",
	}
	if nullCase.SeriesId() != 0 {
		t.Error("Test failed nullCase should be set to 0")
	}
	nullCase.SetSeriesId(1)
	if nullCase.SeriesId() != 1 {
		t.Error("Test failed nullCase should be set to 1")
	}
}

func TestMNERequest_String(t *testing.T) {
	nullCase := MNERequest{
		directionOfInvestment: "reverse",
		seriesId:              0,
		classification:        "bad",
		year:                  "2012",
		country:               "usa",
		industry:              "news",
		getFootnotes:          "yes",
		requestType:           "JSON",
	}
	if nullCase.String() != "reverse 0 bad 2012 usa news yes JSON" {
		t.Error("Test failed for testing string")
	}
}

func TestNewMNERequest(t *testing.T) {
	newCase := NewMNERequest("reverse", 1, "bad", "2012", "usa", "news", "yes", "JSON")
	if newCase.DirectionOfInvestment() != "reverse" {
		t.Error("Test failed for direction of investment")
	}
	if newCase.SeriesId() != 1 {
		t.Error("Test failed for series id")
	}
	if newCase.Classification() != "bad" {
		t.Error("Test failed for classification")
	}
	if newCase.Year() != "2012" {
		t.Error("Test failed for year")
	}
	if newCase.Country() != "usa" {
		t.Error("Test failed for year")
	}
	if newCase.Industry() != "news" {
		t.Error("Test failed for news")
	}
	if newCase.GetFootnotes() != "yes" {
		t.Error("Test failed for get footnotes")
	}
	if newCase.RequestType() != "JSON" {
		t.Error("Test failed for request type")
	}
}

func TestMNERequest_Request(t *testing.T) {
	requestTest := NewMNERequest("Inward", 0, "CountryByIndustry", "2013,2012,2011,2010", "308", "3000", "", "xml")
	if requestTest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=MNE&Year=2013,2012,2011,2010&Country=308&Industry=3000&DirectionOfInvestment=Inward&Classification=CountryByIndustry&ResultFormat=xml" {
		t.Log(requestTest.Request())
		t.Error("Test failed MNE Request")
	}
}

func TestNIPADataRequest_TableName(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed tableName is not nil")
	}
	nonNullCase := NIPADataRequest{
		tableName:   "someName",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nonNullCase.TableName() != "someName" {
		t.Error("someName")
	}
}

func TestNIPADataRequest_SetTableName(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed tableName is not nil")
	}
	nullCase.SetTableName("someName")
	if nullCase.TableName() != "someName" {
		t.Error("Test failed tableName != someName")
	}
}

func TestNIPADataRequest_Request(t *testing.T) {
	requestTest := NewNIPADataRequest("T10101", "A,Q", "ALL", "xml")
	if requestTest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=NIPA&TableName=T10101&Frequency=A,Q&Year=ALL&ResultFormat=xml" {
		t.Log(requestTest.Request())
		t.Error("Test failed request is not valid")
	}
}

func TestNIPADataRequest_Frequency(t *testing.T) {

	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}

	if nullCase.Frequency() != "" {
		t.Error("frequency != nil")
	}

	nonNullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "2003,20",
		years:       "",
		requestType: "",
	}

	if nonNullCase.Frequency() != "2003,20" {
		t.Error("nonNullCase.frequency != 2003,20")
	}
}

func TestNIPADataRequest_Years(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nullCase.Years() != "" {
		t.Error("nullCase.years != empty")
	}
	nonNullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "2000, 2012",
		requestType: "",
	}
	if nonNullCase.Years() != "2000, 2012" {
		t.Error("nonNullCase.years != 2000, 2012")
	}

}

func TestNIPADataRequest_SetYears(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	nullCase.SetYears("2000, 2012")
	if nullCase.Years() != "2000, 2012" {
		t.Error("nullCase.Years()[0] != 2000")
	}
	nullCase.SetYears("1992, 1193")
	if nullCase.Years() != "1992, 1193" {
		t.Error("nullCase.Years() != 1992, 1193")
	}

}

func TestNIPADataRequest_SetFrequency(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	nullCase.SetFrequency("2012, 2000")

	if nullCase.Frequency() != "2012, 2000" {
		t.Error("nullCase.Frequency()[0] != 2012")
	}
	nullCase.SetFrequency("1999, 1998")
	if nullCase.Frequency() != "1999, 1998" {
		t.Error("nullCase.Frequency() != 1999, 1998")
	}
}

func TestNIPADataRequest_AddFrequency(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName: "",
		frequency: "",
		years:     "",
	}

	nullCase.AddFrequency("2012")
	if nullCase.Frequency() != "2012" {
		t.Error("nullCase.Frequency() != 2012")
	}
	nullCase.AddFrequency("2013")
	if nullCase.Frequency() != "2012,2013" {
		t.Error("nullCase.Frequency() != 2012,2013")
	}
}

func TestNIPADataRequest_AddYear(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}

	nullCase.AddYear("210")
	if nullCase.Years() != "210" {
		t.Error("nullCase.Years() != 210")
	}
	nullCase.AddYear("2000")
	if nullCase.Years() != "210,2000" {
		t.Error("nullCase.Years() != 210, 2000")
	}

}

func TestNIPADataRequest_RequestType(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed requestType != nil")
	}
	nonNullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed requestType != nil")
	}
}

func TestNIPADataRequest_SetRequestType(t *testing.T) {
	nullCase := NIPADataRequest{
		tableName:   "",
		frequency:   "",
		years:       "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed requestType != nil")
	}
	nullCase.SetRequestType("JSON")
	if nullCase.RequestType() != "JSON" {
		t.Error("Test failed requestType != JSON")
	}
}

func TestNIPADataRequest_String(t *testing.T) {
	nonNullCase := NIPADataRequest{
		tableName:   "Ryan",
		frequency:   "2003,2004",
		years:       "122,332",
		requestType: "JSON",
	}

	if nonNullCase.String() != "Ryan 2003,2004 122,332 JSON" {
		t.Log(nonNullCase.String())
		t.Error("nonNullCase String()")
	}
}

func TestNewNIPADataRequest(t *testing.T) {
	newCase := NIPADataRequest{
		tableName:   "",
		frequency:   "good",
		years:       "2012",
		requestType: "JSON",
	}

	if newCase.TableName() != "" {
		t.Error("Test failed tableName != nil")
	}
	if newCase.Frequency() != "good" {
		t.Error("Test failed Frequency != good")
	}

	if newCase.Years() != "2012" {
		t.Error("Test failed Years != 2012")
	}

	if newCase.RequestType() != "JSON" {
		t.Error("Test failed RequestType != JSON")
	}
}

func TestNIUnderlyingDetailRequest_Frequency(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}

	if nullCase.Frequency() != "" {
		t.Error("Test failed Frequency should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "moderate",
		year:      "",
	}
	if nonNullCase.Frequency() != "moderate" {
		t.Error("Test failed Frequency should be moderate")
	}

}

func TestNIUnderlyingDetailRequest_Year(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed year field should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "2012",
	}
	if nonNullCase.Year() != "2012" {
		t.Error("Test failed year field should be 2012")
	}

}

func TestNIUnderlyingDetailRequest_TableName(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}

	if nullCase.TableName() != "" {
		t.Error("Test failed Null Case should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName: "Tayne",
		frequency: "",
		year:      "",
	}

	if nonNullCase.TableName() != "Tayne" {
		t.Error("Test Failed nonNull case tableName should be Tayne")
	}
}

func TestNIUnderlyingDetailRequest_SetFrequency(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Frequency() != "" {
		t.Error("Test failed null case should start null")
	}
	nullCase.SetFrequency("moderate")
	if nullCase.Frequency() != "moderate" {
		t.Error("Test failed Frequency should be moderate")
	}

}

func TestNIUnderlyingDetailRequest_SetYear(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.Year() != "" {
		t.Error("Test failed nullCase should start null")
	}
	nullCase.SetYear("2012")
	if nullCase.Year() != "2012" {
		t.Error("Test failed year should be 2012")
	}

}

func TestNIUnderlyingDetailRequest_SetTableName(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "",
		frequency: "",
		year:      "",
	}
	if nullCase.TableName() != "" {
		t.Error("Test failed tableName should start null")
	}
	nullCase.SetTableName("Tayne")
	if nullCase.TableName() != "Tayne" {
		t.Error("Test failed tableName should be Tayne")
	}

}

func TestNIUnderlyingDetailRequest_RequestType(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed null case should be null")
	}

	nonNullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "JSON",
	}
	if nonNullCase.RequestType() != "JSON" {
		t.Error("Test failed request type should be JSON")
	}

}

func TestNIUnderlyingDetailRequest_SetRequestType(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName:   "",
		frequency:   "",
		year:        "",
		requestType: "",
	}
	if nullCase.RequestType() != "" {
		t.Error("Test failed nullCase should start as null")
	}

	nullCase.SetRequestType("XML")
	if nullCase.RequestType() != "XML" {
		t.Error("Test failed nullCase should be XML.")
	}
}

func TestNIUnderlyingDetailRequest_String(t *testing.T) {
	nullCase := NIUnderlyingDetailRequest{
		tableName: "Tayne",
		frequency: "frequent",
		year:      "2012",
	}

	if nullCase.String() != "Tayne frequent 2012" {
		t.Error("Test failed string ouput should be: Tayne frequent 2012")
	}
}

func TestNewNIUnderlyingDetailRequest(t *testing.T) {
	constructCase := NewNIUnderlyingDetailRequest("Tayne", "regular", "2012", "JSON")
	if constructCase.Year() != "2012" {
		t.Error("Test failed constructCase year should be 2012")
	}
	if constructCase.RequestType() != "JSON" {
		t.Error("Test failed constructCase request type should be JSON ")
	}
	if constructCase.TableName() != "Tayne" {
		t.Error("Test failed constructCase tableName should be Tayne")
	}
}

func TestNIUnderlyingDetailRequest_Request(t *testing.T) {
	requestTest := NewNIUnderlyingDetailRequest("U20305", "A,Q,M", "ALL", "xml")
	if requestTest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=NIUnderlyingDetail&TableName=U20305&Frequency=A,Q,M&Year=ALL&ResultFormat=xml" {
		t.Log(requestTest.Request())
		t.Error("Test failed NIU request is not correct")
	}
}

func TestRegionalRequest_TableName(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	if nullCase.tableName != "" {
		t.Error("null case TableName")
	}

	nonNullCase := RegionalRequest{
		tableName: "Table",
		lineCode:  2,
		geoFips:   "392",
		year:      "2005",
	}
	if nonNullCase.tableName != "Table" {
		t.Error("nonNullCase TableName")
	}

}

func TestRegionalRequest_RequestFormat(t *testing.T) {
	nullCase := RegionalRequest{
		tableName:     "",
		lineCode:      0,
		geoFips:       "",
		year:          "",
		requestFormat: "",
	}
	if nullCase.RequestFormat() != "" {
		t.Error("Test failed RequestFormat is not nil")
	}
	nonNullCase := RegionalRequest{
		tableName:     "",
		lineCode:      0,
		geoFips:       "",
		year:          "",
		requestFormat: "xml",
	}
	if nonNullCase.RequestFormat() != "xml" {
		t.Error("Test failed RequestFormat != xml")
	}
}

func TestRegionalRequest_LineCode(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	if nullCase.lineCode != 0 {

		t.Error("nullCase LineCode")

	}
	nonNullCase := RegionalRequest{
		tableName: "name",
		lineCode:  2,
		geoFips:   "what is a fip",
		year:      "2003",
	}
	if nonNullCase.lineCode != 2 {
		t.Error("nonNullCase LineCode")
	}
}

func TestRegionalRequest_GeoFips(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	if nullCase.geoFips != "" {
		t.Error("nullCase Geofips")
	}
	nonNullCase := RegionalRequest{
		tableName: "table",
		lineCode:  3,
		geoFips:   "fils",
		year:      "3002",
	}
	if nonNullCase.geoFips != "fils" {
		t.Error("nonNullCase Geofips")
	}

}

func TestRegionalRequest_Year(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	if nullCase.year != "" {
		t.Error("nullCase Year")
	}
	nonNullCase := RegionalRequest{
		tableName: "ryan",
		lineCode:  3,
		geoFips:   "58",
		year:      "2003",
	}
	if nonNullCase.year != "2003" {
		t.Error("nonNullCase Year")
	}

}

func TestRegionalRequest_SetRequestFormat(t *testing.T) {
	nullCase := RegionalRequest{
		tableName:     "",
		lineCode:      0,
		geoFips:       "",
		year:          "",
		requestFormat: "",
	}
	if nullCase.RequestFormat() != "" {
		t.Error("Test failed RequestFormat != nil")
	}
	nullCase.SetRequestFormat("xml")
	if nullCase.RequestFormat() != "xml" {
		t.Error("Test failed RequestFormat != xml")
	}
}

func TestRegionalRequest_Request(t *testing.T) {

}

func TestRegionalRequest_SetTableName(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	nullCase.SetTableName("newName")
	if nullCase.TableName() != "newName" {
		t.Error("nullCase SetTableName")
	}

	nullCase.SetTableName("otherName")
	if nullCase.TableName() != "otherName" {
		t.Error("nullCase2 SetTableName")
	}
}

func TestRegionalRequest_SetLineCode(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}

	nullCase.SetLineCode(20)
	if nullCase.LineCode() != 20 {
		t.Error("nullCase LineCode")
	}
	nullCase.SetLineCode(30)
	if nullCase.LineCode() != 30 {
		t.Error("nullCase2 LineCode")
	}
}

func TestRegionalRequest_SetGeoFips(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	nullCase.SetGeoFips("newFip")
	if nullCase.GeoFips() != "newFip" {
		t.Error("nullCase GeoFips")
	}

}

func TestRegionalRequest_SetYear(t *testing.T) {
	nullCase := RegionalRequest{
		tableName: "",
		lineCode:  0,
		geoFips:   "",
		year:      "",
	}
	nullCase.SetYear("2003")
	if nullCase.Year() != "2003" {
		t.Error("nullCase SetYear")
	}
	nullCase.SetYear("2004")
	if nullCase.Year() != "2004" {
		t.Error("nullCase2 SetYear")
	}
}

func TestNewRegionalRequest(t *testing.T) {
	newCase := NewRegionalRequest("someName", 0, "A,Q,M", "ALL", "xml")
	if newCase.TableName() != "someName" {
		t.Error("Test failed TableName != someName")
	}
	if newCase.LineCode() != 0 {
		t.Error("Test failed LineCode != 0")
	}
	if newCase.GeoFips() != "A,Q,M" {
		t.Error("Test failed GeoFips != A,Q,M")
	}
	if newCase.Year() != "ALL" {
		t.Error("Test failed Year != ALL")
	}
	if newCase.RequestFormat() != "xml" {
		t.Error("Test failed RequestFormat != xml")
	}
}

func TestRegionalRequest_String(t *testing.T) {
	stringCase := RegionalRequest{
		tableName:     "Ryan",
		lineCode:      0,
		geoFips:       "AM",
		year:          "2013",
		requestFormat: "XML",
	}

	if stringCase.String() != "Ryan 0 AM 2013 XML" {
		t.Log(stringCase.String())
		t.Error("Test failed string case != Ryan 0 AM 2013 XML")
	}
}

func TestUnderLyingGDPByIndustryRequest_Frequency(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nullCase.Frequency() != "" {
		t.Error("nullCase.Frequency() != null")
	}

}

func TestUnderLyingGDPByIndustryRequest_Years(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nullCase.Years() != "" {
		t.Error("")
	}

}

func TestUnderLyingGDPByIndustryRequest_Industry(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	if nullCase.Industry() != "" {
		t.Error("nullCase.Industry() != null ")
	}

}

func TestUnderLyingGDPByIndustryRequest_SetTableId(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	nullCase.SetTableId(12)
	if nullCase.TableId() != 12 {
		t.Error("nullCase.TableId() != 12")
	}

	nullCase.SetTableId(13)
	if nullCase.TableId() != 13 {
		t.Error("nullCase.TableId() != 13")
	}
}

func TestUnderLyingGDPByIndustryRequest_SetFrequency(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	nullCase.SetFrequency("another")
	if nullCase.Frequency() != "another" {
		t.Error("nullCase.Frequency() != \"another\"")
	}
}

func TestUnderLyingGDPByIndustryRequest_SetYears(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	nullCase.SetYears("2004, 2003, 2007")
	if nullCase.Years() != "2004, 2003, 2007" {
		t.Error("nullCase.Years()[0] != \"2004, 2003, 2007\"")
	}

}

func TestUnderLyingGDPByIndustryRequest_SetIndustry(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	nullCase.SetIndustry("steel")
	if nullCase.Industry() != "steel" {
		t.Error("nullCase.Industry() != \"steel\"")
	}

	nullCase.SetIndustry("tech")
	if nullCase.Industry() != "tech" {
		t.Error("nullCase.Industry() != \"tech\"")
	}

}

func TestUnderLyingGDPByIndustryRequest_AddYear(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}

	nullCase.AddYear("2010")
	if nullCase.Years() != "2010" {
		t.Error("nullCase.Years()[0] != \"2010\"")
	}

}

func TestUnderLyingGDPByIndustryRequest_String(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      2,
		frequency:    "yearly",
		years:        "2000,2010",
		industry:     "steel",
		resultFormat: "JSON",
	}

	if nullCase.String() != "yearly steel 2 2000,2010 JSON" {
		t.Error("nullCase.String() != yearly steel 2 2000,2010 JSON")
	}

}

func TestUnderLyingGDPByIndustryRequest_ResultFormat(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nullCase.ResultFormat() != "" {
		t.Error("Test failed resultFormat != nil")
	}
	nonNullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "XML",
	}
	if nonNullCase.ResultFormat() != "XML" {
		t.Error("Test failed resultFormat != XML")
	}
}

func TestUnderLyingGDPByIndustryRequest_SetResultFormat(t *testing.T) {
	nonNullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nonNullCase.ResultFormat() != "" {
		t.Error("Test failed resultFormat != nil")
	}
	nonNullCase.SetResultFormat("JSON")
	if nonNullCase.ResultFormat() != "JSON" {
		t.Error("Test failed resultFormat != JSON")
	}
}

func TestNewUnderLyingGDPByIndustryRequest(t *testing.T) {
	newCase := NewUnderLyingGDPByIndustryRequest(1, "moderate", "2012", "Steel", "JSON")
	if newCase.TableId() != 1 {
		t.Error("Test failed tableId != 1")
	}
	if newCase.Frequency() != "moderate" {
		t.Error("Test failed Frequency != moderate")
	}
	if newCase.Years() != "2012" {
		t.Error("Test failed years != 2012")
	}
	if newCase.Industry() != "Steel" {
		t.Error("Test failed Industry != Steel")
	}
	if newCase.ResultFormat() != "JSON" {
		t.Error("Test failed ResultFormat != JSON")
	}
}

func TestUnderLyingGDPByIndustryRequest_TableId(t *testing.T) {
	nullCase := UnderLyingGDPByIndustryRequest{
		tableId:      0,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nullCase.TableId() != 0 {
		t.Error("Test failed TableId != 0")
	}
	nonNullCase := UnderLyingGDPByIndustryRequest{
		tableId:      1,
		frequency:    "",
		years:        "",
		industry:     "",
		resultFormat: "",
	}
	if nonNullCase.TableId() != 1 {
		t.Error("Test failed TableId != 1")
	}
}

func TestUnderLyingGDPByIndustryRequest_Request(t *testing.T) {
	testRequest := NewUnderLyingGDPByIndustryRequest(210, "A", "2013,2012", "ALL", "xml")
	if testRequest.Request() != "https://apps.bea.gov/api/data/?&UserID=D27C93B3-297B-4BA6-826C-25D85653E54D&method=GetData&DataSetName=underlyingGDPByIndustry&Year=2013,2012&Industry=ALL&tableID=210&Frequency=A&ResultFormat=xml" {
		t.Error("Test failed request is not correct")
	}

}

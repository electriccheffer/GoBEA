package InternationalInvestmentPosition

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

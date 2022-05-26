package NIUnderlyingDetailRequest

type NIUnderlyingDetailRequest struct {
	tableName   string
	frequency   string
	year        string
	requestType string
}

func (N *NIUnderlyingDetailRequest) RequestType() string {
	return N.requestType
}

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

//NewNIUnderlyingDetailRequest is a constuctor for anNIUnderlyingDetailRequest
func NewNIUnderlyingDetailRequest(tableName string, frequency string, year string) *NIUnderlyingDetailRequest {
	return &NIUnderlyingDetailRequest{tableName: tableName, frequency: frequency, year: year}
}

//String is a method to return a string version of an NewNIUnderlyingDetailRequest
func (N NIUnderlyingDetailRequest) String() string {
	return N.tableName + " " + N.frequency + " " + N.year
}

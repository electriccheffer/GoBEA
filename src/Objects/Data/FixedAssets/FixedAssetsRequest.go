package FixedAssets

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

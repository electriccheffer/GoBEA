package FixedAssets

type FixedAssetsRequest struct {
	tableName   string
	year        string
	requestType string
}

func (f *FixedAssetsRequest) TableName() string {
	return f.tableName
}

func (f *FixedAssetsRequest) SetTableName(tableName string) {
	f.tableName = tableName
}

func (f *FixedAssetsRequest) Year() string {
	return f.year
}

func (f *FixedAssetsRequest) SetYear(year string) {
	f.year = year
}

func (f *FixedAssetsRequest) RequestType() string {
	return f.requestType
}

func (f *FixedAssetsRequest) SetRequestType(requestType string) {
	f.requestType = requestType
}

func NewFixedAssetsRequest(tableName string, year string, requestType string) *FixedAssetsRequest {
	return &FixedAssetsRequest{tableName: tableName, year: year, requestType: requestType}
}

func (f *FixedAssetsRequest) String() string {
	return f.tableName + " " + f.year + " " + f.requestType
}

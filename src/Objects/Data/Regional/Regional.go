package Regional

import "strconv"

type Regional struct {
	tableName string
	lineCode  int
	geoFips   string
	year      string
}

func (r *Regional) TableName() string {
	return r.tableName
}

func (r *Regional) SetTableName(tableName string) {
	r.tableName = tableName
}

func (r *Regional) LineCode() int {
	return r.lineCode
}

func (r *Regional) SetLineCode(lineCode int) {
	r.lineCode = lineCode
}

func (r *Regional) GeoFips() string {
	return r.geoFips
}

func (r *Regional) SetGeoFips(geoFips string) {
	r.geoFips = geoFips
}

func (r *Regional) Year() string {
	return r.year
}

func (r *Regional) SetYear(year string) {
	r.year = year
}

func NewRegional(tableName string, lineCode int, geoFips string, year string) *Regional {
	return &Regional{tableName: tableName, lineCode: lineCode, geoFips: geoFips, year: year}
}

func (r *Regional) toString() string {

	return r.tableName + strconv.Itoa(r.lineCode) + r.geoFips + r.year
}

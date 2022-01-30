package InputOutputData

import "strconv"

type InputOutputDataReq struct {
	tableId      int
	years        []string
	returnFormat string
}

func (r InputOutputDataReq) ReturnFormat() string {
	return r.returnFormat
}

func (r InputOutputDataReq) setReturnFormat(format string) {
	r.returnFormat = format
}

func (r InputOutputDataReq) Years() []string {
	return r.years
}

func (r InputOutputDataReq) setYears(newYears []string) {

	for _, year := range newYears {

		r.years = append(newYears, year)

	}
}

func (r InputOutputDataReq) TableId() int {
	return r.tableId
}

func (r InputOutputDataReq) setTableId(id int) {

	r.tableId = id

}

func (r InputOutputDataReq) toString() (string, int) {

	var returnYears string
	for _, year := range r.years {

		returnYears += year

	}

	return strconv.Itoa(r.tableId) + returnYears + r.returnFormat, 0

}

package InputOutput

import (
	"BEALIb/src/Objects/Data/RequestInterface"
	"strconv"
)

//InputOutputDataReq is a type struct which embodies all the required fields
//	of a request made for Input/Output data
type InputOutputDataReq struct {
	tableId      int
	years        []string
	returnFormat string
	RequestInterface.Request
}

//ReturnFormat provides a method for returning the value of the returnFormat
//	variable of the InputOutputDataReq struct
func (r InputOutputDataReq) ReturnFormat() string {
	return r.returnFormat
}

//setReturnFormat provides a method for setting the returnFormat variable of
//	of the InputOutputDataReq
func (r *InputOutputDataReq) setReturnFormat(format string) {
	r.returnFormat = format
}

//Years provides a method for returning the years[] variable of the InputOutputDataReq
//struct
func (r InputOutputDataReq) Years() []string {
	return r.years
}

//addYear provides a method for appending a year to the years[] variable
func (r *InputOutputDataReq) addYear(year string) {
	if r.years == nil {
		r.years = []string{year}
		return
	}
	r.years = append(r.years, year)
}

//setYears provides a method for changing the entire years[] with a new array of years[]
//for the InputOutputDataReq struct
func (r *InputOutputDataReq) setYears(newYears []string) {
	r.years = newYears
}

//TableId provides a method for getting the tableId of the InputOutputDataReq
//struct
func (r InputOutputDataReq) TableId() int {
	return r.tableId
}

//setTableId provides a method for setting the tableId field of the InputOutputDataReq
//struct
func (r *InputOutputDataReq) setTableId(id int) {

	r.tableId = id

}

//toString provides a method for returning the InputOutputDataReq struct as a
//string in the format of a BEA request
func (r *InputOutputDataReq) toString() (string, int) {

	var returnYears string
	for _, year := range r.years {

		returnYears += " " + year

	}
	return strconv.Itoa(r.tableId) + returnYears + " " + r.returnFormat, 0

}

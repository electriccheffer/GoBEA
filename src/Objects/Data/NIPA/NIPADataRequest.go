package NIPA

//NIPADataRequest is a struct used to represent a NIPA request
//to the BEA.  It consists of a tableId, frequency[], years[].
type NIPADataRequest struct {
	tableId   int
	frequency []string
	years     []string
}

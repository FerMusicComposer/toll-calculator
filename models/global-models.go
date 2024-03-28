package models

//=====================
// ON-BOARD UNIT TYPES
//=====================
type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//=========================
// INVOICE AGREGATOR TYPES
//=========================
type Distance struct {
	Value float64 `json:"value"`
	OBUID int     `json:"obuID"`
	Unix  int64   `json:"unix"`
}

type Invoice struct {
	OBUID         int     `json:"obuID"`
	TotalDistance float64 `json:"totalDistance"`
	AmountDue     float64 `json:"amountDue"`
}

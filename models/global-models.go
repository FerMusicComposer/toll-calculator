package models

//=====================
// ON-BOARD UNIT TYPES
//=====================
type OBUData struct {
	OBUID     int     `json:"obuID"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//=====================
// DATA RECEIVER TYPES
//=====================

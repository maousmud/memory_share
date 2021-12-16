package model

import "time"

// Book :
type Movie struct {
	Id                int
	Title             string
	Movie_file        []uint8
	Register_username string
	Remarks           string
	Memory_shubetsu   string
	Created_date      time.Time
	Update_date       time.Time
	Del_flg           bool
}

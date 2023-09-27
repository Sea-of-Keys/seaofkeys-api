package models


// ######TODO######
// ####IDEAS####
// Make it so ID is the day
// Day 1=Monday, 2=Tuesday, 3=Wensday 4=Thursday 5=Friday 6=Saturday 7=Sunday

type Weekday struct {
	ID uint `json:"id" gorm:"primaryKey"`
	Day int `json:"day"`
}

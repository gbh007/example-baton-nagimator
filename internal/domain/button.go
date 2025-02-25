package domain

type ButtonPressed struct {
	Year  int   `json:"year" gorm:"primaryKey;autoIncrement:false"`
	Month int   `json:"month" gorm:"primaryKey;autoIncrement:false"`
	Day   int   `json:"day" gorm:"primaryKey;autoIncrement:false"`
	Count int64 `json:"count"`
}

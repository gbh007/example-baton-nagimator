package domain

import "strconv"

type Button struct {
	UserID int    `json:"-" gorm:"primaryKey;autoIncrement:false"`
	User   User   `json:"-" gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	Year   int    `json:"year" gorm:"primaryKey;autoIncrement:false"`
	Month  int    `json:"month" gorm:"primaryKey;autoIncrement:false"`
	Day    int    `json:"day" gorm:"primaryKey;autoIncrement:false"`
	Count  int64  `json:"count"`
	Text   string `json:"text" gorm:"-"`
}

func (b *Button) UpdateText() {
	b.Text = strconv.Itoa(b.Year) + "." +
		strconv.Itoa(b.Month) + "." +
		strconv.Itoa(b.Day)
}

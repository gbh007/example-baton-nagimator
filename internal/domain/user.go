package domain

type User struct {
	ID    int    `json:"year" gorm:"primaryKey"`
	Name  string `json:"name"`
	Token string `json:"token" gorm:"uniqueIndex"`
}

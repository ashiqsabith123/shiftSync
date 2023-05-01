package domain

type Admin struct {
	ID        int    `json:"admid" gorm:"primaryKey;unique"`
	Name      string `json:"admname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

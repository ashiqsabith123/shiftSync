package domain

type Admin struct {
	ID        int    `json:"admid" gorm:"primaryKey;unique"`
	Name      string `json:"admname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

type Duty struct {
	EmployeeID uint   `json:"empid"`
	Duty       Form   `gorm:"foreignKey:EmployeeID"`
	Duty_type  string `json:"type" gorm:"type:char(1)"`
	Status     string `json:"status" gorm:"type:char(1)"`
}

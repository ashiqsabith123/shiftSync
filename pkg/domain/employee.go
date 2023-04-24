package domain

type Employee_Signup struct {
	Signup_id int    `json:"id"`
	Full_name string `json:"fullname"`
	Email     string `json:"email"`
	Phone     int64  `json:"phone"`
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

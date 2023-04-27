package request

type LoginStruct struct {
	User_name string `json:"username"`
	Pass_word string `json:"password"`
}

type OTPStruct struct {
	Code string `json:"otp"`
}

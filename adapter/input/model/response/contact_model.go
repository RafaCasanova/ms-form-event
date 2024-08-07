package response

type ContactResponse struct {
	Name  string `validate:"min=3,max=70"`
	Age   uint8  `validate:"min=10"`
	Email string
}

type ResponseError struct {
	TypeErro string
	Title    string
	Detail   []string
}

type RecaptchaResponse struct {
	Success bool     `json:"success"`
	Errors  []string `json:"error-codes"`
}

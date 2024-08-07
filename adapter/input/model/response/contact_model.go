package response

type ContactResponse struct {
	Name  string `validate:"min=3,max=40,regexp=^[a-zA-Z]*$"`
	Age   uint8  `validate:"min=10"`
	Email string `validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
}

type ResponseError struct {
	TypeErro string
	Title    string
	Detail   []string
}

package convert

import (
	"challenger/adapter/input/model/request"
	"challenger/adapter/input/model/response"
	"fmt"
	"net/http"
	"net/mail"
	"strconv"

	"gopkg.in/validator.v2"
)

func ConvertContactRequestoToResponse(req request.ContactRequest) (res response.ContactResponse, err []string) {
	res = convertObjectContactReqResp(req)
	err = verifyContactRequest(res)
	return
}

func ConvertHttpRequestToRequestConect(r *http.Request) (request.ContactRequest, error) {
	var formContact request.ContactRequest

	formContact.Name = r.FormValue("name")
	formContact.Email = r.FormValue("email")
	age := r.FormValue("age")

	num, err := strconv.ParseUint(age, 10, 8)
	if err != nil {
		return formContact, err
	}
	formContact.Age = uint8(num)
	return formContact, nil
}
func convertObjectContactReqResp(req request.ContactRequest) response.ContactResponse {
	return response.ContactResponse{
		Name:  req.Name,
		Age:   req.Age,
		Email: req.Email,
	}
}

func verifyContactRequest(rep response.ContactResponse) []string {
	var errOuts []string
	_, err := mail.ParseAddress(rep.Email)
	if err != nil {
		errOuts = append(errOuts, fmt.Sprintf("%s (%s)\n", "Email", err.Error()))
	}

	if err := validator.Validate(rep); err != nil {
		errs := err.(validator.ErrorMap)
		for f, e := range errs {
			errOuts = append(errOuts, fmt.Sprintf("%s (%v)\n", f, e))
		}
	}
	if len(errOuts) > 0 {
		return errOuts
	}
	return nil
}

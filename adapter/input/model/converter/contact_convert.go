package convert

import (
	"challenger/adapter/input/model/request"
	"challenger/adapter/input/model/response"
	"fmt"

	"gopkg.in/validator.v2"
)

func ConvertContactRequestoToResponse(req request.ContactRequest) (res response.ContactResponse, err []string) {
	res = convertObjectContactReqResp(req)
	err = verifyContactRequest(res)
	return
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
	if err := validator.Validate(rep); err != nil {
		errs := err.(validator.ErrorMap)
		for f, e := range errs {
			errOuts = append(errOuts, fmt.Sprintf("\t - %s (%v)\n", f, e))
		}
		return errOuts
	}

	return nil
}

package convert

import (
	"challenger/adapter/input/model/response"
	"challenger/adapter/output/model"
)

func ConvertContactEntityToContactResponse(entity model.UserEntity) response.ContactResponse {
	return response.ContactResponse{
		Name:  entity.Name,
		Email: entity.Email,
		Age:   uint8(entity.Age),
	}
}

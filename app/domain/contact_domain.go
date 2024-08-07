package domain

import (
	"challenger/adapter/input/model/response"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ContactDomain struct {
	Id    primitive.ObjectID `bson:"_id,omitempty"`
	Email string             `bson:"email"`
	Name  string             `bson:"name"`
	Age   uint8              `bson:"age"`
}

func ConvertRequestToDomain(contact response.ContactResponse) ContactDomain {
	return ContactDomain{
		Id:    primitive.NewObjectID(),
		Email: contact.Email,
		Name:  contact.Name,
		Age:   contact.Age,
	}
}

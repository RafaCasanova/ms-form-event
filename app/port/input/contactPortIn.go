package input

import (
	"challenger/adapter/input/model/response"
	"challenger/app/domain"
)

type ContactDomainService interface {
	CreateContactServices(response.ContactResponse) (*domain.ContactDomain, error)
	FindContactByIDServices(id string) (*domain.ContactDomain, error)
	FindContactByEmailServices(email string) (*domain.ContactDomain, error)
}

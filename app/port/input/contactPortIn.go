package input

import "challenger/app/domain"

type ContactDomainService interface {
	CreateContactServices(domain.ContactDomain) (*domain.ContactDomain, error)
	FindContactByIDServices(id string) (*domain.ContactDomain, error)
	FindContactByEmailServices(email string) (*domain.ContactDomain, error)
}

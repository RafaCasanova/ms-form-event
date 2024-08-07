package output

import "challenger/app/domain"

type ContactPort interface {
	CreateContact(userDomain domain.ContactDomain) (*domain.ContactDomain, error)
	FindContactByEmail(email string) (*domain.ContactDomain, error)
	FindContactByID(id string) (*domain.ContactDomain, error)
}

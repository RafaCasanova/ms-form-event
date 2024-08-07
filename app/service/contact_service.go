package service

import (
	"challenger/app/domain"
	"challenger/app/port/input"
	"challenger/app/port/output"
)

func NewContactServoce(contactRepository output.ContactPort) input.ContactDomainService {
	return &contactService{
		contactRepository,
	}
}

type contactService struct {
	repository output.ContactPort
}

func (c *contactService) CreateContactServices(domain.ContactDomain) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *contactService) FindContactByEmailServices(email string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *contactService) FindContactByIDServices(id string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

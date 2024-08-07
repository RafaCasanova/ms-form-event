package repository

import (
	"challenger/app/domain"
	"challenger/app/port/output"
)

func NewContactRepository(database string) output.ContactPort {
	return &ContactRepository{
		database,
	}
}

type ContactRepository struct {
	databaseConnection string
}

func (c *ContactRepository) CreateContact(userDomain domain.ContactDomain) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *ContactRepository) FindContactByEmail(email string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *ContactRepository) FindContactByID(id string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

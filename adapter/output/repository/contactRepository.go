package repository

import (
	"challenger/app/domain"
	"challenger/app/port/output"
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewContactRepository(database *mongo.Database) output.ContactPort {

	return &ContactRepository{
		databaseConnection: database, Collection: "conects",
	}
}

type ContactRepository struct {
	Collection         string
	databaseConnection *mongo.Database
}

func (c *ContactRepository) CreateContact(userDomain domain.ContactDomain) (*domain.ContactDomain, error) {
	collection := c.databaseConnection.Collection(c.Collection)
	_, err := collection.InsertOne(context.Background(), userDomain)
	return &userDomain, err
}

func (c *ContactRepository) FindContactByEmail(email string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *ContactRepository) FindContactByID(id string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

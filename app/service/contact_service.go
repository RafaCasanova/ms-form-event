package service

import (
	"challenger/adapter/input/model/response"
	kafkaMs "challenger/app/config/kafka"
	"challenger/app/domain"
	"challenger/app/port/input"
	"challenger/app/port/output"
	"encoding/json"
	"os"

	"github.com/segmentio/kafka-go"
)

func NewContactServoce(contactRepository output.ContactPort) input.ContactDomainService {
	kafkaW := kafkaMs.GetKafkaWriter(os.Getenv("KAFKA_BROKER"), os.Getenv("KAFKA_EMAIL_TOPIC"))
	return &contactService{
		repository: contactRepository, kafkaWrite: kafkaW,
	}
}

type contactService struct {
	kafkaWrite *kafka.Writer
	repository output.ContactPort
}

func (c *contactService) CreateContactServices(contact response.ContactResponse) (*domain.ContactDomain, error) {
	contactDom := domain.ConvertRequestToDomain(contact)
	_, erro := c.repository.CreateContact(contactDom)
	if erro != nil {
		return &contactDom, erro
	}
	data, err := json.Marshal(contactDom)
	if err != nil {
		return &contactDom, erro
	}
	kafkaMs.SendMessage(c.kafkaWrite, contactDom.Id.String(), string(data))
	return &contactDom, nil
}

func (c *contactService) FindContactByEmailServices(email string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

func (c *contactService) FindContactByIDServices(id string) (*domain.ContactDomain, error) {
	panic("unimplemented")
}

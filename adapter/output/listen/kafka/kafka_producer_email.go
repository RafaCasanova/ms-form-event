package kafka_liste

import (
	"challenger/adapter/output/mail"
	"challenger/adapter/output/model"
	"challenger/adapter/output/model/convert"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/segmentio/kafka-go"
)

func InitConsumer() {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_BROKER")},
		GroupID:  "contact_consumer",
		Topic:    os.Getenv("KAFKA_EMAIL_TOPIC"),
		MaxBytes: 10e6,
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println(err.Error())
			return
		}
		var contact model.UserEntity
		err = json.Unmarshal(m.Value, &contact)
		if err != nil {
			fmt.Println("Error:", err)
			return
		}
		go func(contactEntity model.UserEntity) {
			cotactResp := convert.ConvertContactEntityToContactResponse(contactEntity)
			go mail.SendEmailCompany(cotactResp)
			go mail.SendEmailToUser(cotactResp)

		}(contact)
	}
}

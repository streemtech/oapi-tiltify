package main

import (
	"encoding/json"
	"fmt"
	"time"

	amqp091 "github.com/rabbitmq/amqp091-go"
	"github.com/streemtech/oapi-tiltify/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Donation struct {
	ID        int `gorm:"primarykey"`
	CreatedAt time.Time

	Amount  int64
	Donator string `gorm:"type:text"`
	Message string `gorm:"type:text"`

	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

func main() {

	var err error
	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=admin dbname=tiltify password=password sslmode=disable application_name=tiltify_loader"), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&Donation{})
	if err != nil {
		panic(err.Error())
	}

	// config := "amqp://admin:@192.168.100.101:30820"
	config := "amqp://guest:guest@localhost:5672"
	rabbit, err := amqp091.DialConfig(config, amqp091.Config{
		Heartbeat: time.Second * 10,
		Locale:    "en_US",
		// Vhost:     "qa",
		Vhost: "dev",

		Properties: amqp091.Table{
			"product":         "https://github.com/rabbitmq/amqp091-go",
			"version":         "β",
			"connection_name": "test-connection",
		},
	})
	if err != nil {
		panic(err)
	}

	ch, err := rabbit.Channel()
	if err != nil {
		panic(err)
	}

	err = ch.ExchangeDeclare("Chaos.Donation.Input", amqp091.ExchangeDirect, true, false, false, false, nil)
	if err != nil {
		fmt.Printf("Error declaring exchange testData: %s\n", err.Error())
	}

	min := 4499827
	max := 4501404
	donations := make([]Donation, 0)
	// start := time.Now()
	err = db.Model(&Donation{}).Where("id >= ?", min).Where("id <= ?", max).Order("created_at asc").Find(&donations).Error
	if err != nil {
		panic(err.Error())
	}

	//FEATURE REQUESTS
	//Add pinned options for cookie and disc only.
	//

	// fmt.Printf("loadTime: %+v: %s\n", donations[0:5], time.Since(start).String())

	doDonationsOnTime(donations, func(d Donation) {
		// fmt.Printf("%s: %s: $%0.2f: %s\n", d.CreatedAt, d.Donator, float64(d.Amount)/100.0, d.Message)
		fmt.Printf("%+v\n", d)

		f := float32(d.Amount)

		t := api.CampaignsIdDonations{
			Id:      &d.ID,
			Comment: &d.Message,
			Name:    &d.Donator,
			Amount:  &f,
		}
		var data []byte
		data, err = json.Marshal(t)
		if err != nil {
			fmt.Printf("Error marshaling: %s\n", err.Error())
			return
		}
		err = ch.Publish("Fanout.Tiltify.Campaign", "114159", false, false, amqp091.Publishing{
			ContentType:  "text/text",
			Body:         data,
			DeliveryMode: amqp091.Persistent,
			Timestamp:    time.Now(),
			Headers: amqp091.Table{
				"Campaign": "114159",
			},
		})
		if err != nil {
			fmt.Printf("Error publishing testData: %s\n", err.Error())
		}
	})

}

func doDonationsOnTime(d []Donation, callback func(Donation)) {

	var wait time.Duration

	for i, v := range d {

		next := d[i+1]
		wait = next.CreatedAt.Sub(v.CreatedAt)
		callback(v)
		fmt.Printf("Waiting %s\n", wait)
		time.Sleep(wait)
	}

}

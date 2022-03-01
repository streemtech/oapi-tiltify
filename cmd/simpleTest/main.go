package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http/httputil"
	"os"
	"time"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/rabbitmq/amqp091-go"
	"github.com/streemtech/oapi-tiltify/tiltifyApi"
)

var ch *amqp091.Channel

func main() {
	f, err := os.Open("./test.key")
	if err != nil {
		panic(err)
	}
	testKey, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	config := "amqp://admin:@192.168.100.101:30820"
	// config := "amqp://guest:guest@localhost:5672"
	rabbit, err := amqp091.DialConfig(config, amqp091.Config{
		Heartbeat: time.Second * 10,
		Locale:    "en_US",
		Vhost:     "qa",
		// Vhost: "dev",

		Properties: amqp091.Table{
			"product":         "https://github.com/rabbitmq/amqp091-go",
			"version":         "β",
			"connection_name": "test-connection",
		},
	})
	if err != nil {
		panic(err)
	}

	ch, err = rabbit.Channel()
	if err != nil {
		panic(err)
	}

	// db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=admin dbname=tiltify password=password sslmode=disable application_name=tiltify_loader"), &gorm.Config{
	// 	PrepareStmt: true,
	// })

	// if err != nil {
	// 	panic(err)
	// }

	// err = db.AutoMigrate(&Donation{})

	// if err != nil {
	// 	panic(err.Error())
	// }
	// pullDonations(string(testKey))
	// url := "https://tiltify.com/@supermcgamer/trgc-2021"
	getDonationsFromURL(string(testKey))
}

func getDonationsFromURL(key string) {
	sp, err := securityprovider.NewSecurityProviderBearerToken(key)
	if err != nil {
		panic(err)
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	if err != nil {
		panic(err)
	}
	total := 0
	divider := 500
	newTotal, err := getTotal(client, 157301)
	fmt.Printf("%f\n", newTotal)
	total = int(newTotal)

	total = total / divider
	total = total * divider

	fmt.Printf("%d\n", total)
	for {

		client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))
		if err != nil {
			fmt.Printf("ERROR ENCOUNTERED CREATING CLIENT: %s\n", err.Error())
			time.Sleep(time.Second * 1)
			continue
		}

		newTotal, err := getTotal(client, 157301)

		if err == nil {
			if int(newTotal) >= total+divider {
				total += divider
				fmt.Printf("New Total %d\n", total)
				go sendRush()
			}
		} else {

			fmt.Printf("ERROR ENCOUNTERED: %s\n", err.Error())
		}

		time.Sleep(time.Second * 5)

		// err := db.Clauses(clause.OnConflict{
		// 	UpdateAll: true,
		// }).Create(data).Error

	}
}

func getTotal(client *tiltifyApi.ClientWithResponses, campaign int) (total float32, err error) {

	resp, err := client.GetCampaignsIdWithResponse(context.Background(), campaign)
	if err != nil {
		return 0, fmt.Errorf("error from tiltify: Error making request: %s", err.Error())
	}

	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		d := *resp.JSON200.Data
		return *d.AmountRaised, nil
	}

	dump, err := httputil.DumpResponse(resp.HTTPResponse, true)
	if err != nil {
		return 0, fmt.Errorf("error from tiltify: Error getting response dump: %s", err.Error())
	}
	return 0, fmt.Errorf("received non 200 from tiltify: %s", dump)

}

var message = `{
    "TypeKey": "PARALLEL",
    "Actions": [
        {
            "TypeKey": "MEDIA_VIDEO",
            "Media": "https://api-qa.streem.tech/hostr/file/c70fc458-2a20-44a3-bef2-41c7346c1278/21183535-8718-4daa-bd50-f2e85361ec04"
        },
		{
			"TypeKey": "MEDIA_AUDIO",
			"Media": "https://api-qa.streem.tech/hostr/file/c70fc458-2a20-44a3-bef2-41c7346c1278/64a6fe5d-fa85-4ccc-9a76-b34f03b5b89e"
		}
    ]
}`

const FANFARE_BOARD_INPUT_EXCHANGE = "Fanfare.Boards.Input"

func sendRush() {

	ch.Publish(FANFARE_BOARD_INPUT_EXCHANGE, "", false, false, amqp091.Publishing{
		ContentType:  "text/text",
		Body:         []byte(message),
		DeliveryMode: amqp091.Persistent,
		Timestamp:    time.Now(),
		Headers: amqp091.Table{
			"table": "413f2ae4-1710-4f2d-a56f-e7f8c1658693",
		},
	})
}

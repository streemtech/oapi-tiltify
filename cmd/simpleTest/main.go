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
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var ch *amqp091.Channel
var db *gorm.DB

func main() {
	f, err := os.Open("./test.key")
	if err != nil {
		panic(err)
	}
	testKey, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	db, err = gorm.Open(postgres.Open("host=localhost port=5432 user=admin dbname=tiltify password=password sslmode=disable application_name=tiltify_loader"), &gorm.Config{
		PrepareStmt: true,
	})

	if err != nil {
		panic(err)
	}

	err = db.Exec("drop table if exists donations").Error
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&Donation{})
	if err != nil {
		panic(err)
	}
	// if err != nil {
	// 	panic(err.Error())
	// }
	// pullDonations(string(testKey))
	// url := "https://tiltify.com/@supermcgamer/trgc-2021"
	donos := getDonationsFromURL(string(testKey))

	message := ""
	name := ""
	amount := 0.0
	createdAt := time.Now()
	id := 0

	for _, v := range donos {
		if v.Amount != nil {
			amount = float64(*v.Amount)
		}
		if v.Comment != nil {
			message = *v.Comment
		}
		if v.Id != nil {
			id = *v.Id
		}
		if v.Name != nil {
			name = *v.Name
		}
		if v.CompletedAt != nil {
			createdAt = time.UnixMilli(int64(*v.CompletedAt))
		}

		err = db.Create(&Donation{
			Message: message,
			Donator: name,
			Amount:  amount,
			Time:    createdAt,
			ID:      id,
		}).Error

		if err != nil {
			fmt.Printf("error creating donation %+v\n", v)
		}
	}
}

func getDonationsFromURL(key string) []tiltifyApi.CampaignsIdDonations {
	sp, err := securityprovider.NewSecurityProviderBearerToken(key)

	if err != nil {
		panic(err)
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	if err != nil {
		panic(err)
	}

	prev := -1
	donations, links, err := getTotal(client, 157301, prev)

	allDonations := make([]tiltifyApi.CampaignsIdDonations, 0)
	for {

		//load in the donations.
		if err == nil {
			if len(donations) <= 0 {
				break
			}
			allDonations = append(allDonations, donations...)
			fmt.Printf("Added %d donations. total: %d\n", len(donations), len(allDonations))

		} else {
			fmt.Printf("ERROR ENCOUNTERED: %s\n", err.Error())
		}

		//update next.
		prev = tiltifyApi.ParseLinks(links.Prev).Before

		//sleep.
		time.Sleep(time.Millisecond * 100)

		//create new client and load in next donations.
		client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))
		if err != nil {
			fmt.Printf("ERROR ENCOUNTERED CREATING CLIENT: %s\n", err.Error())
			time.Sleep(time.Second * 1)
			continue
		}
		donations, links, err = getTotal(client, 157301, prev)

		// err := db.Clauses(clause.OnConflict{
		// 	UpdateAll: true,
		// }).Create(data).Error

	}
	fmt.Printf("Total Donations: %d\n", len(allDonations))

	return allDonations
}

func getTotal(client *tiltifyApi.ClientWithResponses, campaign int, Before int) (donations []tiltifyApi.CampaignsIdDonations, links *tiltifyApi.Pagination, err error) {

	count := 100
	params := &tiltifyApi.GetCampaignsIdDonationsParams{
		Count: &count,
	}

	if Before > 0 {
		params.Before = &Before

	}
	resp, err := client.GetCampaignsIdDonationsWithResponse(context.Background(), campaign, params)
	if err != nil {
		return []tiltifyApi.CampaignsIdDonations{}, nil, fmt.Errorf("error from tiltify: Error making request: %s", err.Error())
	}

	// fmt.Printf("Response Dump: %s\n", dump)
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		donations := *resp.JSON200.Data
		return donations, resp.JSON200.Links, nil
	}

	dump, err := httputil.DumpResponse(resp.HTTPResponse, true)
	if err != nil {
		return []tiltifyApi.CampaignsIdDonations{}, nil, fmt.Errorf("error from tiltify: Error getting response dump: %s", err.Error())
	}
	return []tiltifyApi.CampaignsIdDonations{}, nil, fmt.Errorf("received non 200 from tiltify: %s", dump)

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

type Donation struct {
	Message string
	Donator string
	Amount  float64
	Time    time.Time
	ID      int
}

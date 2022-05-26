package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http/httputil"
	"os"
	"time"

	_ "github.com/deepmap/oapi-codegen/pkg/codegen"
	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/streemtech/oapi-tiltify/api"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

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

func getDonationsFromURL(key string) []api.CampaignsIdDonations {
	sp, err := securityprovider.NewSecurityProviderBearerToken(key)

	if err != nil {
		panic(err)
	}

	client, err := api.NewClientWithResponses("https://tiltify.com/api/v3", api.WithRequestEditorFn(sp.Intercept))

	if err != nil {
		panic(err)
	}

	prev := -1
	donations, links, err := getTotal(client, 157301, prev)

	allDonations := make([]api.CampaignsIdDonations, 0)
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
		prev = api.ParseLinks(links.Prev).Before

		//sleep.
		time.Sleep(time.Millisecond * 100)

		//create new client and load in next donations.
		client, err = api.NewClientWithResponses("https://tiltify.com/api/v3", api.WithRequestEditorFn(sp.Intercept))
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

func getTotal(client *api.ClientWithResponses, campaign int, Before int) (donations []api.CampaignsIdDonations, links *api.Pagination, err error) {

	count := 100
	params := &api.GetCampaignsIdDonationsParams{
		Count: &count,
	}

	if Before > 0 {
		params.Before = &Before

	}
	resp, err := client.GetCampaignsIdDonationsWithResponse(context.Background(), campaign, params)
	if err != nil {
		return []api.CampaignsIdDonations{}, nil, fmt.Errorf("error from tiltify: Error making request: %s", err.Error())
	}

	// fmt.Printf("Response Dump: %s\n", dump)
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		donations = *resp.JSON200.Data
		return donations, resp.JSON200.Links, nil
	}

	dump, err := httputil.DumpResponse(resp.HTTPResponse, true)
	if err != nil {
		return []api.CampaignsIdDonations{}, nil, fmt.Errorf("error from tiltify: Error getting response dump: %s", err.Error())
	}
	return []api.CampaignsIdDonations{}, nil, fmt.Errorf("received non 200 from tiltify: %s", dump)

}

// func sendRush() {

// 	var ch *amqp091.Channel
// 	ch.Publish("Fanfare.Boards.Input", "", false, false, amqp091.Publishing{
// 		ContentType: "text/text",
// 		Body: []byte(`{
// 			"TypeKey": "PARALLEL",
// 			"Actions": [
// 				{
// 					"TypeKey": "MEDIA_VIDEO",
// 					"Media": "https://api-qa.streem.tech/hostr/file/c70fc458-2a20-44a3-bef2-41c7346c1278/21183535-8718-4daa-bd50-f2e85361ec04"
// 				},
// 				{
// 					"TypeKey": "MEDIA_AUDIO",
// 					"Media": "https://api-qa.streem.tech/hostr/file/c70fc458-2a20-44a3-bef2-41c7346c1278/64a6fe5d-fa85-4ccc-9a76-b34f03b5b89e"
// 				}
// 			]
// 		}`),
// 		DeliveryMode: amqp091.Persistent,
// 		Timestamp:    time.Now(),
// 		Headers: amqp091.Table{
// 			"table": "413f2ae4-1710-4f2d-a56f-e7f8c1658693",
// 		},
// 	})
// }

type Donation struct {
	Message string
	Donator string
	Amount  float64
	Time    time.Time
	ID      int
}

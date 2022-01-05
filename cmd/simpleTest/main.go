package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/streemtech/oapi-tiltify/tiltifyApi"
)

func main() {
	f, err := os.Open("./test.key")
	if err != nil {
		panic(err)
	}
	testKey, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	pullDataFromTiltify(string(testKey))
}

func pullDataFromTiltify(testKey string) {
	sp, err := securityprovider.NewSecurityProviderBearerToken(testKey)
	if err != nil {
		panic(err)
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	if err != nil {
		panic(err)
	}
	i := 100
	resp, err := client.GetCampaignsIdDonationsWithResponse(context.Background(), 147920, &tiltifyApi.GetCampaignsIdDonationsParams{
		Count: &i,
	})
	if err != nil {
		panic(err)
	}
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		for _, v := range *resp.JSON200.Data {
			if v.Name != nil && v.Comment != nil {
				fmt.Printf("%50s: %s\n", *v.Name, *v.Comment)
			}
		}
	} else {
		fmt.Printf("data null: %+v\n", resp)
	}
}

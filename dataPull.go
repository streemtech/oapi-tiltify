package main

import (
	"context"
	"fmt"

	"github.com/deepmap/oapi-codegen/pkg/securityprovider"
	"github.com/streemtech/oapi-tiltify/tiltifyApi"
)

var testKey = "1dad548e3b94911d751c1b11d9c3be942132ebae308819e1b5835248c52b52eb"

//1dad548e3b94911d751c1b11d9c3be942132ebae308819e1b5835248c52b52eb
//TODO0 validate this key has timed out.
func pullDataFromTiltify() {
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
	fmt.Printf("%+v\n", resp.JSON200.Data)
}

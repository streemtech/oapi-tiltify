package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/k0kubun/pp/v3"
	"github.com/streemtech/oapi-tiltify/api"
)

var Token = "06d77c25dbeb2a0cd831547f7960bfa47c7b459fd53b6a07dd9c65f4db4d9b00"

func main() {
	cmd, err := api.NewClientWithResponses("https://v5api.tiltify.com", true, api.WithRequestEditorFn(ref))
	if err != nil {
		log.Fatal(err.Error())
	}

	resp, err := cmd.V5ApiWebPublicTeamCampaignDonationControllerIndexWithResponse(context.TODO(), "7d47a0ee-49cf-4ed2-9996-ec77c1f847bd", &api.V5ApiWebPublicTeamCampaignDonationControllerIndexParams{
		Limit: aws.Int(10),
	})
	if err != nil {
		log.Fatal(err.Error())
	}

	pp.Println(resp)
}

func ref(ctx context.Context, req *http.Request) error {
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", Token))
	return nil
}

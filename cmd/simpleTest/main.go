package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/url"
	"os"
	"strings"

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

	// pullDonations(string(testKey))
	url := "https://tiltify.com/@supermcgamer/trgc-2021"
	getDonationsFromURL(string(testKey), url)
}

func getDonationsFromURL(key, url string) {

	user, campaignSlug, err := convertUrlToUserCampaignSlug(url)
	if err != nil {
		panic(err)
	}
	userID, err := getUserIDFromUsername(key, user)
	if err != nil {
		panic(err)
	}
	campaignID, err := getCampaignFromUserIDAndSlug(key, campaignSlug, userID)
	if err != nil {
		panic(err)
	}
	getDonations(key, campaignID)
}

func convertUrlToUserCampaignSlug(in string) (username, slug string, err error) {

	u, err := url.Parse(in)
	if err != nil {
		return "", "", err
	}
	p := u.Path
	spl := strings.Split(p, "/")
	if len(spl) != 3 {
		return "", "", fmt.Errorf("unable to split path %s. Got %d components", p, len(spl))
	}
	username = strings.ReplaceAll(spl[1], "@", "")
	slug = spl[2]
	return username, slug, nil
}

func getUserIDFromUsername(testKey, username string) (id int, err error) {
	sp, err := securityprovider.NewSecurityProviderBearerToken(testKey)
	if err != nil {
		return 0, err
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	resp, err := client.GetUsersSlugWithResponse(context.Background(), username)
	if err != nil {
		return 0, err
	}
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		d := resp.JSON200.Data
		if d.Id != nil {
			return *d.Id, nil
		}
	}
	return 0, fmt.Errorf("data not 200")
}

func getCampaignFromUserIDAndSlug(key, campaignSlug string, userID int) (campaignID int, err error) {
	sp, err := securityprovider.NewSecurityProviderBearerToken(key)
	if err != nil {
		return 0, err
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	c := 100
	resp, err := client.GetUsersIdCampaignsWithResponse(context.Background(), userID, &tiltifyApi.GetUsersIdCampaignsParams{
		Count: &c,
	})
	if err != nil {
		return 0, err
	}

	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		for _, v := range *resp.JSON200.Data {
			if *v.Slug == campaignSlug {
				return *v.Id, nil
			}
		}
		return 0, fmt.Errorf("campaign %s not found in data", campaignSlug)
	}
	return 0, fmt.Errorf("data not 200")
}

func getDonations(testKey string, campaign int) {
	sp, err := securityprovider.NewSecurityProviderBearerToken(testKey)
	if err != nil {
		panic(err)
	}

	client, err := tiltifyApi.NewClientWithResponses("https://tiltify.com/api/v3", tiltifyApi.WithRequestEditorFn(sp.Intercept))

	if err != nil {
		panic(err)
	}
	i := 100
	resp, err := client.GetCampaignsIdDonationsWithResponse(context.Background(), campaign, &tiltifyApi.GetCampaignsIdDonationsParams{
		Count: &i,
	})
	if err != nil {
		panic(err)
	}
	if resp.JSON200 != nil && resp.JSON200.Data != nil {
		d := *resp.JSON200.Data
		for _, v := range d {
			if v.Name != nil && v.Comment != nil {
				id := -1
				if v.RewardId != nil {
					id = *v.RewardId
				}
				fmt.Printf("%50s: $%07.2f | %d | %d\n", *v.Name, *v.Amount, *v.Id, id)
			}
		}
		fmt.Printf("%+v | %+v | %+v\n", tiltifyApi.ParseLinks(resp.JSON200.Links.Prev), tiltifyApi.ParseLinks(resp.JSON200.Links.Self), tiltifyApi.ParseLinks(resp.JSON200.Links.Next))
	} else {
		fmt.Printf("data null: %+v\n", resp)
	}
}

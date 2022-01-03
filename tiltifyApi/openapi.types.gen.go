// Package tiltifyApi provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.0 DO NOT EDIT.
package tiltifyApi

// error response body. Only included in Post and Patch requests.
type Error struct {
	// An object describing the error that occurred
	Error *struct {
		Detial *string `json:"detial,omitempty"`
		Title  *string `json:"title,omitempty"`
	} `json:"error,omitempty"`

	// An object pointing to fields that are incorrect in a submitted model
	Errors *map[string]interface{} `json:"errors,omitempty"`

	// This is the HTTP status code that is also sent with the request
	Meta *Meta `json:"meta,omitempty"`
}

// Image defines model for Image.
type Image struct {
	Alt    *string `json:"alt,omitempty"`
	Height *int    `json:"height,omitempty"`
	Src    *string `json:"src,omitempty"`
	Width  *int    `json:"width,omitempty"`
}

// Livestream defines model for Livestream.
type Livestream struct {
	Channel *string `json:"channel,omitempty"`
	Type    *string `json:"type,omitempty"`
}

// This is the HTTP status code that is also sent with the request
type Meta struct {
	Status *int `json:"status,omitempty"`
}

// We use cursor based pagination for our donations and this information is embeded in the response under the links key. You will find a prev and next link that point to the next pages of the paginated response. You may submit an optional count field of up to 100.
type Pagination struct {
	First *string `json:"first,omitempty"`
	Last  *string `json:"last,omitempty"`
	Next  string  `json:"next"`
	Prev  string  `json:"prev"`
	Self  string  `json:"self"`
}

// Social defines model for Social.
type Social struct {
	Discord   *string `json:"discord,omitempty"`
	Facebook  *string `json:"facebook,omitempty"`
	Instagram *string `json:"instagram,omitempty"`
	Mixer     *string `json:"mixer,omitempty"`
	Twitch    *string `json:"twitch,omitempty"`
	Twitter   *string `json:"twitter,omitempty"`
	Website   *string `json:"website,omitempty"`
	Youtube   *string `json:"youtube,omitempty"`
}

// error response body. Only included in Post and Patch requests.
type BadRequest Error

// error response body. Only included in Post and Patch requests.
type Forbidden Error

// error response body. Only included in Post and Patch requests.
type NotAuthorized Error

// error response body. Only included in Post and Patch requests.
type NotFound Error

// error response body. Only included in Post and Patch requests.
type ServerError Error

// error response body. Only included in Post and Patch requests.
type Unprocessable Error

// GetCampaignsIdChallengesParams defines parameters for GetCampaignsIdChallenges.
type GetCampaignsIdChallengesParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCampaignsIdDonationsParams defines parameters for GetCampaignsIdDonations.
type GetCampaignsIdDonationsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCampaignsIdPollsParams defines parameters for GetCampaignsIdPolls.
type GetCampaignsIdPollsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCampaignsIdRewardsParams defines parameters for GetCampaignsIdRewards.
type GetCampaignsIdRewardsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCampaignsIdScheduleParams defines parameters for GetCampaignsIdSchedule.
type GetCampaignsIdScheduleParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCampaignsIdSupportingCampaignsParams defines parameters for GetCampaignsIdSupportingCampaigns.
type GetCampaignsIdSupportingCampaignsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCausesIdCampaignsParams defines parameters for GetCausesIdCampaigns.
type GetCausesIdCampaignsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCausesIdFundraisingEventsParams defines parameters for GetCausesIdFundraisingEvents.
type GetCausesIdFundraisingEventsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetCausesIdVisibilityOptionsParams defines parameters for GetCausesIdVisibilityOptions.
type GetCausesIdVisibilityOptionsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsParams defines parameters for GetFundraisingEvents.
type GetFundraisingEventsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdCampaignsParams defines parameters for GetFundraisingEventsIdCampaigns.
type GetFundraisingEventsIdCampaignsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdIncentivesParams defines parameters for GetFundraisingEventsIdIncentives.
type GetFundraisingEventsIdIncentivesParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdLeaderboardsParams defines parameters for GetFundraisingEventsIdLeaderboards.
type GetFundraisingEventsIdLeaderboardsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdRegistrationFieldsParams defines parameters for GetFundraisingEventsIdRegistrationFields.
type GetFundraisingEventsIdRegistrationFieldsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdRegistrationsParams defines parameters for GetFundraisingEventsIdRegistrations.
type GetFundraisingEventsIdRegistrationsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdScheduleParams defines parameters for GetFundraisingEventsIdSchedule.
type GetFundraisingEventsIdScheduleParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetFundraisingEventsIdVisibilityOptionsParams defines parameters for GetFundraisingEventsIdVisibilityOptions.
type GetFundraisingEventsIdVisibilityOptionsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetTeamsParams defines parameters for GetTeams.
type GetTeamsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetTeamsIdCampaignsParams defines parameters for GetTeamsIdCampaigns.
type GetTeamsIdCampaignsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetUsersIdCampaignsParams defines parameters for GetUsersIdCampaigns.
type GetUsersIdCampaignsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetUsersIdOwnedTeamsParams defines parameters for GetUsersIdOwnedTeams.
type GetUsersIdOwnedTeamsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}

// GetUsersIdTeamsParams defines parameters for GetUsersIdTeams.
type GetUsersIdTeamsParams struct {
	// This is the amount of results to return for the page. This number must be between 1 and 100. Used in Pagination.
	Count *int `json:"count,omitempty"`

	// Used in Pagination.
	Before *int `json:"before,omitempty"`

	// Used in Pagination.
	After *int `json:"after,omitempty"`
}
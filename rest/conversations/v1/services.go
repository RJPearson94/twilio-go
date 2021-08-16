/*
 * Twilio - Conversations
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.20.0
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Optional parameters for the method 'CreateService'
type CreateServiceParams struct {
	// The human-readable name of this service, limited to 256 characters. Optional.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *CreateServiceParams) SetFriendlyName(FriendlyName string) *CreateServiceParams {
	params.FriendlyName = &FriendlyName
	return params
}

// Create a new conversation service on your account
func (c *ApiService) CreateService(params *CreateServiceParams) (*ConversationsV1Service, error) {
	path := "/v1/Services"

	data := url.Values{}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Remove a conversation service with all its nested resources from your account
func (c *ApiService) DeleteService(Sid string) error {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Delete(c.baseURL+path, data, headers)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	return nil
}

// Fetch a conversation service from your account
func (c *ApiService) FetchService(Sid string) (*ConversationsV1Service, error) {
	path := "/v1/Services/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ConversationsV1Service{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListService'
type ListServiceParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListServiceParams) SetPageSize(PageSize int) *ListServiceParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListServiceParams) SetLimit(Limit int) *ListServiceParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of Service records from the API. Request is executed immediately.
func (c *ApiService) PageService(params *ListServiceParams, pageToken string, pageNumber string) (*ListServiceResponse, error) {
	path := "/v1/Services"

	data := url.Values{}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}
	headers := make(map[string]interface{})

	if pageToken != "" {
		data.Set("PageToken", pageToken)
	}
	if pageToken != "" {
		data.Set("Page", pageNumber)
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists Service records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListService(params *ListServiceParams) ([]ConversationsV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ConversationsV1Service

	for response != nil {
		records = append(records, response.Services...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListServiceResponse)
	}

	return records, err
}

// Streams Service records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamService(params *ListServiceParams) (chan ConversationsV1Service, error) {
	if params == nil {
		params = &ListServiceParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageService(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ConversationsV1Service, 1)

	go func() {
		for response != nil {
			for item := range response.Services {
				channel <- response.Services[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListServiceResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListServiceResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListServiceResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListServiceResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

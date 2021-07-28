/*
 * Twilio - Pricing
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
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

func (c *ApiService) FetchPhoneNumberCountry(IsoCountry string) (*PricingV1PhoneNumberPhoneNumberCountryInstance, error) {
	path := "/v1/PhoneNumbers/Countries/{IsoCountry}"
	path = strings.Replace(path, "{"+"IsoCountry"+"}", IsoCountry, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &PricingV1PhoneNumberPhoneNumberCountryInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListPhoneNumberCountry'
type ListPhoneNumberCountryParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListPhoneNumberCountryParams) SetPageSize(PageSize int) *ListPhoneNumberCountryParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListPhoneNumberCountryParams) SetLimit(Limit int) *ListPhoneNumberCountryParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of PhoneNumberCountry records from the API. Request is executed immediately.
func (c *ApiService) PagePhoneNumberCountry(params *ListPhoneNumberCountryParams, pageToken string, pageNumber string) (*ListPhoneNumberCountryResponse, error) {
	path := "/v1/PhoneNumbers/Countries"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

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

	ps := &ListPhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists PhoneNumberCountry records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListPhoneNumberCountry(params *ListPhoneNumberCountryParams) ([]PricingV1PhoneNumberPhoneNumberCountry, error) {
	if params == nil {
		params = &ListPhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PagePhoneNumberCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []PricingV1PhoneNumberPhoneNumberCountry

	for response != nil {
		records = append(records, response.Countries...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListPhoneNumberCountryResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListPhoneNumberCountryResponse)
	}

	return records, err
}

// Streams PhoneNumberCountry records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamPhoneNumberCountry(params *ListPhoneNumberCountryParams) (chan PricingV1PhoneNumberPhoneNumberCountry, error) {
	if params == nil {
		params = &ListPhoneNumberCountryParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PagePhoneNumberCountry(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan PricingV1PhoneNumberPhoneNumberCountry, 1)

	go func() {
		for response != nil {
			for item := range response.Countries {
				channel <- response.Countries[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListPhoneNumberCountryResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListPhoneNumberCountryResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListPhoneNumberCountryResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListPhoneNumberCountryResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

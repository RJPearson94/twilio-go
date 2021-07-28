/*
 * Twilio - Messaging
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

// Optional parameters for the method 'CreateShortCode'
type CreateShortCodeParams struct {
	// The SID of the ShortCode resource being added to the Service.
	ShortCodeSid *string `json:"ShortCodeSid,omitempty"`
}

func (params *CreateShortCodeParams) SetShortCodeSid(ShortCodeSid string) *CreateShortCodeParams {
	params.ShortCodeSid = &ShortCodeSid
	return params
}

func (c *ApiService) CreateShortCode(ServiceSid string, params *CreateShortCodeParams) (*MessagingV1ServiceShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ShortCodeSid != nil {
		data.Set("ShortCodeSid", *params.ShortCodeSid)
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteShortCode(ServiceSid string, Sid string) error {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
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

func (c *ApiService) FetchShortCode(ServiceSid string, Sid string) (*MessagingV1ServiceShortCode, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes/{Sid}"
	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MessagingV1ServiceShortCode{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListShortCode'
type ListShortCodeParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListShortCodeParams) SetPageSize(PageSize int) *ListShortCodeParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListShortCodeParams) SetLimit(Limit int) *ListShortCodeParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of ShortCode records from the API. Request is executed immediately.
func (c *ApiService) PageShortCode(ServiceSid string, params *ListShortCodeParams, pageToken string, pageNumber string) (*ListShortCodeResponse, error) {
	path := "/v1/Services/{ServiceSid}/ShortCodes"

	path = strings.Replace(path, "{"+"ServiceSid"+"}", ServiceSid, -1)

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

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists ShortCode records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListShortCode(ServiceSid string, params *ListShortCodeParams) ([]MessagingV1ServiceShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []MessagingV1ServiceShortCode

	for response != nil {
		records = append(records, response.ShortCodes...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListShortCodeResponse)
	}

	return records, err
}

// Streams ShortCode records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamShortCode(ServiceSid string, params *ListShortCodeParams) (chan MessagingV1ServiceShortCode, error) {
	if params == nil {
		params = &ListShortCodeParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageShortCode(ServiceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan MessagingV1ServiceShortCode, 1)

	go func() {
		for response != nil {
			for item := range response.ShortCodes {
				channel <- response.ShortCodes[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListShortCodeResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListShortCodeResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListShortCodeResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListShortCodeResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

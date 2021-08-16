/*
 * Twilio - Voice
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

// Optional parameters for the method 'CreateIpRecord'
type CreateIpRecordParams struct {
	// An integer representing the length of the [CIDR](https://tools.ietf.org/html/rfc4632) prefix to use with this IP address. By default the entire IP address is used, which for IPv4 is value 32.
	CidrPrefixLength *int `json:"CidrPrefixLength,omitempty"`
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// An IP address in dotted decimal notation, IPv4 only.
	IpAddress *string `json:"IpAddress,omitempty"`
}

func (params *CreateIpRecordParams) SetCidrPrefixLength(CidrPrefixLength int) *CreateIpRecordParams {
	params.CidrPrefixLength = &CidrPrefixLength
	return params
}
func (params *CreateIpRecordParams) SetFriendlyName(FriendlyName string) *CreateIpRecordParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateIpRecordParams) SetIpAddress(IpAddress string) *CreateIpRecordParams {
	params.IpAddress = &IpAddress
	return params
}

func (c *ApiService) CreateIpRecord(params *CreateIpRecordParams) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords"

	data := url.Values{}
	if params != nil && params.CidrPrefixLength != nil {
		data.Set("CidrPrefixLength", fmt.Sprint(*params.CidrPrefixLength))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.IpAddress != nil {
		data.Set("IpAddress", *params.IpAddress)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *ApiService) DeleteIpRecord(Sid string) error {
	path := "/v1/IpRecords/{Sid}"
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

func (c *ApiService) FetchIpRecord(Sid string) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIpRecord'
type ListIpRecordParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIpRecordParams) SetPageSize(PageSize int) *ListIpRecordParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIpRecordParams) SetLimit(Limit int) *ListIpRecordParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IpRecord records from the API. Request is executed immediately.
func (c *ApiService) PageIpRecord(params *ListIpRecordParams, pageToken string, pageNumber string) (*ListIpRecordResponse, error) {
	path := "/v1/IpRecords"

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

	ps := &ListIpRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IpRecord records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIpRecord(params *ListIpRecordParams) ([]VoiceV1IpRecord, error) {
	if params == nil {
		params = &ListIpRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIpRecord(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VoiceV1IpRecord

	for response != nil {
		records = append(records, response.IpRecords...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIpRecordResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListIpRecordResponse)
	}

	return records, err
}

// Streams IpRecord records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIpRecord(params *ListIpRecordParams) (chan VoiceV1IpRecord, error) {
	if params == nil {
		params = &ListIpRecordParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIpRecord(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VoiceV1IpRecord, 1)

	go func() {
		for response != nil {
			for item := range response.IpRecords {
				channel <- response.IpRecords[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIpRecordResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListIpRecordResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListIpRecordResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIpRecordResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateIpRecord'
type UpdateIpRecordParams struct {
	// A descriptive string that you create to describe the resource. It is not unique and can be up to 255 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
}

func (params *UpdateIpRecordParams) SetFriendlyName(FriendlyName string) *UpdateIpRecordParams {
	params.FriendlyName = &FriendlyName
	return params
}

func (c *ApiService) UpdateIpRecord(Sid string, params *UpdateIpRecordParams) (*VoiceV1IpRecord, error) {
	path := "/v1/IpRecords/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

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

	ps := &VoiceV1IpRecord{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

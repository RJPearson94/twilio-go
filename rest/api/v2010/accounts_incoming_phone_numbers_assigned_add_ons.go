/*
 * Twilio - Api
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

// Optional parameters for the method 'CreateIncomingPhoneNumberAssignedAddOn'
type CreateIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID that identifies the Add-on installation.
	InstalledAddOnSid *string `json:"InstalledAddOnSid,omitempty"`
}

func (params *CreateIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *CreateIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateIncomingPhoneNumberAssignedAddOnParams) SetInstalledAddOnSid(InstalledAddOnSid string) *CreateIncomingPhoneNumberAssignedAddOnParams {
	params.InstalledAddOnSid = &InstalledAddOnSid
	return params
}

// Assign an Add-on installation to the Number specified.
func (c *ApiService) CreateIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *CreateIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)

	data := url.Values{}
	if params != nil && params.InstalledAddOnSid != nil {
		data.Set("InstalledAddOnSid", *params.InstalledAddOnSid)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberAssignedAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteIncomingPhoneNumberAssignedAddOn'
type DeleteIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *DeleteIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Remove the assignment of an Add-on installation from the Number specified.
func (c *ApiService) DeleteIncomingPhoneNumberAssignedAddOn(ResourceSid string, Sid string, params *DeleteIncomingPhoneNumberAssignedAddOnParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
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

// Optional parameters for the method 'FetchIncomingPhoneNumberAssignedAddOn'
type FetchIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *FetchIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of an Add-on installation currently assigned to this Number.
func (c *ApiService) FetchIncomingPhoneNumberAssignedAddOn(ResourceSid string, Sid string, params *FetchIncomingPhoneNumberAssignedAddOnParams) (*ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010IncomingPhoneNumberAssignedAddOn{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListIncomingPhoneNumberAssignedAddOn'
type ListIncomingPhoneNumberAssignedAddOnParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetPathAccountSid(PathAccountSid string) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetPageSize(PageSize int) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListIncomingPhoneNumberAssignedAddOnParams) SetLimit(Limit int) *ListIncomingPhoneNumberAssignedAddOnParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of IncomingPhoneNumberAssignedAddOn records from the API. Request is executed immediately.
func (c *ApiService) PageIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams, pageToken string, pageNumber string) (*ListIncomingPhoneNumberAssignedAddOnResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/IncomingPhoneNumbers/{ResourceSid}/AssignedAddOns.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"ResourceSid"+"}", ResourceSid, -1)

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

	ps := &ListIncomingPhoneNumberAssignedAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists IncomingPhoneNumberAssignedAddOn records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) ([]ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	if params == nil {
		params = &ListIncomingPhoneNumberAssignedAddOnParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIncomingPhoneNumberAssignedAddOn(ResourceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010IncomingPhoneNumberAssignedAddOn

	for response != nil {
		records = append(records, response.AssignedAddOns...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIncomingPhoneNumberAssignedAddOnResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListIncomingPhoneNumberAssignedAddOnResponse)
	}

	return records, err
}

// Streams IncomingPhoneNumberAssignedAddOn records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamIncomingPhoneNumberAssignedAddOn(ResourceSid string, params *ListIncomingPhoneNumberAssignedAddOnParams) (chan ApiV2010IncomingPhoneNumberAssignedAddOn, error) {
	if params == nil {
		params = &ListIncomingPhoneNumberAssignedAddOnParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageIncomingPhoneNumberAssignedAddOn(ResourceSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010IncomingPhoneNumberAssignedAddOn, 1)

	go func() {
		for response != nil {
			for item := range response.AssignedAddOns {
				channel <- response.AssignedAddOns[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListIncomingPhoneNumberAssignedAddOnResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListIncomingPhoneNumberAssignedAddOnResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListIncomingPhoneNumberAssignedAddOnResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListIncomingPhoneNumberAssignedAddOnResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

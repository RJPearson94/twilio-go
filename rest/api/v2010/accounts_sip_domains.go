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

// Optional parameters for the method 'CreateSipDomain'
type CreateSipDomainParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that will create the resource.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
	ByocTrunkSid *string `json:"ByocTrunkSid,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`.
	DomainName *string `json:"DomainName,omitempty"`
	// Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
	EmergencyCallerSid *string `json:"EmergencyCallerSid,omitempty"`
	// Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
	EmergencyCallingEnabled *bool `json:"EmergencyCallingEnabled,omitempty"`
	// A descriptive string that you created to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
	Secure *bool `json:"Secure,omitempty"`
	// Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not.
	SipRegistration *bool `json:"SipRegistration,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML from `voice_url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`. Can be: `GET` or `POST`.
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`.
	VoiceStatusCallbackMethod *string `json:"VoiceStatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	VoiceStatusCallbackUrl *string `json:"VoiceStatusCallbackUrl,omitempty"`
	// The URL we should when the domain receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *CreateSipDomainParams) SetPathAccountSid(PathAccountSid string) *CreateSipDomainParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *CreateSipDomainParams) SetByocTrunkSid(ByocTrunkSid string) *CreateSipDomainParams {
	params.ByocTrunkSid = &ByocTrunkSid
	return params
}
func (params *CreateSipDomainParams) SetDomainName(DomainName string) *CreateSipDomainParams {
	params.DomainName = &DomainName
	return params
}
func (params *CreateSipDomainParams) SetEmergencyCallerSid(EmergencyCallerSid string) *CreateSipDomainParams {
	params.EmergencyCallerSid = &EmergencyCallerSid
	return params
}
func (params *CreateSipDomainParams) SetEmergencyCallingEnabled(EmergencyCallingEnabled bool) *CreateSipDomainParams {
	params.EmergencyCallingEnabled = &EmergencyCallingEnabled
	return params
}
func (params *CreateSipDomainParams) SetFriendlyName(FriendlyName string) *CreateSipDomainParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *CreateSipDomainParams) SetSecure(Secure bool) *CreateSipDomainParams {
	params.Secure = &Secure
	return params
}
func (params *CreateSipDomainParams) SetSipRegistration(SipRegistration bool) *CreateSipDomainParams {
	params.SipRegistration = &SipRegistration
	return params
}
func (params *CreateSipDomainParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *CreateSipDomainParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *CreateSipDomainParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *CreateSipDomainParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *CreateSipDomainParams) SetVoiceMethod(VoiceMethod string) *CreateSipDomainParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *CreateSipDomainParams) SetVoiceStatusCallbackMethod(VoiceStatusCallbackMethod string) *CreateSipDomainParams {
	params.VoiceStatusCallbackMethod = &VoiceStatusCallbackMethod
	return params
}
func (params *CreateSipDomainParams) SetVoiceStatusCallbackUrl(VoiceStatusCallbackUrl string) *CreateSipDomainParams {
	params.VoiceStatusCallbackUrl = &VoiceStatusCallbackUrl
	return params
}
func (params *CreateSipDomainParams) SetVoiceUrl(VoiceUrl string) *CreateSipDomainParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Create a new Domain
func (c *ApiService) CreateSipDomain(params *CreateSipDomainParams) (*ApiV2010SipDomain, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

	data := url.Values{}
	if params != nil && params.ByocTrunkSid != nil {
		data.Set("ByocTrunkSid", *params.ByocTrunkSid)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.EmergencyCallerSid != nil {
		data.Set("EmergencyCallerSid", *params.EmergencyCallerSid)
	}
	if params != nil && params.EmergencyCallingEnabled != nil {
		data.Set("EmergencyCallingEnabled", fmt.Sprint(*params.EmergencyCallingEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.SipRegistration != nil {
		data.Set("SipRegistration", fmt.Sprint(*params.SipRegistration))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceStatusCallbackMethod != nil {
		data.Set("VoiceStatusCallbackMethod", *params.VoiceStatusCallbackMethod)
	}
	if params != nil && params.VoiceStatusCallbackUrl != nil {
		data.Set("VoiceStatusCallbackUrl", *params.VoiceStatusCallbackUrl)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipDomain{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'DeleteSipDomain'
type DeleteSipDomainParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to delete.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *DeleteSipDomainParams) SetPathAccountSid(PathAccountSid string) *DeleteSipDomainParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Delete an instance of a Domain
func (c *ApiService) DeleteSipDomain(Sid string, params *DeleteSipDomainParams) error {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
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

// Optional parameters for the method 'FetchSipDomain'
type FetchSipDomainParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to fetch.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
}

func (params *FetchSipDomainParams) SetPathAccountSid(PathAccountSid string) *FetchSipDomainParams {
	params.PathAccountSid = &PathAccountSid
	return params
}

// Fetch an instance of a Domain
func (c *ApiService) FetchSipDomain(Sid string, params *FetchSipDomainParams) (*ApiV2010SipDomain, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipDomain{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListSipDomain'
type ListSipDomainParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resources to read.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListSipDomainParams) SetPathAccountSid(PathAccountSid string) *ListSipDomainParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *ListSipDomainParams) SetPageSize(PageSize int) *ListSipDomainParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListSipDomainParams) SetLimit(Limit int) *ListSipDomainParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of SipDomain records from the API. Request is executed immediately.
func (c *ApiService) PageSipDomain(params *ListSipDomainParams, pageToken string, pageNumber string) (*ListSipDomainResponse, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains.json"

	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}

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

	ps := &ListSipDomainResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists SipDomain records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListSipDomain(params *ListSipDomainParams) ([]ApiV2010SipDomain, error) {
	if params == nil {
		params = &ListSipDomainParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipDomain(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []ApiV2010SipDomain

	for response != nil {
		records = append(records, response.Domains...)

		var record interface{}
		if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipDomainResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListSipDomainResponse)
	}

	return records, err
}

// Streams SipDomain records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamSipDomain(params *ListSipDomainParams) (chan ApiV2010SipDomain, error) {
	if params == nil {
		params = &ListSipDomainParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageSipDomain(params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan ApiV2010SipDomain, 1)

	go func() {
		for response != nil {
			for item := range response.Domains {
				channel <- response.Domains[item]
			}

			var record interface{}
			if record, err = client.GetNext(c.baseURL, response, &curRecord, params.Limit, c.getNextListSipDomainResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListSipDomainResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListSipDomainResponse(nextPageUrl string) (interface{}, error) {
	if nextPageUrl == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(nextPageUrl, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListSipDomainResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

// Optional parameters for the method 'UpdateSipDomain'
type UpdateSipDomainParams struct {
	// The SID of the [Account](https://www.twilio.com/docs/iam/api/account) that created the SipDomain resource to update.
	PathAccountSid *string `json:"PathAccountSid,omitempty"`
	// The SID of the BYOC Trunk(Bring Your Own Carrier) resource that the Sip Domain will be associated with.
	ByocTrunkSid *string `json:"ByocTrunkSid,omitempty"`
	// The unique address you reserve on Twilio to which you route your SIP traffic. Domain names can contain letters, digits, and \\\"-\\\" and must end with `sip.twilio.com`.
	DomainName *string `json:"DomainName,omitempty"`
	// Whether an emergency caller sid is configured for the domain. If present, this phone number will be used as the callback for the emergency call.
	EmergencyCallerSid *string `json:"EmergencyCallerSid,omitempty"`
	// Whether emergency calling is enabled for the domain. If enabled, allows emergency calls on the domain from phone numbers with validated addresses.
	EmergencyCallingEnabled *bool `json:"EmergencyCallingEnabled,omitempty"`
	// A descriptive string that you created to describe the resource. It can be up to 64 characters long.
	FriendlyName *string `json:"FriendlyName,omitempty"`
	// Whether secure SIP is enabled for the domain. If enabled, TLS will be enforced and SRTP will be negotiated on all incoming calls to this sip domain.
	Secure *bool `json:"Secure,omitempty"`
	// Whether to allow SIP Endpoints to register with the domain to receive calls. Can be `true` or `false`. `true` allows SIP Endpoints to register with the domain to receive calls, `false` does not.
	SipRegistration *bool `json:"SipRegistration,omitempty"`
	// The HTTP method we should use to call `voice_fallback_url`. Can be: `GET` or `POST`.
	VoiceFallbackMethod *string `json:"VoiceFallbackMethod,omitempty"`
	// The URL that we should call when an error occurs while retrieving or executing the TwiML requested by `voice_url`.
	VoiceFallbackUrl *string `json:"VoiceFallbackUrl,omitempty"`
	// The HTTP method we should use to call `voice_url`
	VoiceMethod *string `json:"VoiceMethod,omitempty"`
	// The HTTP method we should use to call `voice_status_callback_url`. Can be: `GET` or `POST`.
	VoiceStatusCallbackMethod *string `json:"VoiceStatusCallbackMethod,omitempty"`
	// The URL that we should call to pass status parameters (such as call ended) to your application.
	VoiceStatusCallbackUrl *string `json:"VoiceStatusCallbackUrl,omitempty"`
	// The URL we should call when the domain receives a call.
	VoiceUrl *string `json:"VoiceUrl,omitempty"`
}

func (params *UpdateSipDomainParams) SetPathAccountSid(PathAccountSid string) *UpdateSipDomainParams {
	params.PathAccountSid = &PathAccountSid
	return params
}
func (params *UpdateSipDomainParams) SetByocTrunkSid(ByocTrunkSid string) *UpdateSipDomainParams {
	params.ByocTrunkSid = &ByocTrunkSid
	return params
}
func (params *UpdateSipDomainParams) SetDomainName(DomainName string) *UpdateSipDomainParams {
	params.DomainName = &DomainName
	return params
}
func (params *UpdateSipDomainParams) SetEmergencyCallerSid(EmergencyCallerSid string) *UpdateSipDomainParams {
	params.EmergencyCallerSid = &EmergencyCallerSid
	return params
}
func (params *UpdateSipDomainParams) SetEmergencyCallingEnabled(EmergencyCallingEnabled bool) *UpdateSipDomainParams {
	params.EmergencyCallingEnabled = &EmergencyCallingEnabled
	return params
}
func (params *UpdateSipDomainParams) SetFriendlyName(FriendlyName string) *UpdateSipDomainParams {
	params.FriendlyName = &FriendlyName
	return params
}
func (params *UpdateSipDomainParams) SetSecure(Secure bool) *UpdateSipDomainParams {
	params.Secure = &Secure
	return params
}
func (params *UpdateSipDomainParams) SetSipRegistration(SipRegistration bool) *UpdateSipDomainParams {
	params.SipRegistration = &SipRegistration
	return params
}
func (params *UpdateSipDomainParams) SetVoiceFallbackMethod(VoiceFallbackMethod string) *UpdateSipDomainParams {
	params.VoiceFallbackMethod = &VoiceFallbackMethod
	return params
}
func (params *UpdateSipDomainParams) SetVoiceFallbackUrl(VoiceFallbackUrl string) *UpdateSipDomainParams {
	params.VoiceFallbackUrl = &VoiceFallbackUrl
	return params
}
func (params *UpdateSipDomainParams) SetVoiceMethod(VoiceMethod string) *UpdateSipDomainParams {
	params.VoiceMethod = &VoiceMethod
	return params
}
func (params *UpdateSipDomainParams) SetVoiceStatusCallbackMethod(VoiceStatusCallbackMethod string) *UpdateSipDomainParams {
	params.VoiceStatusCallbackMethod = &VoiceStatusCallbackMethod
	return params
}
func (params *UpdateSipDomainParams) SetVoiceStatusCallbackUrl(VoiceStatusCallbackUrl string) *UpdateSipDomainParams {
	params.VoiceStatusCallbackUrl = &VoiceStatusCallbackUrl
	return params
}
func (params *UpdateSipDomainParams) SetVoiceUrl(VoiceUrl string) *UpdateSipDomainParams {
	params.VoiceUrl = &VoiceUrl
	return params
}

// Update the attributes of a domain
func (c *ApiService) UpdateSipDomain(Sid string, params *UpdateSipDomainParams) (*ApiV2010SipDomain, error) {
	path := "/2010-04-01/Accounts/{AccountSid}/SIP/Domains/{Sid}.json"
	if params != nil && params.PathAccountSid != nil {
		path = strings.Replace(path, "{"+"AccountSid"+"}", *params.PathAccountSid, -1)
	} else {
		path = strings.Replace(path, "{"+"AccountSid"+"}", c.requestHandler.Client.AccountSid(), -1)
	}
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	if params != nil && params.ByocTrunkSid != nil {
		data.Set("ByocTrunkSid", *params.ByocTrunkSid)
	}
	if params != nil && params.DomainName != nil {
		data.Set("DomainName", *params.DomainName)
	}
	if params != nil && params.EmergencyCallerSid != nil {
		data.Set("EmergencyCallerSid", *params.EmergencyCallerSid)
	}
	if params != nil && params.EmergencyCallingEnabled != nil {
		data.Set("EmergencyCallingEnabled", fmt.Sprint(*params.EmergencyCallingEnabled))
	}
	if params != nil && params.FriendlyName != nil {
		data.Set("FriendlyName", *params.FriendlyName)
	}
	if params != nil && params.Secure != nil {
		data.Set("Secure", fmt.Sprint(*params.Secure))
	}
	if params != nil && params.SipRegistration != nil {
		data.Set("SipRegistration", fmt.Sprint(*params.SipRegistration))
	}
	if params != nil && params.VoiceFallbackMethod != nil {
		data.Set("VoiceFallbackMethod", *params.VoiceFallbackMethod)
	}
	if params != nil && params.VoiceFallbackUrl != nil {
		data.Set("VoiceFallbackUrl", *params.VoiceFallbackUrl)
	}
	if params != nil && params.VoiceMethod != nil {
		data.Set("VoiceMethod", *params.VoiceMethod)
	}
	if params != nil && params.VoiceStatusCallbackMethod != nil {
		data.Set("VoiceStatusCallbackMethod", *params.VoiceStatusCallbackMethod)
	}
	if params != nil && params.VoiceStatusCallbackUrl != nil {
		data.Set("VoiceStatusCallbackUrl", *params.VoiceStatusCallbackUrl)
	}
	if params != nil && params.VoiceUrl != nil {
		data.Set("VoiceUrl", *params.VoiceUrl)
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ApiV2010SipDomain{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

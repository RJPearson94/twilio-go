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
)

// Retrieve voice dialing permissions inheritance for the sub-account
func (c *ApiService) FetchDialingPermissionsSettings() (*VoiceV1DialingPermissionsSettings, error) {
	path := "/v1/Settings"

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1DialingPermissionsSettings{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateDialingPermissionsSettings'
type UpdateDialingPermissionsSettingsParams struct {
	// `true` for the sub-account to inherit voice dialing permissions from the Master Project; otherwise `false`.
	DialingPermissionsInheritance *bool `json:"DialingPermissionsInheritance,omitempty"`
}

func (params *UpdateDialingPermissionsSettingsParams) SetDialingPermissionsInheritance(DialingPermissionsInheritance bool) *UpdateDialingPermissionsSettingsParams {
	params.DialingPermissionsInheritance = &DialingPermissionsInheritance
	return params
}

// Update voice dialing permissions inheritance for the sub-account
func (c *ApiService) UpdateDialingPermissionsSettings(params *UpdateDialingPermissionsSettingsParams) (*VoiceV1DialingPermissionsSettings, error) {
	path := "/v1/Settings"

	data := url.Values{}
	if params != nil && params.DialingPermissionsInheritance != nil {
		data.Set("DialingPermissionsInheritance", fmt.Sprint(*params.DialingPermissionsInheritance))
	}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VoiceV1DialingPermissionsSettings{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

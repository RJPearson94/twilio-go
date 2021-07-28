/*
 * Twilio - Video
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
	"net/url"

	"strings"
)

// Returns a list of Recording Rules for the Room.
func (c *ApiService) FetchRoomRecordingRule(RoomSid string) (*VideoV1RoomRoomRecordingRule, error) {
	path := "/v1/Rooms/{RoomSid}/RecordingRules"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRoomRecordingRule{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'UpdateRoomRecordingRule'
type UpdateRoomRecordingRuleParams struct {
	// A JSON-encoded array of recording rules.
	Rules *map[string]interface{} `json:"Rules,omitempty"`
}

func (params *UpdateRoomRecordingRuleParams) SetRules(Rules map[string]interface{}) *UpdateRoomRecordingRuleParams {
	params.Rules = &Rules
	return params
}

// Update the Recording Rules for the Room
func (c *ApiService) UpdateRoomRecordingRule(RoomSid string, params *UpdateRoomRecordingRuleParams) (*VideoV1RoomRoomRecordingRule, error) {
	path := "/v1/Rooms/{RoomSid}/RecordingRules"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.Rules != nil {
		v, err := json.Marshal(params.Rules)

		if err != nil {
			return nil, err
		}

		data.Set("Rules", string(v))
	}

	resp, err := c.requestHandler.Post(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRoomRecordingRule{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

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
	"fmt"
	"net/url"

	"strings"

	"github.com/twilio/twilio-go/client"
)

// Returns a single Track resource represented by &#x60;track_sid&#x60;.  Note: This is one resource with the Video API that requires a SID, be Track Name on the subscriber side is not guaranteed to be unique.
func (c *ApiService) FetchRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, Sid string) (*VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks/{Sid}"
	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListRoomParticipantSubscribedTrack'
type ListRoomParticipantSubscribedTrackParams struct {
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int `json:"PageSize,omitempty"`
	// Max number of records to return.
	Limit *int `json:"limit,omitempty"`
}

func (params *ListRoomParticipantSubscribedTrackParams) SetPageSize(PageSize int) *ListRoomParticipantSubscribedTrackParams {
	params.PageSize = &PageSize
	return params
}
func (params *ListRoomParticipantSubscribedTrackParams) SetLimit(Limit int) *ListRoomParticipantSubscribedTrackParams {
	params.Limit = &Limit
	return params
}

// Retrieve a single page of RoomParticipantSubscribedTrack records from the API. Request is executed immediately.
func (c *ApiService) PageRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams, pageToken string, pageNumber string) (*ListRoomParticipantSubscribedTrackResponse, error) {
	path := "/v1/Rooms/{RoomSid}/Participants/{ParticipantSid}/SubscribedTracks"

	path = strings.Replace(path, "{"+"RoomSid"+"}", RoomSid, -1)
	path = strings.Replace(path, "{"+"ParticipantSid"+"}", ParticipantSid, -1)

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

	ps := &ListRoomParticipantSubscribedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Lists RoomParticipantSubscribedTrack records from the API as a list. Unlike stream, this operation is eager and loads 'limit' records into memory before returning.
func (c *ApiService) ListRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams) ([]VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack, error) {
	if params == nil {
		params = &ListRoomParticipantSubscribedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipantSubscribedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	var records []VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack

	for response != nil {
		records = append(records, response.SubscribedTracks...)

		var record interface{}
		if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListRoomParticipantSubscribedTrackResponse); record == nil || err != nil {
			return records, err
		}

		response = record.(*ListRoomParticipantSubscribedTrackResponse)
	}

	return records, err
}

// Streams RoomParticipantSubscribedTrack records from the API as a channel stream. This operation lazily loads records as efficiently as possible until the limit is reached.
func (c *ApiService) StreamRoomParticipantSubscribedTrack(RoomSid string, ParticipantSid string, params *ListRoomParticipantSubscribedTrackParams) (chan VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack, error) {
	if params == nil {
		params = &ListRoomParticipantSubscribedTrackParams{}
	}
	params.SetPageSize(client.ReadLimits(params.PageSize, params.Limit))

	response, err := c.PageRoomParticipantSubscribedTrack(RoomSid, ParticipantSid, params, "", "")
	if err != nil {
		return nil, err
	}

	curRecord := 0
	//set buffer size of the channel to 1
	channel := make(chan VideoV1RoomRoomParticipantRoomParticipantSubscribedTrack, 1)

	go func() {
		for response != nil {
			for item := range response.SubscribedTracks {
				channel <- response.SubscribedTracks[item]
			}

			var record interface{}
			if record, err = client.GetNext(response, &curRecord, params.Limit, c.getNextListRoomParticipantSubscribedTrackResponse); record == nil || err != nil {
				close(channel)
				return
			}

			response = record.(*ListRoomParticipantSubscribedTrackResponse)
		}
		close(channel)
	}()

	return channel, err
}

func (c *ApiService) getNextListRoomParticipantSubscribedTrackResponse(nextPageUri string) (interface{}, error) {
	if nextPageUri == "" {
		return nil, nil
	}
	resp, err := c.requestHandler.Get(c.baseURL+nextPageUri, nil, nil)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListRoomParticipantSubscribedTrackResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}
	return ps, nil
}

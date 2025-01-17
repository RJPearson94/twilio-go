/*
 * Twilio - Monitor
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.16.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
	"net/url"

	"strings"
	"time"

	twilio "github.com/twilio/twilio-go/client"
)

type DefaultApiService struct {
	baseURL        string
	requestHandler *twilio.RequestHandler
}

func NewDefaultApiService(requestHandler *twilio.RequestHandler) *DefaultApiService {
	return &DefaultApiService{
		requestHandler: requestHandler,
		baseURL:        "https://monitor.twilio.com",
	}
}

func NewDefaultApiServiceWithClient(client twilio.BaseClient) *DefaultApiService {
	return NewDefaultApiService(twilio.NewRequestHandler(client))
}

func (c *DefaultApiService) FetchAlert(Sid string) (*MonitorV1AlertInstance, error) {
	path := "/v1/Alerts/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1AlertInstance{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

func (c *DefaultApiService) FetchEvent(Sid string) (*MonitorV1Event, error) {
	path := "/v1/Events/{Sid}"
	path = strings.Replace(path, "{"+"Sid"+"}", Sid, -1)

	data := url.Values{}
	headers := make(map[string]interface{})

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &MonitorV1Event{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListAlert'
type ListAlertParams struct {
	// Only show alerts for this log-level.  Can be: `error`, `warning`, `notice`, or `debug`.
	LogLevel *string `json:"LogLevel,omitempty"`
	// Only include alerts that occurred on or after this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only include alerts that occurred on or before this date and time. Specify the date and time in GMT and format as `YYYY-MM-DD` or `YYYY-MM-DDThh:mm:ssZ`. Queries for alerts older than 30 days are not supported.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListAlertParams) SetLogLevel(LogLevel string) *ListAlertParams {
	params.LogLevel = &LogLevel
	return params
}
func (params *ListAlertParams) SetStartDate(StartDate time.Time) *ListAlertParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListAlertParams) SetEndDate(EndDate time.Time) *ListAlertParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListAlertParams) SetPageSize(PageSize int32) *ListAlertParams {
	params.PageSize = &PageSize
	return params
}

func (c *DefaultApiService) ListAlert(params *ListAlertParams) (*ListAlertResponse, error) {
	path := "/v1/Alerts"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.LogLevel != nil {
		data.Set("LogLevel", *params.LogLevel)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListAlertResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

// Optional parameters for the method 'ListEvent'
type ListEventParams struct {
	// Only include events initiated by this Actor. Useful for auditing actions taken by specific users or API credentials.
	ActorSid *string `json:"ActorSid,omitempty"`
	// Only include events of this [Event Type](https://www.twilio.com/docs/usage/monitor-events#event-types).
	EventType *string `json:"EventType,omitempty"`
	// Only include events that refer to this resource. Useful for discovering the history of a specific resource.
	ResourceSid *string `json:"ResourceSid,omitempty"`
	// Only include events that originated from this IP address. Useful for tracking suspicious activity originating from the API or the Twilio Console.
	SourceIpAddress *string `json:"SourceIpAddress,omitempty"`
	// Only include events that occurred on or after this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	StartDate *time.Time `json:"StartDate,omitempty"`
	// Only include events that occurred on or before this date. Specify the date in GMT and [ISO 8601](https://en.wikipedia.org/wiki/ISO_8601) format.
	EndDate *time.Time `json:"EndDate,omitempty"`
	// How many resources to return in each list page. The default is 50, and the maximum is 1000.
	PageSize *int32 `json:"PageSize,omitempty"`
}

func (params *ListEventParams) SetActorSid(ActorSid string) *ListEventParams {
	params.ActorSid = &ActorSid
	return params
}
func (params *ListEventParams) SetEventType(EventType string) *ListEventParams {
	params.EventType = &EventType
	return params
}
func (params *ListEventParams) SetResourceSid(ResourceSid string) *ListEventParams {
	params.ResourceSid = &ResourceSid
	return params
}
func (params *ListEventParams) SetSourceIpAddress(SourceIpAddress string) *ListEventParams {
	params.SourceIpAddress = &SourceIpAddress
	return params
}
func (params *ListEventParams) SetStartDate(StartDate time.Time) *ListEventParams {
	params.StartDate = &StartDate
	return params
}
func (params *ListEventParams) SetEndDate(EndDate time.Time) *ListEventParams {
	params.EndDate = &EndDate
	return params
}
func (params *ListEventParams) SetPageSize(PageSize int32) *ListEventParams {
	params.PageSize = &PageSize
	return params
}

// Returns a list of events in the account, sorted by event-date.
func (c *DefaultApiService) ListEvent(params *ListEventParams) (*ListEventResponse, error) {
	path := "/v1/Events"

	data := url.Values{}
	headers := make(map[string]interface{})

	if params != nil && params.ActorSid != nil {
		data.Set("ActorSid", *params.ActorSid)
	}
	if params != nil && params.EventType != nil {
		data.Set("EventType", *params.EventType)
	}
	if params != nil && params.ResourceSid != nil {
		data.Set("ResourceSid", *params.ResourceSid)
	}
	if params != nil && params.SourceIpAddress != nil {
		data.Set("SourceIpAddress", *params.SourceIpAddress)
	}
	if params != nil && params.StartDate != nil {
		data.Set("StartDate", fmt.Sprint((*params.StartDate).Format(time.RFC3339)))
	}
	if params != nil && params.EndDate != nil {
		data.Set("EndDate", fmt.Sprint((*params.EndDate).Format(time.RFC3339)))
	}
	if params != nil && params.PageSize != nil {
		data.Set("PageSize", fmt.Sprint(*params.PageSize))
	}

	resp, err := c.requestHandler.Get(c.baseURL+path, data, headers)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	ps := &ListEventResponse{}
	if err := json.NewDecoder(resp.Body).Decode(ps); err != nil {
		return nil, err
	}

	return ps, err
}

/*
 * Twilio - Insights
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.8.0
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// ListEventResponse struct for ListEventResponse
type ListEventResponse struct {
	Events []InsightsV1CallEvent `json:"Events,omitempty"`
	Meta ListVideoRoomSummaryResponseMeta `json:"Meta,omitempty"`
}
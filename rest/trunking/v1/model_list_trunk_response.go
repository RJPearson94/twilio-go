/*
 * Twilio - Trunking
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListTrunkResponse struct for ListTrunkResponse
type ListTrunkResponse struct {
	Meta   ListTrunkResponseMeta `json:"meta,omitempty"`
	Trunks []TrunkingV1Trunk     `json:"trunks,omitempty"`
}

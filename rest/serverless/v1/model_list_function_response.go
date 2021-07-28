/*
 * Twilio - Serverless
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ListFunctionResponse struct for ListFunctionResponse
type ListFunctionResponse struct {
	Functions []ServerlessV1ServiceFunction `json:"functions,omitempty"`
	Meta      ListServiceResponseMeta       `json:"meta,omitempty"`
}

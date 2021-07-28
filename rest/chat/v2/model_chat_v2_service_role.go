/*
 * Twilio - Chat
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"time"
)

// ChatV2ServiceRole struct for ChatV2ServiceRole
type ChatV2ServiceRole struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The ISO 8601 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// An array of the permissions the role has been granted
	Permissions *[]string `json:"permissions,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The type of role
	Type *string `json:"type,omitempty"`
	// The absolute URL of the Role resource
	Url *string `json:"url,omitempty"`
}

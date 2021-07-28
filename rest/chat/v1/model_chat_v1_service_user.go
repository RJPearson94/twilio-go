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

// ChatV1ServiceUser struct for ChatV1ServiceUser
type ChatV1ServiceUser struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The JSON string that stores application-specific data
	Attributes *string `json:"attributes,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was created
	DateCreated *time.Time `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT when the resource was last updated
	DateUpdated *time.Time `json:"date_updated,omitempty"`
	// The string that you assigned to describe the resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The string that identifies the resource's User
	Identity *string `json:"identity,omitempty"`
	// Whether the User has a potentially valid Push Notification registration for the Service instance
	IsNotifiable *bool `json:"is_notifiable,omitempty"`
	// Whether the User is actively connected to the Service instance and online
	IsOnline *bool `json:"is_online,omitempty"`
	// The number of Channels this User is a Member of
	JoinedChannelsCount *int `json:"joined_channels_count,omitempty"`
	// The absolute URLs of the Channel and Binding resources related to the user
	Links *map[string]interface{} `json:"links,omitempty"`
	// The SID of the assigned to the user
	RoleSid *string `json:"role_sid,omitempty"`
	// The SID of the Service that the resource is associated with
	ServiceSid *string `json:"service_sid,omitempty"`
	// The unique string that identifies the resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the User resource
	Url *string `json:"url,omitempty"`
}

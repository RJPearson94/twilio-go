/*
 * Twilio - Api
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// ApiV2010AccountAuthorizedConnectApp struct for ApiV2010AccountAuthorizedConnectApp
type ApiV2010AccountAuthorizedConnectApp struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The company name set for the Connect App
	ConnectAppCompanyName *string `json:"connect_app_company_name,omitempty"`
	// A detailed description of the app
	ConnectAppDescription *string `json:"connect_app_description,omitempty"`
	// The name of the Connect App
	ConnectAppFriendlyName *string `json:"connect_app_friendly_name,omitempty"`
	// The public URL for the Connect App
	ConnectAppHomepageUrl *string `json:"connect_app_homepage_url,omitempty"`
	// The SID that we assigned to the Connect App
	ConnectAppSid *string `json:"connect_app_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// Permissions authorized to the app
	Permissions *[]string `json:"permissions,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}

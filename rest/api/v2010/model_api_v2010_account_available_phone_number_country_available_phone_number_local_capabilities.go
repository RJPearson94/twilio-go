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

// ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities Whether a phone number can receive calls or messages
type ApiV2010AccountAvailablePhoneNumberCountryAvailablePhoneNumberLocalCapabilities struct {
	Fax   bool `json:"fax,omitempty"`
	Mms   bool `json:"mms,omitempty"`
	Sms   bool `json:"sms,omitempty"`
	Voice bool `json:"voice,omitempty"`
}

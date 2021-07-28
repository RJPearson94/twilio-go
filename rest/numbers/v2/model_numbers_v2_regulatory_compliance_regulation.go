/*
 * Twilio - Numbers
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.19.1
 * Contact: support@twilio.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

// NumbersV2RegulatoryComplianceRegulation struct for NumbersV2RegulatoryComplianceRegulation
type NumbersV2RegulatoryComplianceRegulation struct {
	// The type of End User of the Regulation resource
	EndUserType *string `json:"end_user_type,omitempty"`
	// A human-readable description of the Regulation resource
	FriendlyName *string `json:"friendly_name,omitempty"`
	// The ISO country code of the phone number's country
	IsoCountry *string `json:"iso_country,omitempty"`
	// The type of phone number restricted by the regulatory requirement
	NumberType *string `json:"number_type,omitempty"`
	// The sid of a regulation object that dictates requirements
	Requirements *map[string]interface{} `json:"requirements,omitempty"`
	// The unique string that identifies the Regulation resource
	Sid *string `json:"sid,omitempty"`
	// The absolute URL of the Regulation resource
	Url *string `json:"url,omitempty"`
}

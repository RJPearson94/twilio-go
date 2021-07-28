/*
 * Twilio - Events
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

// EventsV1Schema struct for EventsV1Schema
type EventsV1Schema struct {
	// Schema Identifier.
	Id *string `json:"id,omitempty"`
	// Latest schema version.
	LatestVersion *int `json:"latest_version,omitempty"`
	// The date that the latest schema version was created.
	LatestVersionDateCreated *time.Time `json:"latest_version_date_created,omitempty"`
	// Nested resource URLs.
	Links *map[string]interface{} `json:"links,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}

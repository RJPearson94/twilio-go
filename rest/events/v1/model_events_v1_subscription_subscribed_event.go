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

// EventsV1SubscriptionSubscribedEvent struct for EventsV1SubscriptionSubscribedEvent
type EventsV1SubscriptionSubscribedEvent struct {
	// Account SID.
	AccountSid *string `json:"account_sid,omitempty"`
	// The schema version that the subscription should use.
	SchemaVersion *int `json:"schema_version,omitempty"`
	// Subscription SID.
	SubscriptionSid *string `json:"subscription_sid,omitempty"`
	// Type of event being subscribed to.
	Type *string `json:"type,omitempty"`
	// The URL of this resource.
	Url *string `json:"url,omitempty"`
}

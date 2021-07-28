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

// ApiV2010AccountConferenceParticipant struct for ApiV2010AccountConferenceParticipant
type ApiV2010AccountConferenceParticipant struct {
	// The SID of the Account that created the resource
	AccountSid *string `json:"account_sid,omitempty"`
	// The SID of the Call the resource is associated with
	CallSid *string `json:"call_sid,omitempty"`
	// The SID of the participant who is being `coached`
	CallSidToCoach *string `json:"call_sid_to_coach,omitempty"`
	// Indicates if the participant changed to coach
	Coaching *bool `json:"coaching,omitempty"`
	// The SID of the conference the participant is in
	ConferenceSid *string `json:"conference_sid,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was created
	DateCreated *string `json:"date_created,omitempty"`
	// The RFC 2822 date and time in GMT that the resource was last updated
	DateUpdated *string `json:"date_updated,omitempty"`
	// Whether the conference ends when the participant leaves
	EndConferenceOnExit *bool `json:"end_conference_on_exit,omitempty"`
	// Whether the participant is on hold
	Hold *bool `json:"hold,omitempty"`
	// The label of this participant
	Label *string `json:"label,omitempty"`
	// Whether the participant is muted
	Muted *bool `json:"muted,omitempty"`
	// Whether the conference starts when the participant joins the conference
	StartConferenceOnEnter *bool `json:"start_conference_on_enter,omitempty"`
	// The status of the participant's call in a session
	Status *string `json:"status,omitempty"`
	// The URI of the resource, relative to `https://api.twilio.com`
	Uri *string `json:"uri,omitempty"`
}

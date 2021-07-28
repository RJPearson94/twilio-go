/*
 * Twilio - Insights
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

// InsightsV1VideoRoomSummaryVideoParticipantSummary struct for InsightsV1VideoRoomSummaryVideoParticipantSummary
type InsightsV1VideoRoomSummaryVideoParticipantSummary struct {
	// Account SID associated with the room.
	AccountSid *string `json:"account_sid,omitempty"`
	// Codecs detected from the participant.
	Codecs *[]string `json:"codecs,omitempty"`
	// Amount of time in seconds the participant was in the room.
	DurationSec *int `json:"duration_sec,omitempty"`
	// Name of the edge location the participant connected to.
	EdgeLocation *string `json:"edge_location,omitempty"`
	// Reason the participant left the room.
	EndReason *string `json:"end_reason,omitempty"`
	// Errors encountered by the participant.
	ErrorCode *int `json:"error_code,omitempty"`
	// Twilio error code dictionary link.
	ErrorCodeUrl *string `json:"error_code_url,omitempty"`
	// When the participant joined the room.
	JoinTime *time.Time `json:"join_time,omitempty"`
	// When the participant left the room
	LeaveTime *time.Time `json:"leave_time,omitempty"`
	// Twilio media region the participant connected to.
	MediaRegion *string `json:"media_region,omitempty"`
	// The application-defined string that uniquely identifies the participant within a Room.
	ParticipantIdentity *string `json:"participant_identity,omitempty"`
	// Unique identifier for the participant.
	ParticipantSid *string `json:"participant_sid,omitempty"`
	// Object containing information about the participant's data from the room.
	Properties *map[string]interface{} `json:"properties,omitempty"`
	// Object containing information about the SDK name and version.
	PublisherInfo *map[string]interface{} `json:"publisher_info,omitempty"`
	// Unique identifier for the room.
	RoomSid *string `json:"room_sid,omitempty"`
	// Status of the room.
	Status *string `json:"status,omitempty"`
	// URL of the participant resource.
	Url *string `json:"url,omitempty"`
}

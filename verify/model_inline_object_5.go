/*
 * Twilio - Verify
 *
 * This is the public Twilio REST API.
 *
 * API version: 1.0.1
 * Contact: support@twilio.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package twilio
// InlineObject5 struct for InlineObject5
type InlineObject5 struct {
	// A unique binding for this Factor that identifies it. E.g. the algorithm and public key for `push` factors. It must be a json string with the required properties for the given factor type. Required when creating a new Factor. This value is never returned because it can contain customer secrets.
	Binding string `json:"Binding"`
	// The config required for this Factor. It must be a json string with the required properties for the given factor type
	Config string `json:"Config"`
	// The Type of this Factor. Currently only `push` is supported
	FactorType string `json:"FactorType"`
	// The friendly name of this Factor
	FriendlyName string `json:"FriendlyName"`
}
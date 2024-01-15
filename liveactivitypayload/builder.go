// Package payload is a helper package which contains a payload
// builder to make constructing notification payloads easier.
package liveacvititypayload

import (
	"encoding/json"
)

// InterruptionLevel defines the value for the payload aps interruption-level
type EInterruptionLevel string

const (
	// InterruptionLevelPassive is used to indicate that notification be delivered in a passive manner.
	InterruptionLevelPassive EInterruptionLevel = "passive"

	// InterruptionLevelActive is used to indicate the importance and delivery timing of a notification.
	InterruptionLevelActive EInterruptionLevel = "active"

	// InterruptionLevelTimeSensitive is used to indicate the importance and delivery timing of a notification.
	InterruptionLevelTimeSensitive EInterruptionLevel = "time-sensitive"

	// InterruptionLevelCritical is used to indicate the importance and delivery timing of a notification.
	// This interruption level requires an approved entitlement from Apple.
	// See: https://developer.apple.com/documentation/usernotifications/unnotificationinterruptionlevel/
	InterruptionLevelCritical EInterruptionLevel = "critical"
)

// Payload represents a notification which holds the content that will be
// marshalled as JSON.
type Payload struct {
	content map[string]interface{}
}

type aps struct {
	Alert          interface{} `json:"alert,omitempty"`
	Timestamp      int64       `json:"timestamp"`
	Event          string      `json:"event"`
	ContentState   interface{} `json:"content-state"`
	AttributesType string      `json:"attributes-type"`
	Attributes     interface{} `json:"attributes,omitempty"`
	DismissalDate  int64       `json:"dismissal-date,omitempty"`
}

// NewPayload returns a new Payload struct
func NewPayload() *Payload {
	return &Payload{
		map[string]interface{}{
			"aps": &aps{},
		},
	}
}

// Alert sets the aps alert on the payload.
// This will display a notification alert message to the user.
//
//	{"aps":{"alert":alert}}`
func (p *Payload) Alert(alert interface{}) *Payload {
	p.aps().Alert = alert
	return p
}

// Custom payload

// Custom sets a custom key and value on the payload.
// This will add custom key/value data to the notification payload at root level.
//
//	{"aps":{}, key:value}
func (p *Payload) Custom(key string, val interface{}) *Payload {
	p.content[key] = val
	return p
}

// Mdm sets the mdm on the payload.
// This is for Apple Mobile Device Management (mdm) payloads.
//
//	{"aps":{}:"mdm":mdm}
func (p *Payload) Mdm(mdm string) *Payload {
	p.content["mdm"] = mdm
	return p
}

func (p *Payload) Timestamp(t int64) *Payload {
	p.aps().Timestamp = t
	return p
}

func (p *Payload) Event(event string) *Payload {
	p.aps().Event = event
	return p
}

func (p *Payload) ContentState(contentState interface{}) *Payload {
	p.aps().ContentState = contentState
	return p
}

func (p *Payload) AttributesType(attributesType string) *Payload {
	p.aps().AttributesType = attributesType
	return p
}

func (p *Payload) Attributes() *Payload {
	var obj = make(map[string]interface{})

	p.aps().Attributes = obj
	return p
}

func (p *Payload) DismissalDate(t int64) *Payload {

	p.aps().DismissalDate = t
	return p
}

// MarshalJSON returns the JSON encoded version of the Payload
func (p *Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.content)
}

func (p *Payload) aps() *aps {
	return p.content["aps"].(*aps)
}

// Package payload is a helper package which contains a payload
// builder to make constructing notification payloads easier.
package payload1

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
	ContentState   interface{} `json:"content_state"`
	AttributesType string      `json:"attributes_type"`
	Attributes     interface{} `json:"attributes"`
}

type alert struct {
	Action          string   `json:"action,omitempty"`
	ActionLocKey    string   `json:"action-loc-key,omitempty"`
	Body            string   `json:"body,omitempty"`
	LaunchImage     string   `json:"launch-image,omitempty"`
	LocArgs         []string `json:"loc-args,omitempty"`
	LocKey          string   `json:"loc-key,omitempty"`
	Title           string   `json:"title,omitempty"`
	Subtitle        string   `json:"subtitle,omitempty"`
	TitleLocArgs    []string `json:"title-loc-args,omitempty"`
	TitleLocKey     string   `json:"title-loc-key,omitempty"`
	SummaryArg      string   `json:"summary-arg,omitempty"`
	SummaryArgCount int      `json:"summary-arg-count,omitempty"`
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

// Alert dictionary

// AlertTitle sets the aps alert title on the payload.
// This will display a short string describing the purpose of the notification.
// Apple Watch & Safari display this string as part of the notification interface.
//
//	{"aps":{"alert":{"title":title}}}
func (p *Payload) AlertTitle(title string) *Payload {
	p.aps().alert().Title = title
	return p
}

// AlertTitleLocKey sets the aps alert title localization key on the payload.
// This is the key to a title string in the Localizable.strings file for the
// current localization. See Localized Formatted Strings in Apple documentation
// for more information.
//
//	{"aps":{"alert":{"title-loc-key":key}}}
func (p *Payload) AlertTitleLocKey(key string) *Payload {
	p.aps().alert().TitleLocKey = key
	return p
}

// AlertTitleLocArgs sets the aps alert title localization args on the payload.
// These are the variable string values to appear in place of the format
// specifiers in title-loc-key. See Localized Formatted Strings in Apple
// documentation for more information.
//
//	{"aps":{"alert":{"title-loc-args":args}}}
func (p *Payload) AlertTitleLocArgs(args []string) *Payload {
	p.aps().alert().TitleLocArgs = args
	return p
}

// AlertSubtitle sets the aps alert subtitle on the payload.
// This will display a short string describing the purpose of the notification.
// Apple Watch & Safari display this string as part of the notification interface.
//
//	{"aps":{"alert":{"subtitle":"subtitle"}}}
func (p *Payload) AlertSubtitle(subtitle string) *Payload {
	p.aps().alert().Subtitle = subtitle
	return p
}

// AlertBody sets the aps alert body on the payload.
// This is the text of the alert message.
//
//	{"aps":{"alert":{"body":body}}}
func (p *Payload) AlertBody(body string) *Payload {
	p.aps().alert().Body = body
	return p
}

// AlertLaunchImage sets the aps launch image on the payload.
// This is the filename of an image file in the app bundle. The image is used
// as the launch image when users tap the action button or move the action
// slider.
//
//	{"aps":{"alert":{"launch-image":image}}}
func (p *Payload) AlertLaunchImage(image string) *Payload {
	p.aps().alert().LaunchImage = image
	return p
}

// AlertLocArgs sets the aps alert localization args on the payload.
// These are the variable string values to appear in place of the format
// specifiers in loc-key. See Localized Formatted Strings in Apple
// documentation for more information.
//
//	{"aps":{"alert":{"loc-args":args}}}
func (p *Payload) AlertLocArgs(args []string) *Payload {
	p.aps().alert().LocArgs = args
	return p
}

// AlertLocKey sets the aps alert localization key on the payload.
// This is the key to an alert-message string in the Localizable.strings file
// for the current localization. See Localized Formatted Strings in Apple
// documentation for more information.
//
//	{"aps":{"alert":{"loc-key":key}}}
func (p *Payload) AlertLocKey(key string) *Payload {
	p.aps().alert().LocKey = key
	return p
}

// AlertAction sets the aps alert action on the payload.
// This is the label of the action button, if the user sets the notifications
// to appear as alerts. This label should be succinct, such as “Details” or
// “Read more”. If omitted, the default value is “Show”.
//
//	{"aps":{"alert":{"action":action}}}
func (p *Payload) AlertAction(action string) *Payload {
	p.aps().alert().Action = action
	return p
}

// AlertActionLocKey sets the aps alert action localization key on the payload.
// This is the the string used as a key to get a localized string in the current
// localization to use for the notfication right button’s title instead of
// “View”. See Localized Formatted Strings in Apple documentation for more
// information.
//
//	{"aps":{"alert":{"action-loc-key":key}}}
func (p *Payload) AlertActionLocKey(key string) *Payload {
	p.aps().alert().ActionLocKey = key
	return p
}

// AlertSummaryArg sets the aps alert summary arg key on the payload.
// This is the string that is used as a key to fill in an argument
// at the bottom of a notification to provide more context, such as
// a name associated with the sender of the notification.
//
//	{"aps":{"alert":{"summary-arg":key}}}
func (p *Payload) AlertSummaryArg(key string) *Payload {
	p.aps().alert().SummaryArg = key
	return p
}

// AlertSummaryArgCount sets the aps alert summary arg count key on the payload.
// This integer sets a custom "weight" on the notification, effectively
// allowing a notification to be viewed internally as two. For example if
// a notification encompasses 3 messages, you can set it to 3.
//
//	{"aps":{"alert":{"summary-arg-count":key}}}
func (p *Payload) AlertSummaryArgCount(key int) *Payload {
	p.aps().alert().SummaryArgCount = key
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

// MarshalJSON returns the JSON encoded version of the Payload
func (p *Payload) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.content)
}

func (p *Payload) aps() *aps {
	return p.content["aps"].(*aps)
}

func (a *aps) alert() *alert {
	if _, ok := a.Alert.(*alert); !ok {
		a.Alert = &alert{}
	}
	return a.Alert.(*alert)
}

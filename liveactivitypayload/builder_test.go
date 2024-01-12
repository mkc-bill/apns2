package liveacvititypayload_test

import (
	"encoding/json"
	. "github.com/mkc-bill/apns2/liveactivitypayload"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmptyPayload(t *testing.T) {
	payload := NewPayload()
	b, _ := json.Marshal(payload)
	assert.Equal(t, `{"aps":{}}`, string(b))
}

func TestAlert(t *testing.T) {
	payload := NewPayload().Alert("hello")
	b, _ := json.Marshal(payload)
	assert.Equal(t, `{"aps":{"alert":"hello"}}`, string(b))
}

func TestCustom(t *testing.T) {
	payload := NewPayload().Custom("key", "val")
	b, _ := json.Marshal(payload)
	assert.Equal(t, `{"aps":{},"key":"val"}`, string(b))
}

func TestCustomMap(t *testing.T) {
	payload := NewPayload().Custom("key", map[string]interface{}{
		"map": 1,
	})
	b, _ := json.Marshal(payload)
	assert.Equal(t, `{"aps":{},"key":{"map":1}}`, string(b))
}

func TestMdm(t *testing.T) {
	payload := NewPayload().Mdm("996ac527-9993-4a0a-8528-60b2b3c2f52b")
	b, _ := json.Marshal(payload)
	assert.Equal(t, `{"aps":{},"mdm":"996ac527-9993-4a0a-8528-60b2b3c2f52b"}`, string(b))
}

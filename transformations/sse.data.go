package transformations

// https://api.slack.com/outgoing-webhooks

import (
	"encoding/json"
	_ "log"

	"github.com/whosonfirst/go-webhookd"
)

type SSEDataTransformation struct {
	webhookd.WebhookTransformation
}

func NewSSEDataTransformation() (*SSEDataTransformation, error) {

	p := SSEDataTransformation{}

	return &p, nil
}

func (p *SSEDataTransformation) Transform(body []byte) ([]byte, *webhookd.WebhookError) {

	var obj map[string]interface{}
	if err := json.Unmarshal([]byte(body), &obj); err != nil {
		panic(err)
	}

	b, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}

	s := "data: " + string(b)

	return []byte(s), nil
}

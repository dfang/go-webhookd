package dispatchers

import (
	"github.com/whosonfirst/go-webhookd"
	"github.com/whosonfirst/go-webhookd/config"
)

type NullDispatcher struct {
	webhookd.WebhookDispatcher
}

func NewNullDispatcher(cfg *config.WebhookDispatcherConfig) (*NullDispatcher, error) {

	n := NullDispatcher{}
	return &n, nil
}

func (n *NullDispatcher) Dispatch(body []byte) *webhookd.WebhookError {

	return nil
}

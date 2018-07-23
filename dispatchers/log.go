package dispatchers

// PLEASE MAKE ME MORE SOPHISTICATED (20161016/thisisaaronland)

import (
	"github.com/whosonfirst/go-webhookd"
	"github.com/whosonfirst/go-webhookd/config"
	"log"
)

type LogDispatcher struct {
	webhookd.WebhookDispatcher
}

func NewLogDispatcher(cfg *config.WebhookDispatcherConfig) (*LogDispatcher, error) {

	n := LogDispatcher{}
	return &n, nil
}

func (n *LogDispatcher) Dispatch(body []byte) *webhookd.WebhookError {

	log.Println(string(body))
	return nil
}

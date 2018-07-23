package dispatchers

import (
	"errors"
	"fmt"
	"github.com/whosonfirst/go-webhookd"
	"github.com/whosonfirst/go-webhookd/config"
	"strings"
)

func NewDispatcherFromConfig(cfg *config.WebhookDispatcherConfig) (webhookd.WebhookDispatcher, error) {

	switch strings.ToUpper(cfg.Name) {

	case "LOG":
		return NewLogDispatcher(cfg)
	case "NULL":
		return NewNullDispatcher(cfg)
	case "PUBSUB":
		return NewPubSubDispatcher(cfg)
	case "S3":
		return NewS3Dispatcher(cfg)
	case "SLACK":
		return NewSlackDispatcher(cfg)
	default:
		msg := fmt.Sprintf("Undefined dispatcher: '%s'", cfg.Name)
		return nil, errors.New(msg)
	}
}

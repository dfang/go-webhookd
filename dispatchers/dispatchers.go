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
		return NewLogDispatcher()
	case "NULL":
		return NewNullDispatcher()
	case "PUBSUB":
		return NewPubSubDispatcher(cfg.Host, cfg.Port, cfg.Channel)
	case "S3":
		return NewS3Dispatcher("fix me")
	case "SLACK":
		return NewSlackDispatcher(cfg.Config)
	default:
		msg := fmt.Sprintf("Undefined dispatcher: '%s'", cfg.Name)
		return nil, errors.New(msg)
	}
}

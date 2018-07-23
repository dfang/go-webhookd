package dispatchers

import (
	"github.com/whosonfirst/go-webhookd"
	"github.com/whosonfirst/go-webhookd/config"
	"github.com/whosonfirst/go-writer-slackcat"
)

type SlackDispatcher struct {
	webhookd.WebhookDispatcher
	writer *slackcat.Writer
}

func NewSlackDispatcher(cfg *config.WebhookDispatcherConfig) (*SlackDispatcher, error) {

	slackcat_cfg := cfg.Config // PLEASE REPLACE WITH cfg.DSN
	writer, err := slackcat.NewWriter(slackcat_cfg)

	if err != nil {
		return nil, err
	}

	slack := SlackDispatcher{
		writer: writer,
	}

	return &slack, nil
}

func (sl *SlackDispatcher) Dispatch(body []byte) *webhookd.WebhookError {

	_, err := sl.writer.Write(body)

	if err != nil {
		code := 999
		message := err.Error()

		err := &webhookd.WebhookError{Code: code, Message: message}
		return err
	}

	return nil
}

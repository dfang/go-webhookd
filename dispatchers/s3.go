package dispatchers

import (
	"bytes"
	"github.com/whosonfirst/go-webhookd"
	"github.com/whosonfirst/go-webhookd/config"
	"github.com/whosonfirst/go-whosonfirst-aws/s3"
	"io/ioutil"
)

type S3Dispatcher struct {
	webhookd.WebhookDispatcher
	conn *s3.S3Connection
}

func NewS3Dispatcher(cfg *config.WebhookDispatcherConfig) (*S3Dispatcher, error) {

	s3_cfg, err := s3.NewS3ConfigFromString(cfg.DSN)

	if err != nil {
		return nil, err
	}

	conn, err := s3.NewS3Connection(s3_cfg)

	if err != nil {
		return nil, err
	}

	dispatcher := S3Dispatcher{
		conn: conn,
	}

	return &dispatcher, nil
}

func (dispatcher *S3Dispatcher) Dispatch(body []byte) *webhookd.WebhookError {

	key := "fixme"

	r := bytes.NewReader(body)
	fh := ioutil.NopCloser(r)

	err := dispatcher.conn.Put(key, fh)

	if err != nil {

		code := 999
		message := err.Error()

		err := &webhookd.WebhookError{Code: code, Message: message}
		return err
	}

	return nil
}

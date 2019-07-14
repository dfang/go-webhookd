package transformations

// https://work.weixin.qq.com/api/doc#90000/90136/91770

import (
	"encoding/json"
	"log"

	"github.com/whosonfirst/go-webhookd"
)

type WechatWorkTextTransformation struct {
	webhookd.WebhookTransformation
	msgtype string
}

func NewWechatWorkTextTransformation() (*WechatWorkTextTransformation, error) {

	p := WechatWorkTextTransformation{
		msgtype: "text",
	}

	return &p, nil
}

func (p *WechatWorkTextTransformation) Transform(body []byte) ([]byte, *webhookd.WebhookError) {

	x := WechatGroupBotTextMsg{
		Msgtype: "text",
		Text: struct {
			Content             string   `json:"content"`
			MentionedList       []string `json:"mentioned_list"`
			MentionedMobileList []string `json:"mentioned_mobile_list"`
		}{
			Content: string(body),
		},
	}

	text, err := json.Marshal(x)
	if err != nil {
		log.Println(err)
	}

	if len(text) == 0 {

		code := 999
		message := "Unable to parse WechatWork text"

		err := &webhookd.WebhookError{Code: code, Message: message}
		return nil, err
	}

	return []byte(text), nil
}

type WechatGroupBotTextMsg struct {
	Msgtype string `json:"msgtype"`
	Text    struct {
		Content             string   `json:"content"`
		MentionedList       []string `json:"mentioned_list"`
		MentionedMobileList []string `json:"mentioned_mobile_list"`
	} `json:"text"`
}

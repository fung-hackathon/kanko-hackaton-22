package message

import (
	"bytes"
	"encoding/json"
	"kanko-hackaton-22/app/config"
	"net/http"
)

type Event struct {
	Type            string          `json:"type"`
	Message         Message         `json:"message"`
	WebhookEventID  string          `json:"webhookEventId"`
	DeliveryContext DeliveryContext `json:"deliveryContext"`
	Timestamp       int64           `json:"timestamp"`
	Source          Source          `json:"source"`
	ReplyToken      string          `json:"replyToken"`
	Mode            string          `json:"mode"`
}

type Message struct {
	Type               string      `json:"type"`
	Text               string      `json:"text,omitempty"`
	QuickReply         interface{} `json:"quickReply,omitempty"`
	OriginalContentUrl string      `json:"originalContentUrl,omitempty"`
	PreviewImageUrl    string      `json:"previewImageUrl,omitempty"`
}

type ReplyMessage struct {
	ReplyToken string    `json:"replyToken"`
	Messages   []Message `json:"messages"`
}

type DeliveryContext struct {
	IsRedelivery bool `json:"isRedelivery"`
}

type Source struct {
	Type   string `json:"type"`
	UserID string `json:"userId"`
}

type Request struct {
	Destination string  `json:"destination"`
	Events      []Event `json:"events"`
}

func SendMessage(reply any) error {
	payload, err := json.Marshal(reply)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", "https://api.line.me/v2/bot/message/reply", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.LINE_CHANNEL_ACCESS_TOKEN)
	client := &http.Client{}

	res, err := client.Do(req)
	if err != nil {
		return err
	}

	defer res.Body.Close()

	return nil
}

package dtos

import "time"

type EmailNotificationResp struct {
	StatusCode    int       `json:"statusCode"`
	StatusMessage string    `json:"statusMessage"`
	Message       string    `json:"message"`
	Timestamp     time.Time `json:"timestamp"`
}

type EmailNotificationReq struct {
	TemplateId string      `json:"templateId"`
	Data       interface{} `json:"data"`
}

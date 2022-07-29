package sendgrid

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
)

func Send(ctx context.Context, body string) error {

	client := resty.New()
	resp, err := client.R().
		SetHeader("Content-Type", "application/json").
		SetHeader("X-RapidAPI-Key", configs.Instance.RapidAPIKey).
		SetHeader("X-RapidAPI-Host", configs.Instance.RapidAPIHost).
		SetContext(ctx).
		SetBody(body).
		//SetResult(&SuccessResponse{}).
		//SetError(&errRes).
		Post(configs.Instance.RapidAPIBaseURL + "/mail/send")

	if err != nil || resp.IsError() {
		log.Error("Mail failed to send. %+v: %+v", err, resp.String())
		return err
	}

	return nil
}

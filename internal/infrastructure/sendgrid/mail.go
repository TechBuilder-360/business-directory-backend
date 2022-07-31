package sendgrid

import (
	"context"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/go-resty/resty/v2"
	log "github.com/sirupsen/logrus"
	"github.com/sendgrid/rest"
	"github.com/sendgrid/sendgrid-go"
	m "github.com/sendgrid/sendgrid-go/helpers/mail"
	
)

func Send(ctx context.Context, body string) error {

	client := v2.New()
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

func SendMail(subjectTop string, toMail string, bodyHtml string, toName string) (*rest.Response, error) {
	from := m.NewEmail("TechBuilder Developer", configs.Instance.SendGridFromEmail)
	subject := subjectTop
	to := m.NewEmail(toName, toMail)
	htmlContent := bodyHtml
	message := m.NewSingleEmail(from, subject, to,"", htmlContent)
	client := sendgrid.NewSendClient(configs.Instance.SendGridAPIKey)
	res,err := client.Send(message)
	if err != nil {
		
		return res , err
	} else {
	
		return res, nil
	}
}


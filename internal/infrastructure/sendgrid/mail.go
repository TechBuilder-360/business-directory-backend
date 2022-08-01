package sendgrid

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/sendgrid/sendgrid-go"
	m "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendMail(subject string, toMail string, bodyHtml string, toName string) error {
	from := m.NewEmail("TechBuilder Developer", configs.Instance.SendGridFromEmail)
	to := m.NewEmail(toName, toMail)
	message := m.NewSingleEmail(from, subject, to, "", bodyHtml)
	client := sendgrid.NewSendClient(configs.Instance.SendGridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	} else {
		return nil
	}
}

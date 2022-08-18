package sendgrid

import (
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/flosch/pongo2"
	"github.com/sendgrid/sendgrid-go"
	m "github.com/sendgrid/sendgrid-go/helpers/mail"
)

func parseHTML(body map[string]interface{}) (string, error) {
	var tpl = pongo2.Must(pongo2.FromFile("./templates/template.html"))
	dt, err := tpl.Execute(body)
	if err != nil {
		return "", err
	}

	return dt, nil

}

func sendMail(body *MailBuilder) error {
	from := m.NewEmail("TechBuilder Developer", configs.Instance.SendGridFromEmail)
	to := m.NewEmail(body.ToName, body.ToMail)
	bodyTemplate, err := parseHTML(body.Content)
	if err != nil {
		return err
	}
	message := m.NewSingleEmail(from, body.Subject, to, "", bodyTemplate)
	client := sendgrid.NewSendClient(configs.Instance.SendGridAPIKey)
	_, err = client.Send(message)
	if err != nil {
		return err
	}
	return nil
}

func SentActivateMail(activate *MailBuilder) error {

}

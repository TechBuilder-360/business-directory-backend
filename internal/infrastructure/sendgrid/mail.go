package sendgrid

import (
	"fmt"
	"github.com/TechBuilder-360/business-directory-backend/internal/configs"
	"github.com/flosch/pongo2"
	"github.com/sendgrid/sendgrid-go"
	m "github.com/sendgrid/sendgrid-go/helpers/mail"
	"path"
	"runtime"
)

func parseHTML(body map[string]interface{}, templateName Template) (string, error) {
	_, filename, _, _ := runtime.Caller(0)
	filepath := path.Join(path.Dir(filename), fmt.Sprintf("./templates/%s.html", templateName))
	tpl, err := pongo2.FromFile(filepath)
	dt, err := tpl.Execute(body)
	if err != nil {
		return "", err
	}

	return dt, nil

}

func sendMail(body *mail) error {
	from := m.NewEmail("TechBuilder Developer", configs.Instance.SendGridFromEmail)
	to := m.NewEmail(body.ToName, body.ToMail)
	message := m.NewSingleEmail(from, body.Subject, to, "", body.Template)
	client := sendgrid.NewSendClient(configs.Instance.SendGridAPIKey)
	_, err := client.Send(message)
	if err != nil {
		return err
	}
	return nil
}

func SendActivateMail(activate *ActivationMailRequest) error {
	content := make(map[string]interface{})
	content["fullname"] = activate.FullName
	content["appName"] = configs.Instance.AppName
	content["link"] = fmt.Sprintf("%s/auth/activate?token=%s&uid=%s", configs.Instance.Host, activate.Token, activate.UID)

	template, err := parseHTML(content, ACTIVATIONTEMPLATE)
	if err != nil {
		return err
	}

	message := mail{
		ToName:   activate.ToName,
		ToMail:   activate.ToMail,
		Subject:  configs.Instance.AppName + " account activation",
		Template: template,
	}

	return sendMail(&message)
}

func SendOTPMail(otp *OTPMailRequest) error {
	content := make(map[string]interface{})
	content["name"] = otp.Name
	content["code"] = otp.Code
	content["duration"] = otp.Duration

	template, err := parseHTML(content, OTPTEMPLATE)
	if err != nil {
		return err
	}

	message := mail{
		ToName:   otp.ToName,
		ToMail:   otp.ToMail,
		Subject:  configs.Instance.AppName + " OTP",
		Template: template,
	}

	return sendMail(&message)
}

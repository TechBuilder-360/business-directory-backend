package sendgrid

const (
	ACTIVATIONTEMPLATE Template = "activation_template"
	OTPTEMPLATE        Template = "otp_template"
	GENERALTEMPLATE    Template = "general_template"
)

type ActivationMailRequest struct {
	ToName   string
	ToMail   string
	FullName string
	Token    string
	UID      string
}

type GeneralMailRequest struct {
	ToName  string
	ToMail  string
	Subject string
	Message string
}

type OTPMailRequest struct {
	ToName   string
	ToMail   string
	Code     string
	Name     string
	Duration uint
}

type mail struct {
	ToName   string
	ToMail   string
	Subject  string
	Template string
}

type Template string

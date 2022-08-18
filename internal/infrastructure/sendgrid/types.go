package sendgrid

type MailBuilder struct {
	ToName  string
	ToMail  string
	Subject string
	Content interface{}
}

type ActivationMailRequest struct {
	FullName string
	Token    string
}

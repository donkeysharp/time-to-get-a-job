package providers

type EmailProvider struct {
}

func NewEmailProvider() *EmailProvider {
	return &EmailProvider{}
}

func (me *EmailProvider) SendEmail(destination string, subject string, body string) error {
	return nil
}

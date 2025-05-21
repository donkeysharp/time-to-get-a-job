package providers

type EmailProvider struct {
}

func (me *EmailProvider) SendEmail(destination string, subject string, body string) error {
	return nil
}

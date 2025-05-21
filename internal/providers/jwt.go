package providers

type JWTProvider struct{}

func (me *JWTProvider) CreateJWT(claims map[string]interface{}) (string, error) {
	return "", nil
}

func (me *JWTProvider) ValidateJWT(token string) error {
	return nil
}

package services

func (s *authService) Authenticate(tokenString string) error {
	return s.tokenGenerator.VerifyToken(tokenString)
}

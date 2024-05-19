package services_auth

func (s *authService) ValidateToken(tokenString string) error {
	return s.tokenGenerator.VerifyToken(tokenString)
}

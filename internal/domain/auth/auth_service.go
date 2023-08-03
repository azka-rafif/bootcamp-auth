package auth

type AuthService interface {
	Generate(payload AuthPayload) (JwtResponseFormat, error)
}

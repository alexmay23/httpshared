package user

type UseCase interface {
	GetValidModelFromToken(token string) *Model
	SendCode(phone string)
	VerifyCode(phone string, code int)(*AuthResponse, error)
	Update(model *Model, data map[string]interface{})*Model
	SignWithFacebook(token string)(*AuthResponse,error)
}



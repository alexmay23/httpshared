package user

import (
	"net/http"
	"strings"
	"github.com/alexmay23/httputils"

	"github.com/alexmay23/httpshared/shared"
)



func CreateUserMiddleware(useCase UseCase) shared.Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			tokenStr := ""
			if AuthHeader := req.Header.Get("Authorization"); AuthHeader != "" {
				tokenStr = strings.Split(AuthHeader, " ")[1]
			}
			if tokenStr == "" {
				token :=  httputils.GetValueFromURLInRequest(req, "token")
				if token != nil {
					tokenStr = *token
				}
			}
			if tokenStr == "" {
				httputils.HTTP401().Write(w)
				return
			}
			model := useCase.GetValidModelFromToken(tokenStr)
			if model ==  nil{
				httputils.HTTP401().Write(w)
				return
			}
			next.ServeHTTP(w, httputils.SetInContext(model, "CurrentUser", req))
		})
	}
}


func GetModelFromRequest(req *http.Request) *Model {
	user, ok := req.Context().Value("CurrentUser").(*Model)
	if !ok {
		panic("No current user")
	}
	return user
}





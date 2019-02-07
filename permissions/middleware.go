package permissions

import (
	"github.com/alexmay23/httpshared/shared"
	"github.com/alexmay23/httpshared/user"
	"github.com/alexmay23/httputils"
	"net/http"
)

type MiddlewareFactory =  func(permission string)shared.Middleware


func CreateMiddlewareFactory(useCase UseCase) MiddlewareFactory {
	return func(permission string)shared.Middleware {
		return func(next http.Handler) http.Handler {
			return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				usr := user.GetModelFromRequest(req)
				allowed := useCase.CheckUserPermission(usr.ID, permission)
				if !allowed{
					httputils.HTTP403().Write(w)
					return
				}
				next.ServeHTTP(w, req)
			})
		}
	}
}
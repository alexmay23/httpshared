package permissions

import (
	"github.com/alexmay23/httpshared/shared"
	"github.com/alexmay23/httpshared/user"
	"github.com/alexmay23/httputils"
	"net/http"
)

type Transport struct {
	useCase UseCase
	defaultMiddleWare shared.Middleware
	userMiddleware shared.Middleware
}



func NewTransport(useCase UseCase, defaultMiddleWare shared.Middleware, userMiddleware shared.Middleware) *Transport{
	return &Transport{useCase:useCase, defaultMiddleWare:defaultMiddleWare, userMiddleware:userMiddleware}
}




func (self *Transport)GetPermissionsHandler(w http.ResponseWriter, r *http.Request){
	usr := user.GetModelFromRequest(r)
	permissions := self.useCase.GetUserPermissions(usr.ID)
	httputils.JSON(w, map[string][]string{"permissions": permissions}, 200)
}
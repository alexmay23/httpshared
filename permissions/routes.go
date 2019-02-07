package permissions

import (
	"github.com/alexmay23/httputils"
	"net/http"
)

func (self *Transport) RegisterInRouter(router *httputils.Router) {
	router.Get("/permissions", self.defaultMiddleWare(self.userMiddleware(http.HandlerFunc(self.GetPermissionsHandler))))
}

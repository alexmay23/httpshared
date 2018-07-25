package user

import (
	"github.com/alexmay23/httputils"
	"net/http"
)

func (self *Transport) RegisterInRouter(router *httputils.Router) {
	middle := CreateUserMiddleware(self.useCase)
	router.Post("/verify", httputils.DefaultMiddlewares(http.HandlerFunc(self.VerifyHandler)))
	router.Post("/send", httputils.DefaultMiddlewares(http.HandlerFunc(self.SendHandler)))
	router.Post("/sign", httputils.DefaultMiddlewares(http.HandlerFunc(self.SignSocialHandler)))
	router.Put("/user", httputils.DefaultMiddlewares(middle(http.HandlerFunc(self.UpdateHandler))))
	router.Get("/user", httputils.DefaultMiddlewares(middle(http.HandlerFunc(self.GetCurrentUserHandler))))
}

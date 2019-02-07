package user

import (
	"github.com/alexmay23/httputils"
	"net/http"
)

func (self *Transport) RegisterInRouter(router *httputils.Router) {
	middle := CreateUserMiddleware(self.useCase)
	router.Post("/verify",  self.defaultMiddleWare(http.HandlerFunc(self.VerifyHandler)))
	router.Post("/send", self.defaultMiddleWare(http.HandlerFunc(self.SendHandler)))
	router.Post("/sign", self.defaultMiddleWare(http.HandlerFunc(self.SignSocialHandler)))
	router.Put("/user", self.defaultMiddleWare(middle(http.HandlerFunc(self.UpdateHandler))))
	router.Get("/user", self.defaultMiddleWare(middle(http.HandlerFunc(self.GetCurrentUserHandler))))
}

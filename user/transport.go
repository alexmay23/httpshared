package user

import (
	"net/http"
	"github.com/alexmay23/httputils"
	"bitbucket.org/alexmay_23/pp_backend/shared"
	"strconv"
)

type Transport struct {
	useCase UseCase
}



func NewTransport(useCase UseCase) *Transport{
	return &Transport{useCase:useCase}
}

func SendCodeValidatorMap()httputils.VMap{
	return httputils.VMap{
		"phone": httputils.RequiredStringValidators("phone", shared.PhoneValidation("phone")),
	}
}

func (self *Transport) SendHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, SendCodeValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	self.useCase.SendCode(body["phone"].(string))
	httputils.JSON(w, shared.OKJSON, 200)
}

func SignSocialValidatorMap()httputils.VMap{
	return httputils.VMap{
		"network": httputils.RequiredStringValidators("network", httputils.StringContainsValidator("network", []string{"fb"})),
		"token": httputils.RequiredStringValidators("token"),
	}
}


func (self *Transport) SignSocialHandler(w http.ResponseWriter, r *http.Request){

	body, err := httputils.GetValidatedBody(r, SignSocialValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	m, err := self.useCase.SignWithFacebook(body["token"].(string))
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	httputils.JSON(w, m, 200)
}

func VerifyCodeValidatorMap()httputils.VMap{
	return httputils.VMap{
		"phone": httputils.RequiredStringValidators("phone", shared.PhoneValidation("phone")),
		"code": httputils.RequiredStringValidators("code"),
	}
}

func (self *Transport) VerifyHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, VerifyCodeValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	code, err := strconv.ParseInt(body["code"].(string), 10, 32)
	if err != nil{
		httputils.HTTP400().Write(w)
		return
	}
	response, err := self.useCase.VerifyCode(body["phone"].(string), int(code))
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	httputils.JSON(w, response, 200)
}

func UpdateValidatorMap()httputils.VMap{
	return httputils.VMap{
		"name": httputils.RequiredStringValidators("name"),
	}
}


func (self *Transport) GetCurrentUserHandler(w http.ResponseWriter, r *http.Request){
	httputils.JSON(w, GetModelFromRequest(r), 200)
}

func (self *Transport) UpdateHandler(w http.ResponseWriter, r *http.Request){
	body, err := httputils.GetValidatedBody(r, UpdateValidatorMap())
	if err != nil{
		err.(httputils.ServerError).Write(w)
		return
	}
	m := GetModelFromRequest(r)
	m = self.useCase.Update(m, body)
	httputils.JSON(w, m, 200)
}
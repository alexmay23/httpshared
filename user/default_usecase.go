package user

import (
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"errors"
	"math/rand"
	"github.com/alexmay23/httputils"
	"context"
	"time"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"github.com/alexmay23/httpshared/shared"
)

type DefaultUseCase struct {
	repository Repository
	deliverer shared.SMSDeliverer
	config shared.Config
}




func (self *DefaultUseCase) SendCode(phone string) {
	code := rand.Intn(89999) + 10000
	m := self.repository.GetByPhone(phone)
	if  m == nil{
		m = self.repository.CreateWithPhone(phone, code)
	}else{
		m.Code = code
		self.repository.Update(m)
	}
	self.sendVerificationSMS(phone, code)
}


type FBResponse struct {
	Name string `json:"name"`
	ID   string `json:"id"`
}

func HTTP400(err error)httputils.ServerError{
	return httputils.ServerError{400, httputils.Errors{[]httputils.Error{httputils.UndefinedKeyError("INVALID_REQUEST", err.Error())}}}
}

func (self *DefaultUseCase) SignWithFacebook(token string)(*AuthResponse, error) {
	ctx, _ := context.WithTimeout(context.Background(), 5 *time.Second)
	url := fmt.Sprintf(
	"https://graph.facebook.com/v3.0/me?fields=name&access_token=%s", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil{
		return nil, HTTP400(err)
	}
	req.WithContext(ctx)
	var fbResponse FBResponse
	resp, err := http.DefaultClient.Do(req)
	if err != nil{
		return nil, HTTP400(err)
	}
	buf, _ := ioutil.ReadAll(resp.Body)
	err = json.Unmarshal(buf, &fbResponse)
	if err != nil{
		return nil, HTTP400(err)
	}
	m := self.repository.CreateWithFB(fbResponse.ID, fbResponse.Name,
		fmt.Sprintf("http://graph.facebook.com/%s/picture?type=square&width=200&height=200", fbResponse.ID), nil)
	jsonWebToken, err := self.generateTokenFromModel(m)
	if err != nil{
		panic(err)
	}
	response := AuthResponse{Token:jsonWebToken, User:*m}
	return &response, nil
}

func (self *DefaultUseCase) VerifyCode(phone string, code int)(*AuthResponse, error){
	m := self.repository.GetByPhone(phone)
	if m == nil{
		return nil, httputils.HTTP404(phone)
	}
	if m.Code != code{
		return nil, shared.NewServerError(403, "code", "invalid  code", "INVALID_CODE")
	}
	token, err := self.generateTokenFromModel(m)
	if err != nil{
		panic(err)
	}
	response := AuthResponse{Token:token, User:*m}
	return &response, nil
}


func (self *DefaultUseCase)Update(model *Model, data map[string]interface{})*Model{
	name := data["name"].(string)
	model.Name = name
	self.repository.Update(model)
	return model
}

func (self *DefaultUseCase) GetValidModelFromToken(token string) (*Model) {
	userData, err := self.getUserDataFromToken(token)
	if err != nil {
		return nil
	}
	model := self.repository.GetById(userData["user_id"].(string))
	if model == nil {
		return nil
	}
	if self.generateTokenHash(model) != userData["hash"] {
		return nil
	}
	return model
}

func (self *DefaultUseCase) getUserDataFromToken(token string)(map[string]interface{}, error) {
	tokenObj, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(self.config.GetValueForKey(shared.SecretKey)), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := tokenObj.Claims.(jwt.MapClaims)
	if ok && tokenObj.Valid {
		return claims, nil
	}
	return nil, errors.New("Failed jwt")
}

func (self *DefaultUseCase) generateTokenHash(model *Model) string {
	return self.config.GetValueForKey(shared.SecretKey) + model.Secret
}

func (self *DefaultUseCase)sendVerificationSMS(phone string, code int){
	self.deliverer.SendMessage(phone, fmt.Sprintf("Activation Code %d", code))
}



func (self *DefaultUseCase)generateTokenFromModel(model *Model) (string, error) {
	hash := self.generateTokenHash(model)
	claims := struct {
		UserID string `json:"user_id"`
		Hash   string `json:"hash"`
		jwt.StandardClaims
	}{
		model.ID,
		hash,
		jwt.StandardClaims{
			Issuer: "auth",
		},
	}
	tokenObj := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err := tokenObj.SignedString([]byte(self.config.GetValueForKey(shared.SecretKey)))
	if err != nil {
		return "", err
	}
	return token, nil
}

func NewDefaultUseCase(repository Repository, config shared.Config, deliverer shared.SMSDeliverer) *DefaultUseCase{
	if config.GetValueForKey(shared.SecretKey) == ""{
		panic("empty shared secret")
	}
	return &DefaultUseCase{repository:repository, config:config, deliverer:deliverer}
}


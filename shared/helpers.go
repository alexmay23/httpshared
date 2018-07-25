package shared

import (
	"gopkg.in/mgo.v2/bson"
	"github.com/alexmay23/httputils"
	"net/http"
)

var OKJSON = bson.M{"ok": 1}

type IDContainer struct {
	ID string `json:"id"`
}

func NewIDContainer(id string)*IDContainer{
	return &IDContainer{id}
}


var PendingState = "pending"
var ApprovedState = "approved"
var DeclinedState = "declined"
var StateMachine = []string{PendingState, ApprovedState, DeclinedState}


type Middleware func(http.Handler) http.Handler

func NewServerError(statusCode int, key string, description string, code string)httputils.ServerError{
	return httputils.ServerError{StatusCode: statusCode, Errors: httputils.Errors{Errors: []httputils.Error{
		{
			Key:         key,
			Description: description,
			Code:        code,
		}}}}
}




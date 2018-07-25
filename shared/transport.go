package shared

import "github.com/alexmay23/httputils"

type TransportRegisterer interface {
	RegisterInRouter(router *httputils.Router)
}

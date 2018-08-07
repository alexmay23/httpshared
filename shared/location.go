package shared

import (
	"encoding/json"
	"github.com/alexmay23/httputils"
	"strconv"
	"bitbucket.org/alexmay_23/true_fund_server/shared"
	"net/http"
)

type LocationParameters struct {
	Latitude float64
	Longitude float64
	MinDistance int64
	MaxDistance int64
}


func NewObjectLocation(latitude , longitude float64)*ObjectLocation{
	return &ObjectLocation{"Point", []float64{longitude, latitude}}
}

type ObjectLocation struct {
	Type        string
	Coordinates []float64
}

func (self ObjectLocation) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	}{
		Latitude:  self.Coordinates[1],
		Longitude: self.Coordinates[0],
	})
}


func DistanceValidator() httputils.Validator {
	bottom := 0
	upper := 32000000000
	return httputils.IntInRangeValidator("max_distance", httputils.IntRange{Bottom: &bottom, Upper: &upper})
}



func ParseGeoLocation(longitudeRaw string, latitudeRaw string)(float64, float64, error){
	longitude, err := strconv.ParseFloat(longitudeRaw, 64)
	if err != nil{
		return 0, 0, shared.NewServerError(400, "longitude", "INVALID_FLOAT", "INVALID_FLOAT")
	}
	latitude, err := strconv.ParseFloat(latitudeRaw, 64)
	if err != nil{
		return 0, 0, shared.NewServerError(400, "longitude", "INVALID_FLOAT", "INVALID_FLOAT")
	}

	errs := httputils.ValidateValue(longitude, shared.LongitudeValidators("longitude"))
	if errs != nil{
		return 0, 0, httputils.ServerError{StatusCode: 400, Errors: httputils.Errors{Errors: errs}}
	}
	errs = httputils.ValidateValue(latitude, shared.LongitudeValidators("latitude"))
	if errs != nil{
		return 0, 0, httputils.ServerError{StatusCode: 400, Errors: httputils.Errors{Errors: errs}}
	}
	return longitude, latitude, nil
}

func LocationParametersFromRequest(req *http.Request, defaultMaxDistance int64)(*LocationParameters, error){
	longitudeRaw := httputils.GetValueFromURLInRequest(req, "longitude")
	latitudeRaw := httputils.GetValueFromURLInRequest(req, "latitude")
	if longitudeRaw == nil || latitudeRaw == nil{
		return nil, nil
	}

	longitude, latitude, err :=  ParseGeoLocation(*longitudeRaw, *latitudeRaw)

	var minDistance int64  = 0
	maxDistance := defaultMaxDistance

	maxDistanceRaw := httputils.GetValueFromURLInRequest(req, "max_distance")
	minDistanceRaw := httputils.GetValueFromURLInRequest(req, "min_distance")

	if maxDistanceRaw != nil{

		maxDistance, err = strconv.ParseInt(*maxDistanceRaw, 10, 64)
		if err != nil{
			return nil, shared.NewServerError(400, "max_distance", "INVALID_INT", "INVALID_INT")
		}

		errs := httputils.ValidateValue(maxDistanceRaw, httputils.RequiredIntValidators("max_distance", DistanceValidator()))
		if errs != nil{
			return nil, httputils.ServerError{StatusCode: 400, Errors: httputils.Errors{Errors: errs}}
		}
	}

	if minDistanceRaw != nil{

		minDistance, err = strconv.ParseInt(*minDistanceRaw, 10, 64)
		if err != nil{
			return nil, shared.NewServerError(400, "min_distance", "INVALID_FLOAT", "INVALID_FLOAT")
		}

		errs := httputils.ValidateValue(maxDistanceRaw, httputils.RequiredIntValidators("min_distance", DistanceValidator()))
		if errs != nil{
			return nil, httputils.ServerError{StatusCode: 400, Errors: httputils.Errors{Errors: errs}}
		}
	}

	if minDistance > maxDistance{
		return nil, shared.NewServerError(400, "min_distance", "Min distance more tham max", "MIN_MAX_ERROR")
	}

	return &LocationParameters{Latitude:latitude, Longitude:longitude, MinDistance:minDistance, MaxDistance:maxDistance}, nil

}


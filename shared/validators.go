package shared

import (
	"github.com/alexmay23/httputils"
	"github.com/ttacon/libphonenumber"
)

func PhoneValidation(key string) httputils.Validator {
	return func(value interface{}) error {
		stringValue := value.(string)
		_, err := libphonenumber.Parse(stringValue, "")
		if err != nil {
			return httputils.Error{Key: key, Description: "Invalid phone", Code: "INVALID_PHONE_ERROR"}
		}
		return nil
	}
}

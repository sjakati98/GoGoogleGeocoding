package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

const (
	// GeocodeAddressResponseValidateErrorPrefix ...
	GeocodeAddressResponseValidateErrorPrefix = "GeocodeAddressResponse.Validate.error: %v"
)

// GeocodeAddressResponse defines the response type for geocoding
type GeocodeAddressResponse struct {
	Latitude  float64 `validate:"required,max=90,min=-90" json:"latitude"`
	Longitude float64 `validate:"required,max=180,min=-180" json:"longitude"`
}

// Validate validates the response struct
func (g GeocodeAddressResponse) Validate(val validator.Validate) error {
	// perform validations over struct
	if err := val.Struct(g); err != nil {
		logrus.Errorf("geocode address response struct did not pass validations: %v\n", err)
		return fmt.Errorf(GeocodeAddressResponseValidateErrorPrefix, err)
	}
	return nil
}

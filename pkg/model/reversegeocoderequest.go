package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

const (
	// ReverseGeocodeRequestValidateErrorPrefix ...
	ReverseGeocodeRequestValidateErrorPrefix = "ReverseGeocodeRequest.Validate.error: %v"
)

// ReverseGeocodeRequest defines the response type for geocoding
type ReverseGeocodeRequest struct {
	Latitude  float64 `validate:"required,max=90,min=-90" json:"latitude"`
	Longitude float64 `validate:"required,max=180,min=-180" json:"longitude"`
}

// Validate validates the response struct
func (g ReverseGeocodeRequest) Validate(val validator.Validate) error {
	// perform validations over struct
	if err := val.Struct(g); err != nil {
		logrus.Errorf("geocode address response struct did not pass validations: %v\n", err)
		return fmt.Errorf(ReverseGeocodeRequestValidateErrorPrefix, err)
	}
	return nil
}

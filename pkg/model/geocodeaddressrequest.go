package model

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

const (
	// GeocodeAddressRequestValidateErrorPrefix ...
	GeocodeAddressRequestValidateErrorPrefix = "GeocodeAddressRequest.Validate.error: %v"
)

// GeocodeAddressRequest defines the request message
// for geocoding
type GeocodeAddressRequest struct {
	AddressLine1 *string `json:"addressLine1"`
	AddressLine2 *string `json:"addressLine2"`
	City         *string `json:"city"`
	State        *string `json:"state"`
	Zipcode      *string `json:"zipcode"`
}

// Validate validates the request struct
func (g GeocodeAddressRequest) Validate(val validator.Validate) error {
	if g.AddressLine1 == nil &&
		g.AddressLine2 == nil &&
		g.City == nil &&
		g.State == nil &&
		g.Zipcode == nil {
		logrus.Errorf(GeocodeAddressRequestValidateErrorPrefix, "request empty")
		return fmt.Errorf(GeocodeAddressRequestValidateErrorPrefix, "request empty")
	}
	return nil
}

package configs

import (
	"fmt"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
)

// GeocodingClientConfig ...
type GeocodingClientConfig struct {
	APIKey     string `validate:"required" env:"GEOCODING_SERVICE_API_KEY"`
	APIBaseURL string `validate:"required,url" env:"GEOCODING_SERVICE_BASE_URL"`
}

const (
	// GeocodingClientConfigErrorPrefix ...
	GeocodingClientConfigErrorPrefix = "GeocodingClientConfig.Validate.error: %v"
)

// Validate validates the configuration struct
func (g GeocodingClientConfig) Validate(val validator.Validate) error {

	// perform validations over struct
	if err := val.Struct(g); err != nil {
		logrus.Errorf("geocode client config struct did not pass validations: %v\n", err)
		return fmt.Errorf(GeocodingClientConfigErrorPrefix, err)
	}
	return nil
}

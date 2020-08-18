package geocodingclient

import (
	"context"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"github.com/sjakati98/GoGoogleGeocoding/configs"
	"github.com/sjakati98/GoGoogleGeocoding/pkg/model"
)

// GeocodingClientInterface defines the necessary contract for
// a geocoding client
type GeocodingClientInterface interface {
	GeocodeAddress(
		ctx context.Context,
		request model.GeocodeAddressRequest,
	) (*model.ReverseGeocodeRequest, error)
	ReverseGeocode(
		ctx context.Context,
		request model.ReverseGeocodeRequest,
	) (*model.ReverseGeocodeResponse, error)
}

// GeocodingClient defines the explicit client
// to be used for geocoding
type GeocodingClient struct {
	APIKey     string
	APIBaseURL string
	HTTPClient http.Client
	Validator  validator.Validate
}

// GetGeocodingClient returns a new geocoding client
func GetGeocodingClient(config configs.GeocodingClientConfig) (*GeocodingClient, error) {

	// validate client config
	val := validator.New()
	err := config.Validate(*val)
	if err != nil {
		logrus.Errorf("error validating geocoding client config: %v", err)
		return nil, err
	}

	client := GeocodingClient{
		APIBaseURL: config.APIBaseURL,
		APIKey:     config.APIKey,
		HTTPClient: *http.DefaultClient,
		Validator:  *val,
	}

	logrus.Infof("successfully created geocoding client")

	return &client, nil
}

package geocodingclient

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/sjakati98/GoGoogleGeocoding/pkg/model"
)

// GeocodeAddress implements the geocoding functionality
func (g GeocodingClient) GeocodeAddress(
	ctx context.Context,
	request model.GeocodeAddressRequest,
) (*model.GeocodeAddressResponse, error) {

	var requestID = uuid.New().String()

	// validate request
	err := request.Validate(g.Validator)
	if err != nil {
		logrus.Errorf("error handling geocding address request: %v, for request: %s", err, requestID)
		return nil, err
	}
	logrus.Infof("validated geocode address request for request: %s", requestID)

	// gather address
	formattedAddress := ""

	if request.AddressLine1 != nil {
		formattedAddress = formattedAddress + *request.AddressLine1 + ", "
	}
	if request.AddressLine2 != nil {
		formattedAddress = formattedAddress + *request.AddressLine2 + ", "
	}
	if request.City != nil {
		formattedAddress = formattedAddress + *request.City + ", "
	}
	if request.State != nil {
		formattedAddress = formattedAddress + *request.State + ", "
	}
	if request.Zipcode != nil {
		formattedAddress = formattedAddress + *request.Zipcode + ", "
	}

	// get request for geocoding service
	req, err := g.getGeocodeRequest(formattedAddress)
	if err != nil {
		logrus.Errorf("error creating geocode address HTTP request request: %v, for request: %s", err, requestID)
		return nil, err
	}
	logrus.Infof("created geocode address HTTP request for request: %s", requestID)

	// make call to geocoding service
	resp, err := g.HTTPClient.Do(req)
	if err != nil {
		logrus.Errorf("error doing request to geocoding service: %v, for request: %s", err, requestID)
		return nil, err
	}
	logrus.Infof("successfully did request to geocoding service for request: %s", requestID)

	// marshal response
	var geocodeResponse GeocodingResponse
	defer resp.Body.Close()

	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logrus.Errorf("error getting response: %v", err)
	}

	err = json.Unmarshal(bodyBytes, &geocodeResponse)
	if err != nil {
		logrus.Errorf("error decoding response from geocoding service: %v, for request: %s", err, requestID)
		return nil, err
	}

	// return
	return &model.GeocodeAddressResponse{
		Latitude:  geocodeResponse.Results[0].Geometry.Location.Lat,
		Longitude: geocodeResponse.Results[0].Geometry.Location.Lng,
	}, nil
}

// getGeocodeRequest gets the formatted  request
func (g GeocodingClient) getGeocodeRequest(formattedAddress string) (*http.Request, error) {
	req, err := http.NewRequest("GET", g.APIBaseURL, nil)
	if err != nil {
		logrus.Errorf("error creating get request: %v", err)
		return nil, err
	}

	q := req.URL.Query()
	q.Add("key", g.APIKey)
	q.Add("address", formattedAddress)
	req.URL.RawQuery = q.Encode()

	return req, nil
}

package geocodingclient

// GeocodingResponse defines the result of an API
// call to the Google geocoding service
type GeocodingResponse struct {
	Results []GeocodingResult `json:"results"`
	Status  string            `json:"status"`
}

// GeocodingResult ...
type GeocodingResult struct {
	AddressComponents []GeocodingAddressComponents `json:"address_components"`
	FormattedAddress  string                       `json:"formatted_address"`
	Geometry          GeocodingGeometry            `json:"geometry"`
	PlaceID           string                       `json:"place_id"`
	Types             []string                     `json:"types"`
}

// GeocodingAddressComponents ...
type GeocodingAddressComponents struct {
	LongName  string   `json:"long_name"`
	ShortName string   `json:"short_name"`
	Types     []string `json:"types"`
}

// GeocodingGeometry ...
type GeocodingGeometry struct {
	Bounds       Bounds   `json:"bounds"`
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

// Location ...
type Location struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Northeast ...
type Northeast struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Southwest ...
type Southwest struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}

// Viewport ...
type Viewport struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}

// Geometry ...
type Geometry struct {
	Location     Location `json:"location"`
	LocationType string   `json:"location_type"`
	Viewport     Viewport `json:"viewport"`
}

// Bounds ...
type Bounds struct {
	Northeast Northeast `json:"northeast"`
	Southwest Southwest `json:"southwest"`
}

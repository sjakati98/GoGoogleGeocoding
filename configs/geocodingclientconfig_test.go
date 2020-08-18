package configs

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestGeocodingClientConfig_Validate(t *testing.T) {
	type fields struct {
		APIKey     string
		APIBaseURL string
	}
	type args struct {
		val validator.Validate
	}

	// create validator for mocks
	val := validator.New()

	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name:   "Test empty",
			fields: fields{},
			args: args{
				*val,
			},
			wantErr: true,
		},
		{
			name: "Test invalid base url",
			fields: fields{
				APIKey:     "random key string",
				APIBaseURL: "not a url",
			},
			args: args{
				*val,
			},
			wantErr: true,
		},
		{
			name: "Test valid",
			fields: fields{
				APIKey:     "random key string",
				APIBaseURL: "https://maps.googleapis.com/maps/api/geocode/json",
			},
			args: args{
				*val,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GeocodingClientConfig{
				APIKey:     tt.fields.APIKey,
				APIBaseURL: tt.fields.APIBaseURL,
			}
			if err := g.Validate(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("GeocodingClientConfig.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

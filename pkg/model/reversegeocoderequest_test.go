package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestReverseGeocodeRequest_Validate(t *testing.T) {
	type fields struct {
		Latitude  float64
		Longitude float64
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
			name: "Test valid",
			fields: fields{
				Latitude:  -34.444,
				Longitude: 101.343,
			},
			args: args{
				*val,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := ReverseGeocodeRequest{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if err := g.Validate(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("ReverseGeocodeRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

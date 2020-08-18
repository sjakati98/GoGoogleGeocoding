package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestGeocodeAddressResponse_Validate(t *testing.T) {
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
				Latitude:  -43.555,
				Longitude: 98.32423,
			},
			args: args{
				*val,
			},
			wantErr: false,
		},
		{
			name: "Test invalid latitude",
			fields: fields{
				Latitude:  -100.555,
				Longitude: 98.32423,
			},
			args: args{
				*val,
			},
			wantErr: true,
		},
		{
			name: "Test invalid longitude",
			fields: fields{
				Latitude:  -43.555,
				Longitude: 300.32423,
			},
			args: args{
				*val,
			},
			wantErr: true,
		},
		{
			name: "Test invalid",
			fields: fields{
				Latitude:  -100.555,
				Longitude: 300.32423,
			},
			args: args{
				*val,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GeocodeAddressResponse{
				Latitude:  tt.fields.Latitude,
				Longitude: tt.fields.Longitude,
			}
			if err := g.Validate(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("GeocodeAddressResponse.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

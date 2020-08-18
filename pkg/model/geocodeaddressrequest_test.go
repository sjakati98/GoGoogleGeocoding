package model

import (
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestGeocodeAddressRequest_Validate(t *testing.T) {
	type fields struct {
		AddressLine1 *string
		AddressLine2 *string
		City         *string
		State        *string
		Zipcode      *string
	}
	type args struct {
		val validator.Validate
	}

	// create validator for mocks
	val := validator.New()

	randomString := "random string"

	mockAddress := fields{
		AddressLine1: &randomString,
		City:         &randomString,
		State:        &randomString,
		Zipcode:      &randomString,
	}

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
			name:   "Test valid",
			fields: mockAddress,
			args: args{
				*val,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := GeocodeAddressRequest{
				AddressLine1: tt.fields.AddressLine1,
				AddressLine2: tt.fields.AddressLine2,
				City:         tt.fields.City,
				State:        tt.fields.State,
				Zipcode:      tt.fields.Zipcode,
			}
			if err := g.Validate(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("GeocodeAddressRequest.Validate() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

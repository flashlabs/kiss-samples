package main_test

import (
	"testing"

	"github.com/flashlabs/kiss-samples/payloadschema"
)

func TestSchema(t *testing.T) {
	type args struct {
		payload map[string]any
		schema  string
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "valid",
			args: args{
				payload: map[string]any{
					"id": "bar",
				},
				schema: "schema/payload.json",
			},
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				payload: map[string]any{
					"name": "bar",
				},
				schema: "schema/payload.json",
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := main.ValidateSchema(tt.args.payload, tt.args.schema); (err != nil) != tt.wantErr {
				t.Errorf("ValidateSchema() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

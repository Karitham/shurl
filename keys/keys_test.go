package keys

import (
	"encoding/base32"
	"testing"
)

func TestGenerator_Generate(t *testing.T) {
	type fields struct {
		length int
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		{
			name: "basic",
			fields: fields{
				length: 8,
			},
			wantErr: false,
		},
		{
			name: "short",
			fields: fields{
				length: 4,
			},
			wantErr: false,
		},
	}
	errEncoder :=
		base32.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZ234567").WithPadding(base32.NoPadding)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Generator{
				length: tt.fields.length,
			}
			got, err := g.Generate()
			if (err != nil) != tt.wantErr {
				t.Errorf("Generator.Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) == tt.fields.length {
				return
			}
			t.Error("got a weird result: ", errEncoder.EncodeToString(got))
		})
	}
}

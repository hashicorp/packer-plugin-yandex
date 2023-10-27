package yandex

import (
	"errors"
	"os"
	"testing"

	"github.com/hashicorp/packer-plugin-sdk/template/interpolate"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAccessConfig_Prepare(t *testing.T) {
	bytes, err := os.ReadFile(TestServiceAccountKeyFile)
	require.NoErrorf(t, err, "failed to read file %s", TestServiceAccountKeyFile)

	var TestServiceAccountKeyFileContent = string(bytes)

	type fields struct {
		Endpoint              string
		ServiceAccountKeyFile string
		Token                 string
		MaxRetries            int
		saKeyType             keyType
	}
	tests := []struct {
		name   string
		fields fields
		want   []error
	}{
		{
			name: "sa_key_as_file", fields: fields{
				Endpoint:              "",
				ServiceAccountKeyFile: TestServiceAccountKeyFile,
				Token:                 "",
			},
			want: nil,
		},
		{
			name: "sa_key_as_json_content", fields: fields{
				ServiceAccountKeyFile: TestServiceAccountKeyFileContent,
				Token:                 "",
			},
			want: nil,
		},
		{
			name: "both_identities", fields: fields{
				ServiceAccountKeyFile: TestServiceAccountKeyFileContent,
				Token:                 "t1.super-token",
			},
			want: []error{errors.New("one of token or service account key file must be specified, not both")},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := &interpolate.Context{}
			c := &AccessConfig{
				Endpoint:              tt.fields.Endpoint,
				ServiceAccountKeyFile: tt.fields.ServiceAccountKeyFile,
				Token:                 tt.fields.Token,
				MaxRetries:            tt.fields.MaxRetries,
				saKeyType:             tt.fields.saKeyType,
			}
			assert.Equalf(t, tt.want, c.Prepare(ctx), "Prepare(%v)")
		})
	}
}

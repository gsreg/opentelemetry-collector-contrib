package eginnovations

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"go.opentelemetry.io/collector/config/configcompression"
	"go.opentelemetry.io/collector/config/configgrpc"
	"go.opentelemetry.io/collector/config/configopaque"
	"go.opentelemetry.io/collector/config/configtls"
)

func TestValidate(t *testing.T) {

	tests := []struct {
		name string
		cfg  *Config
		err  string
	}{
		{
			name: "no endpoint",
			cfg:  &Config{},
			err:  "endpoint not specified, please fix the configuration",
		},
		{
			name: "TLS settings are valid",
			cfg: &Config{
				ClientConfig: configgrpc.ClientConfig{
					Endpoint: "eginnovations.com",
					TLSSetting: configtls.ClientConfig{
						InsecureSkipVerify: true,
					},
				},
			},
		},
		{
			name: "With configgrpc client configs",
			cfg: &Config{
				UserID: "user1",
				Token:  "xxxxxxxxx",
				ClientConfig: configgrpc.ClientConfig{
					Endpoint:    "eginnovations.com",
					Compression: configcompression.TypeGzip,
					TLSSetting: configtls.ClientConfig{
						Config:             configtls.Config{},
						Insecure:           false,
						InsecureSkipVerify: false,
						ServerName:         "",
					},
					ReadBufferSize:  0,
					WriteBufferSize: 0,
					WaitForReady:    false,
					Headers: map[string]configopaque.String{
						"hdr1": "value1",
					},
					BalancerName: "",
				},
			},
		},
	}
	for _, testInstance := range tests {
		t.Run(testInstance.name, func(t *testing.T) {
			err := testInstance.cfg.Validate()
			if testInstance.err != "" {
				assert.EqualError(t, err, testInstance.err)
			} else {
				assert.NoError(t, err)
			}
		})
	}

}

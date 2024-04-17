package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetRmsClient(t *testing.T) {
	endpointConfig = map[string]string{
		"rms": "https://rms.myhuaweicloud.com",
	}
	client := getRMSClient()
	assert.NotNil(t, client)
}

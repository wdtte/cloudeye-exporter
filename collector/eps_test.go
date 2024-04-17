package collector

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEPSClient(t *testing.T) {
	endpointConfig = map[string]string{
		"eps": "https://eps.myhuaweicloud.com",
	}
	epsClient := getEPSClient()
	assert.NotNil(t, epsClient)
}

package collector

import (
	"errors"
	"os"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/huaweicloud/cloudeye-exporter/logs"
)

func TestName(t *testing.T) {
	err := InitMetricConf()
	assert.Equal(t, true, err != nil)
	metricConf = map[string]MetricConf{
		"SYS.ECS": {
			Resource: "rms",
			DimMetricName: map[string][]string{
				"instance_id": []string{"cpu_util", "mem_util", "disk_util_inband"},
			},
		},
	}
	mconf := getMetricConfigMap("TEST.ECS")
	assert.Equal(t, true, mconf == nil)

	mconf = getMetricConfigMap("SYS.ECS")
	assert.Equal(t, 1, len(mconf))
}

func TestInitConfigSecurityModIsTrue(t *testing.T) {
	SecurityMod = true
	TmpSK = "tmpSK"
	TmpAK = "tmpAK"
	CloudConf.Auth.ProjectID = "testProjectId"
	CloudConf.Auth.ProjectName = "testProjectName"
	err := InitConfig()
	if err != nil {
		return
	}

	assert.Equal(t, TmpAK, conf.AccessKey)
	assert.Equal(t, TmpSK, conf.SecretKey)
}

func TestInitConfigSecurityModIsFalse(t *testing.T) {
	SecurityMod = false
	CloudConf.Auth.AccessKey = "tmpSK"
	CloudConf.Auth.SecretKey = "tmpAK"
	CloudConf.Auth.ProjectID = "testProjectId"
	CloudConf.Auth.ProjectName = "testProjectName"
	err := InitConfig()
	if err != nil {
		return
	}

	assert.Equal(t, CloudConf.Auth.AccessKey, conf.AccessKey)
	assert.Equal(t, CloudConf.Auth.SecretKey, conf.SecretKey)
}

func TestInitEndpointConfig(t *testing.T) {
	testCases := []struct {
		name    string
		patches func() *gomonkey.Patches
		expect  func(t *testing.T)
	}{
		{
			"ReadConfigFileError",
			func() *gomonkey.Patches {
				patches := gomonkey.NewPatches()
				patches.ApplyFuncReturn(os.ReadFile, nil, errors.New("read file error"))
				return patches
			},
			func(t *testing.T) {
				assert.NotNil(t, &endpointConfig)
			},
		},
		{
			"UnMarshalConfigError",
			func() *gomonkey.Patches {
				patches := gomonkey.NewPatches()
				patches.ApplyFuncReturn(os.ReadFile, []byte("\"rms\":\r\n  \"https://rms.xxx.xxx.com\"\r\n\"eps\":\r\n  \"https://eps.xxx.xxx.com\""), nil)
				patches.ApplyFuncReturn(yaml.Unmarshal, errors.New("unmarshal yaml error"))
				return patches
			},
			func(t *testing.T) {
				assert.NotNil(t, &endpointConfig)
			},
		},
		{
			"Success",
			func() *gomonkey.Patches {
				patches := gomonkey.NewPatches()
				patches.ApplyFuncReturn(os.ReadFile, []byte("\"rms\":\r\n  \"https://rms.xxx.xxx.com\"\r\n\"eps\":\r\n  \"https://eps.xxx.xxx.com\""), nil)
				return patches
			},
			func(t *testing.T) {
				assert.NotNil(t, &endpointConfig)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			patches := testCase.patches()
			patches.ApplyMethod(&logs.Logger, "Errorf", func(logger *logs.LoggerConstructor, template string, args ...interface{}) {})
			defer patches.Reset()
			initEndpointConfig()
			testCase.expect(t)
		})
	}
}

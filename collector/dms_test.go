package collector

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
)

func TestDmsGetResourceInfo(t *testing.T) {
	sysConfig := map[string][]string{"instance_id": {"cpu_util"}}
	patches := getPatches()
	patches = gomonkey.ApplyFuncReturn(getMetricConfigMap, sysConfig)
	patches.ApplyFuncReturn(listResources, mockRmsResource(), nil)
	defer patches.Reset()

	var dmsInfo DMSInfo
	labels, metrics := dmsInfo.GetResourceInfo()
	assert.Equal(t, 1, len(labels))
	assert.Equal(t, 0, len(metrics))
}

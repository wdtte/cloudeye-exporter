package collector

import (
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
	rmmModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
	"github.com/stretchr/testify/assert"
)

func TestEvsGetResourceInfo(t *testing.T) {
	var evsgetter EVSInfo
	patches := getPatches()
	sysConfig := map[string][]string{"disk_name": {"disk_device_read_bytes_rate"}}
	patches = gomonkey.ApplyFuncReturn(getResourceFromRMS, true)
	patches = patches.ApplyFuncReturn(getMetricConfigMap, sysConfig)
	volumes := mockRmsResource()
	volumes[0].Properties = map[string]interface{}{
		"attachments": []Attachment{
			{Device: "vda", ServerId: "0001-0001-00000001"},
		},
	}
	id := "xxxx2"
	name := "resource2"
	epId := "0"
	volumes1 := rmmModel.ResourceEntity{
		Id:   &id,
		Name: &name,
		EpId: &epId,
		Properties: map[string]interface{}{
			"attachments": []Attachment{
				{Device: "vdb", ServerId: "0002-0002-00000002"},
			},
		},
	}
	volumes = append(volumes, volumes1)
	patches.ApplyFuncReturn(listResources, volumes, nil)

	metrics := []model.MetricInfoList{
		{
			Namespace:  "SYS.ECS",
			MetricName: "disk_device_io_iops_qos_num",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "disk_name",
					Value: "0001-0001-00000001-vda",
				},
			},
		},
		{
			Namespace:  "SYS.ECS",
			MetricName: "disk_device_io_iops_qos_num",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "disk_name",
					Value: "0002-0002-00000002-volume-xxxx2",
				},
			},
		},
	}
	patches.ApplyFuncReturn(listAllMetrics, metrics, nil)
	defer patches.Reset()

	labels, metrics := evsgetter.GetResourceInfo()
	assert.Equal(t, 2, len(labels))
	assert.Equal(t, 2, len(metrics))
}

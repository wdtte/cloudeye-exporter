package collector

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/ces/v1/model"
	rmsModel "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
)

func TestGetVPNResourceInfo(t *testing.T) {
	metrics := []model.MetricInfoList{
		{
			Namespace:  "SYS.VPN",
			MetricName: "tunnel_average_latency",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "evpn_connection_id",
					Value: "0001-0001-00000001",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "gateway_recv_pkt_rate",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "evpn_gateway_id",
					Value: "0002-0002-00000002",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "gateway_connection_num",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "p2c_vpn_gateway_id",
					Value: "0001-0001-00000003",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "connection_status",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "vgw_ipsec_connect_id",
					Value: "0002-0002-00000004",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "private_average_latency",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "vpn_connection_id",
					Value: "0001-0001-00000005",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "sa_send_pkt_rate",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "evpn_connection_id",
					Value: "0001-0001-00000001",
				},
				{
					Name:  "evpn_sa_id",
					Value: "0002-0002-00000006",
				},
			},
		},
		{
			Namespace:  "SYS.VPN",
			MetricName: "sa_send_pkt_rate",
			Unit:       "count/op",
			Dimensions: []model.MetricsDimension{
				{
					Name:  "evpn_gateway_id",
					Value: "0002-0002-00000007",
				},
			},
		},
	}

	id1 := "0001-0001-00000001"
	name1 := "evpn_connect"
	epId := "0"
	id2 := "0002-0002-00000002"
	name2 := "evpn_gateway"
	id3 := "0001-0001-00000003"
	name3 := "p2c_vpn_gateway_id"
	id4 := "0002-0002-00000004"
	name4 := "vgw_ipsec_connect_id"
	id5 := "0001-0001-00000005"
	name5 := "vpn_connection_id"
	id6 := "0002-0002-00000006"
	name6 := "evpn_sa_id"
	id7 := "0002-0002-00000007"
	name7 := "evpn_gateway"
	resp := []rmsModel.ResourceEntity{
		{Id: &id1, Name: &name1, EpId: &epId},
		{Id: &id2, Name: &name2, EpId: &epId},
		{Id: &id3, Name: &name3, EpId: &epId},
		{Id: &id4, Name: &name4, EpId: &epId},
		{Id: &id5, Name: &name5, EpId: &epId},
		{Id: &id6, Name: &name6, EpId: &epId},
		{Id: &id7, Name: &name7, EpId: &epId},
	}

	patches := gomonkey.ApplyFuncReturn(listAllMetrics, metrics, nil)
	patches = patches.ApplyFuncReturn(listResources, resp, nil)

	defer patches.Reset()

	var vpnGetter VPNInfo
	labels, metrics := vpnGetter.GetResourceInfo()
	assert.Equal(t, 7, len(labels))

}

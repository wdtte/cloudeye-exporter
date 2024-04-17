package collector

import (
	http_client "github.com/huaweicloud/huaweicloud-sdk-go-v3/core"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/auth/global"
	v1 "github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/model"
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/services/rms/v1/region"
)

func getRMSClient() *v1.RmsClient {
	return v1.NewRmsClient(getRMSClientBuilder().Build())
}

func getRMSClientBuilder() *http_client.HcHttpClientBuilder {
	builder := v1.RmsClientBuilder().WithCredential(global.NewCredentialsBuilder().WithAk(conf.AccessKey).WithSk(conf.SecretKey).WithDomainId(conf.DomainID).Build())
	if endpoint, ok := endpointConfig["rms"]; ok {
		builder.WithEndpoint(endpoint)
	} else {
		builder.WithRegion(region.ValueOf("cn-north-4"))
	}
	return builder
}

func listResources(provider, resourceType string) ([]model.ResourceEntity, error) {
	limit := int32(200)
	req := &model.ListResourcesRequest{
		Provider: provider,
		Type:     resourceType,
		RegionId: &conf.Region,
		Limit:    &limit,
	}
	var resources []model.ResourceEntity
	for {
		response, err := getRMSClient().ListResources(req)
		if err != nil {
			return resources, err
		}
		resources = append(resources, *response.Resources...)
		if response.PageInfo.NextMarker == nil {
			break
		}
		req.Marker = response.PageInfo.NextMarker
	}
	return resources, nil
}

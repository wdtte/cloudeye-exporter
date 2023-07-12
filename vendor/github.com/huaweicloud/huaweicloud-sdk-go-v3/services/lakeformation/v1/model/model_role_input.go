package model

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-v3/core/utils"

	"strings"
)

// role信息
type RoleInput struct {

	// role名字
	RoleName string `json:"role_name"`

	// 描述信息
	Description *string `json:"description,omitempty"`
}

func (o RoleInput) String() string {
	data, err := utils.Marshal(o)
	if err != nil {
		return "RoleInput struct{}"
	}

	return strings.Join([]string{"RoleInput", string(data)}, " ")
}

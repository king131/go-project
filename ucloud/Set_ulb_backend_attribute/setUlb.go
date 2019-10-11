package main

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"os"
	"strconv"
)

const region = "cn-sh2"
const projectId =  "org-xxxxxx"

func loadConfig() (*ucloud.Config, *auth.Credential) {
	cfg := ucloud.NewConfig()
	cfg.Region = region
	cfg.ProjectId = projectId



	credential := auth.NewCredential()
	credential.PrivateKey ="xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	credential.PublicKey = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"

	return &cfg, &credential
}

func main() {
	cfg, credential := loadConfig()
	ulbClient := ulb.NewClient(cfg, credential)

	backend0 := ulbClient.NewUpdateBackendAttributeRequest()
	//获取ulb id
	backend0.ULBId = ucloud.String(os.Args[1])
	//获取backend id
	backend0.BackendId = ucloud.String(os.Args[2])
	//获取backend 状态值,0为禁止，1为启用
	statusValue , err := strconv.Atoi(os.Args[3])
	if err != nil {
		log.Error(err)
	}
	backend0.Enabled = ucloud.Int(statusValue)
	//设置backend状态
	_ , err = ulbClient.UpdateBackendAttribute(backend0)
	if err != nil {
		log.Error(err)
	} else {
		switch statusValue {
		case 1:
			fmt.Printf("%s 上线成功！", os.Args[2])
		case 0:
			fmt.Printf("%s 下线成功！", os.Args[2])
		}
	}


}

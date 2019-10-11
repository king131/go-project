package main

import (
	"fmt"
	"github.com/ucloud/ucloud-sdk-go/services/ulb"
	"github.com/ucloud/ucloud-sdk-go/ucloud"
	"github.com/ucloud/ucloud-sdk-go/ucloud/auth"
	"github.com/ucloud/ucloud-sdk-go/ucloud/log"
	"os"
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

	req := ulbClient.NewDescribeULBRequest()
	req.ULBId = ucloud.String(os.Args[1])

	data , err := ulbClient.DescribeULB(req)
	if err != nil {
		log.Error(err)
	}
	//循环Vserver
	for i :=0 ; i < len(data.DataSet[0].VServerSet); i++ {
		fmt.Printf("VServer Name:%s \n",data.DataSet[0].VServerSet[i].VServerName)
		//循环backendid和后端服务器名称
		for j := 0; j < len(data.DataSet[0].VServerSet[i].BackendSet); j++ {
			fmt.Printf("    BackendID: %s 后端服务器: %s\n",data.DataSet[0].VServerSet[i].BackendSet[j].BackendId,data.DataSet[0].VServerSet[i].BackendSet[j].ResourceName)
		}
	}

}

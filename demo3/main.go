package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	"io/ioutil"
	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset      *kubernetes.Clientset
		deployYaml     []byte
		deployJson     []byte
		deployment     = apps_v1.Deployment{}
		containers     []v1.Container
		nginxContainer v1.Container
		err            error
	)

	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		goto FAIL
	}

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		goto FAIL
	}

	// YAML转JSON
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		goto FAIL
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		goto FAIL
	}

	// 定义的container
	nginxContainer.Name = "nginx"
	nginxContainer.Image = "nginx:1.13.8"
	containers = append(containers, nginxContainer)

	// 修改podTemplate, 定义container列表
	deployment.Spec.Template.Spec.Containers = containers

	// 更新deployment
	if _, err = clientset.AppsV1().Deployments("default").Update(context.TODO(), &deployment, metav1.UpdateOptions{}); err != nil {
		goto FAIL
	}

	fmt.Println("apply成功!")
	return

FAIL:
	fmt.Println(err)
	return
}

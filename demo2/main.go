package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/owenliang/k8s-client-go/common"
	"io/ioutil"
	apps_v1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	yaml2 "k8s.io/apimachinery/pkg/util/yaml"
	"k8s.io/client-go/kubernetes"
)

func main() {
	var (
		clientset  *kubernetes.Clientset
		deployYaml []byte
		deployJson []byte
		deployment = apps_v1.Deployment{}
		replicas   int32
		err        error
	)

	//return
	// 初始化k8s客户端
	if clientset, err = common.InitClient(); err != nil {
		fmt.Println(err)
	}

	// 读取YAML
	if deployYaml, err = ioutil.ReadFile("./nginx.yaml"); err != nil {
		fmt.Println(err)
	}

	// YAML转JSON
	if deployJson, err = yaml2.ToJSON(deployYaml); err != nil {
		fmt.Println(err)
	}

	// JSON转struct
	if err = json.Unmarshal(deployJson, &deployment); err != nil {
		fmt.Println(err)
	}

	// 修改replicas数量为1
	replicas = 2
	deployment.Spec.Replicas = &replicas
	var c = context.TODO()
	// 查询k8s是否有该deployment
	_, err = clientset.AppsV1().Deployments("default").Get(c, deployment.Name, meta_v1.GetOptions{})
	fmt.Println(err)
	fmt.Println(deployment.Name, meta_v1.GetOptions{})
	if _, err = clientset.AppsV1().Deployments("default").Get(c, deployment.Name, meta_v1.GetOptions{}); err != nil {
		if !errors.IsNotFound(err) {
			fmt.Println("52")
			fmt.Println(err)
		}
		// 不存在则创建
		if _, err = clientset.AppsV1().Deployments("default").Create(c, &deployment, meta_v1.CreateOptions{}); err != nil {
			fmt.Println("57")
			fmt.Println(err)
		}
	} else { // 已存在则更新
		if _, err = clientset.AppsV1().Deployments("default").Update(c, &deployment, meta_v1.UpdateOptions{}); err != nil {
			fmt.Println(err)
		}
	}

	fmt.Println("apply成功!")
	return

}

package main

import (
	"context"
	"fmt"
	"github.com/bndr/gojenkins"
	"log"
)

func main() {
	jenkins := gojenkins.CreateJenkins(nil, "http://jenkins.p.nxin.com", "lemon9921", "lnsh1234")
	_, err := jenkins.Init(context.TODO())
	if err != nil {
		fmt.Println(err)
	}
	//连接成功
	fmt.Println("is ok")
	//获取节点状态
	nodes, _ := jenkins.GetAllNodes(context.TODO())
	for _, node := range nodes {
		node.Poll(context.TODO())
		if ok, _ := node.IsOnline(context.TODO()); ok {
			nodename := node.GetName()
			log.Printf("node is %s", nodename)
		}
	}
	//获取任务信息
	jobs, err := jenkins.GetAllJobNames(context.TODO())
	if err != nil {
		fmt.Println(err)
	}

	for _, job := range jobs {
		fmt.Println(job.Name, job.Url)
	}
	test := map[string]string{"aa": "/etc", "bb": "cc"}
	a, err := jenkins.BuildJob(context.TODO(), "test", test)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(a)
	al, err := jenkins.CopyJob(context.TODO(), "test", "test-all")
	fmt.Println(al)
	if err != nil {
		fmt.Println(err)
		return
	}
	//aaa ,err := jenkins.GetJob("test",)
	//if err != nil  {
	//	fmt.Println(err)
	//	return
	//}
}

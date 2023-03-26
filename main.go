package main

import (
	"context"
	"encoding/json"
	"fmt"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
	"log"
	"os"
)

// arg0: Namespace
// arg1: Primary Name

func main() {
	args := os.Args
	if len(args) >= 2 {
		config, err := clientcmd.BuildConfigFromFlags("", "/home/ec2-user/.kube/config")
		if err != nil {
			panic(err.Error())
		}
		clientset, err := kubernetes.NewForConfig(config)
		if err != nil {
			panic(err.Error())
		}

		//获取POD
		pods, err := clientset.CoreV1().Pods("carts").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			log.Println(err.Error())
		}
		fmt.Println(len(pods.Items))
		primary, _ := clientset.AppsV1().Deployments(args[1]).Get(context.TODO(), args[2], metav1.GetOptions{})
		jsonString, err := json.Marshal(primary)
		fmt.Println(primary.Status.Replicas)
		fmt.Println(string(jsonString))
	}
	if len(args) < 2 {
		fmt.Println("please add 2 arguments\n arg0: Namespace, arg1: Primary Name")
	}
}

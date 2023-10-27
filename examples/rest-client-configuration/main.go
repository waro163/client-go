package main

import (
	"context"
	"fmt"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
)

const (
	bearerToken = "jwt.token.xxx"
	apiHost     = "https://kubernetes.platform.example.com/"
	cluster     = "cn-pvg16-eng-general"
)

func main() {
	config := &rest.Config{
		BearerToken: bearerToken,
		Host:        apiHost + cluster,
		TLSClientConfig: rest.TLSClientConfig{
			Insecure: true,
		},
		QPS:   30,
		Burst: 60,
	}
	// create the clientset
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	nsList, err := clientset.CoreV1().Namespaces().List(context.Background(), metav1.ListOptions{})
	if err != nil {
		fmt.Println("list namespace error: ", err)
		return
	}
	for _, item := range nsList.Items {
		fmt.Println(item.ObjectMeta.Name)
	}
}

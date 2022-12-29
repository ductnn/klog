package main

import (
	"context"
	"fmt"

	"github.com/ductnn/klog/klient"
	"github.com/ductnn/klog/pkg/streams"
	"github.com/ductnn/klog/utils/color"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var banner = `
 ____        __  __
|  _ \ _   _|  \/  | ___
| | | | | | | |\/| |/ _ \
| |_| | |_| | |  | |  __/
|____/ \__,_|_|  |_|\___|
`

func main() {
	fmt.Println(color.Colorize(banner, 34))
	ctx := context.TODO()
	kubeConfig := klient.GetKubeConfigPath()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}

	// Create clientser
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	pod, err := clientset.CoreV1().Pods("").List(ctx, metav1.ListOptions{})
	if err != nil {
		panic(err)
	}

	for _, pod := range pod.Items {
		fmt.Printf("Pod name=/%s\n", pod.GetName())
		fmt.Println(streams.GetPodLogs(pod))
	}

}

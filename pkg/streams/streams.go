package streams

import (
	"bufio"
	"context"
	"fmt"
	"time"

	"github.com/ductnn/klog/klient"
	"github.com/ductnn/klog/utils/color"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	delay = time.Second
)

func GetPodLogs(pod corev1.Pod) string {
	ctx := context.Background()
	podLogOpts := corev1.PodLogOptions{}
	kubeConfig := klient.GetKubeConfigPath()

	config, err := clientcmd.BuildConfigFromFlags("", kubeConfig)
	if err != nil {
		panic(err)
	}

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	req := clientset.CoreV1().Pods(pod.Namespace).GetLogs(pod.Name, &podLogOpts)

	podLogs, err := req.Stream(ctx)
	if err != nil {
		panic(err)
	}

	defer podLogs.Close()

	for i := 0; i < 10; i++ {
		reader := bufio.NewScanner(podLogs)
		for reader.Scan() {
			line := reader.Text()
			fmt.Printf("%v: %v\n", color.Apply(pod.Name, 33), color.Apply(line, 32))
		}
		time.Sleep(1 * delay)
	}

	return ""
}

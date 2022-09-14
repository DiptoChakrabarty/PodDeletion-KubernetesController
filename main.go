package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"github.com/DiptoChakrabarty/podDeletionController/logger"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	podDelete := make(chan struct{})
	defer close(podDelete)
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "location of kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig file")
	}

	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)
	if err != nil {

	}
	//fmt.Println(config)

	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err)
	}

	informerfactory := informers.NewSharedInformerFactory(clientset, 30*time.Second)

	podinformer := informerfactory.Core().V1().Pods().Informer()
	podinformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			fmt.Println("Add was called")
		},
		UpdateFunc: func(old, new interface{}) {
			fmt.Println("Update was called")
		},
		DeleteFunc: func(obj interface{}) {
			fmt.Println("Delete was called")
			fmt.Println(obj)
			data, err := json.Marshal(obj)
			if err != nil {
				logger.Error("Unable to generate data", err)
			}
			fmt.Println(data)
		},
	})

	/* 	informerfactory.Start(wait.NeverStop)
	   	informerfactory.WaitForCacheSync(wait.NeverStop)
	   	pod, _ := podinformer.Lister().Pods("default").Get("default")
	   	fmt.Println(pod) */

	go podinformer.Run(podDelete)
	<-podDelete
}

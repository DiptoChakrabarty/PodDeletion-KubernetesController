package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"path/filepath"
	"time"

	"github.com/DiptoChakrabarty/podDeletionController/logger"
	"github.com/DiptoChakrabarty/podDeletionController/notif"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
)

func main() {
	var kubeconfig *string
	s := "This is deletion operation"
	ClientModel := notif.NewSlackClient()
	podDelete := make(chan struct{})
	defer close(podDelete)
	if home := homedir.HomeDir(); home != "" {
		kubeconfig = flag.String("kubeconfig", filepath.Join(home, ".kube", "config"), "location of kubeconfig file")
	} else {
		kubeconfig = flag.String("kubeconfig", "", "absolute path to kubeconfig file")
	}
	fmt.Println("KubeConfig Read")
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
			fmt.Println(s)
			//fmt.Println(obj)
			data, err := json.Marshal(obj)
			if err != nil {
				logger.Error("Unable to generate data", err)
			}
			//fmt.Println(data)
			channelID, timestamp, err := ClientModel.SendMessage(string(data))
			if err != nil {
				logger.Error(fmt.Sprintf("Unable to send message to channel ID %s", channelID), err)
			}
			logger.Info(fmt.Sprintf("Message sent successfully at %s to channel ID %s", timestamp, channelID))
		},
	})

	go podinformer.Run(podDelete)
	<-podDelete
}

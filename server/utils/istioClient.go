package utils

import (
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"path/filepath"
	versionedclient "istio.io/client-go/pkg/clientset/versioned"
)

func IstioClient() (istioClient *versionedclient.Clientset)  {
	var kubeconfigpath string
	if home := homedir.HomeDir(); home != "" {
		kubeconfigpath = filepath.Join(home, ".kube")
	} else {
		kubeconfigpath = "~/.kube/"
	}
	restConfig, err := clientcmd.BuildConfigFromFlags("", kubeconfigpath)
	if err != nil {
		log.Fatalf("Failed to create k8s rest client: %s", err)
	}

	ic, err := versionedclient.NewForConfig(restConfig)
	if err != nil {
		log.Fatalf("Failed to create istio client: %s", err)
	}
	return ic
}

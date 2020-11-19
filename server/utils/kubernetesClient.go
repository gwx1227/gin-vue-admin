package utils

import (
	"io/ioutil"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)


func KubeClient() (clientset *kubernetes.Clientset, err error) {
	var (
		restConf *rest.Config
	)

	if restConf, err = GetRestConf(); err != nil {
		return
	}
	// 生成clientset配置
	if clientset, err = kubernetes.NewForConfig(restConf); err != nil {
		return
	}

	return
}

// 获取k8s restful client 配置
func GetRestConf() (restConf *rest.Config, err error) {
	var (
		kubeconfig []byte
	)

	// 读取kubeconfig文件
	if kubeconfig, err = ioutil.ReadFile("/Users/gwx/.kube/config"); err != nil {
		return
	}
	// 生成rest client配置
	if restConf, err = clientcmd.RESTConfigFromKubeConfig(kubeconfig); err != nil {
		return
	}

	return
}
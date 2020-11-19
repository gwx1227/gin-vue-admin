package service

import (
	"fmt"
	resp "gin-vue-admin/model/response"
	"gin-vue-admin/utils"
	v1 "k8s.io/api/core/v1"
	appv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)
var (
	kubeClient *kubernetes.Clientset
	podList   *v1.PodList
	DeploymentList *appv1.DeploymentList
	err       error
	k8sDeployment *appv1.Deployment
)

func DeletePod(namespace string, pod_name string ) (res string,err error) {
	if kubeClient, err = utils.KubeClient(); err != nil {
		fmt.Println(err)
	}
	err = kubeClient.CoreV1().Pods(namespace).Delete(pod_name, &metav1.DeleteOptions{})
	if err != nil {
		res = fmt.Sprintf("重建POD失败: %v",err)
	}
	res = "重建POD成功"
	fmt.Printf("重建结果: %v\n", res)
	return res,err
}
func GetPodList(namespace string, app_name string, pod_env string ) (Pods []resp.PodInfo,err error)  {
	if kubeClient, err = utils.KubeClient(); err != nil {
		fmt.Println(err)
	}
	labels := map[string]string{
		"app.kubernetes.io/instance": app_name,
		"env": pod_env,
	}
	labelSelectorString := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels})
	options := metav1.ListOptions{
		LabelSelector: labelSelectorString,
	}
	podList,err = kubeClient.CoreV1().Pods(namespace).List(options)
	var PodInfo resp.PodInfo
    for PodIndex := range podList.Items {
    	PodInfo.CreationTimestamp = podList.Items[PodIndex].CreationTimestamp
    	PodInfo.Name = podList.Items[PodIndex].Name
    	PodInfo.ContainerCount = len(podList.Items[PodIndex].Spec.Containers)
    	PodInfo.Env = podList.Items[PodIndex].Labels["env"]
    	PodInfo.HostIp = podList.Items[PodIndex].Status.HostIP
    	PodInfo.NodeName = podList.Items[PodIndex].Spec.NodeName
    	PodInfo.PodIp = podList.Items[PodIndex].Status.PodIP
    	PodInfo.RestartCount = podList.Items[PodIndex].Status.ContainerStatuses[0].RestartCount
    	PodInfo.Phase = string(podList.Items[PodIndex].Status.Phase)
    	var ContainersInfo []resp.ContainerInfo
    	var ContainerInfo resp.ContainerInfo
    	for ContainerIndex := range podList.Items[PodIndex].Status.ContainerStatuses {
			ContainerInfo.Name = podList.Items[PodIndex].Status.ContainerStatuses[ContainerIndex].Name
			ContainerInfo.Ready = podList.Items[PodIndex].Status.ContainerStatuses[ContainerIndex].Ready
			ContainerInfo.RestartCount = podList.Items[PodIndex].Status.ContainerStatuses[ContainerIndex].RestartCount
			ContainersInfo = append(ContainersInfo, ContainerInfo)
		}
		PodInfo.ContainersInfo = ContainersInfo
		Pods = append(Pods, PodInfo)
	}
	//fmt.Printf("获取POD列表: %v\n", Pods)
	return Pods, err
}


func GetdeploymentList(namespace string, app_name string ) (Deployments []resp.DeploymentInfo,err error)  {
	if kubeClient, err = utils.KubeClient(); err != nil {
		fmt.Println(err)
	}
	labels := map[string]string{
		"app.kubernetes.io/instance": app_name,
	}
	labelSelectorString := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels})
	options := metav1.ListOptions{
		LabelSelector: labelSelectorString,
	}
	DeploymentList,err = kubeClient.AppsV1().Deployments(namespace).List(options)
	var DeploymentInfo resp.DeploymentInfo
	for DeploymentIndex :=range DeploymentList.Items{
		DeploymentInfo.PodEnv = DeploymentList.Items[DeploymentIndex].Spec.Template.Labels["env"]
		DeploymentInfo.DeploymentName = DeploymentList.Items[DeploymentIndex].Name
		DeploymentInfo.CreationTimestamp = DeploymentList.Items[DeploymentIndex].CreationTimestamp
		DeploymentInfo.Replicas = DeploymentList.Items[DeploymentIndex].Status.Replicas
		DeploymentInfo.UnavailableReplicas = DeploymentList.Items[DeploymentIndex].Status.UnavailableReplicas
		Deployments = append(Deployments, DeploymentInfo)
	}
	return Deployments,err
}

func Getdeployment(namespace string,app_name string, deployment_name string ) (*appv1.Deployment, error)  {
	if kubeClient, err = utils.KubeClient(); err != nil {
		fmt.Println(err)
	}
	labels := map[string]string{
		"app.kubernetes.io/instance": app_name,
	}
	labelSelectorString := metav1.FormatLabelSelector(&metav1.LabelSelector{MatchLabels: labels})
	options := metav1.ListOptions{
		LabelSelector: labelSelectorString,
	}
	fmt.Printf("....2%v\n", options)
	k8sDeployment,err = kubeClient.AppsV1().Deployments(namespace).Get(deployment_name,metav1.GetOptions{})
	return k8sDeployment,err
}

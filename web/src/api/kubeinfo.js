import service from '@/utils/request'



 export const getPodList = (params) => {
     return service({
         url: "/kubeinfo/getPodList",
         method: 'get',
         params
     })
}
 export const restartPod = (params) => {
     return service({
         url: "/kubeinfo/restartPod",
         method: 'post',
         params
     })
}
 export const getDeploymentList = (params) => {
     return service({
         url: "/kubeinfo/getDeploymentList",
         method: 'get',
         params
     })
    }
 export const getDeployment = (params) => {
     return service({
         url: "/kubeinfo/getDeployment",
         method: 'get',
         params
     })
 }    
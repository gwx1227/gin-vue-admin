package initialize

import (
	_ "gin-vue-admin/docs"
	"gin-vue-admin/global"
	"gin-vue-admin/middleware"
	"gin-vue-admin/router"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()
	Router.StaticFS(global.GVA_CONFIG.Local.Path, http.Dir(global.GVA_CONFIG.Local.Path)) // 为用户头像和文件提供静态地址
	// Router.Use(middleware.LoadTls())  // 打开就能玩https了
	global.GVA_LOG.Info("use middleware logger")
	// 跨域
	Router.Use(middleware.Cors())
	global.GVA_LOG.Info("use middleware cors")
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.GVA_LOG.Info("register swagger handler")
	// 方便统一添加路由组前缀 多服务器上线使用
	ApiGroup := Router.Group("")
	router.InitUserRouter(ApiGroup)                  // 注册用户路由
	router.InitBaseRouter(ApiGroup)                  // 注册基础功能路由 不做鉴权
	router.InitMenuRouter(ApiGroup)                  // 注册menu路由
	router.InitAuthorityRouter(ApiGroup)             // 注册角色路由
	router.InitApiRouter(ApiGroup)                   // 注册功能api路由
	router.InitFileUploadAndDownloadRouter(ApiGroup) // 文件上传下载功能路由
	router.InitSimpleUploaderRouter(ApiGroup)        // 断点续传（插件版）
	router.InitWorkflowRouter(ApiGroup)              // 工作流相关路由
	router.InitCasbinRouter(ApiGroup)                // 权限相关路由
	router.InitJwtRouter(ApiGroup)                   // jwt相关路由
	router.InitSystemRouter(ApiGroup)                // system相关路由
	router.InitCustomerRouter(ApiGroup)              // 客户路由
	router.InitAutoCodeRouter(ApiGroup)              // 创建自动化代码
	router.InitSysDictionaryDetailRouter(ApiGroup)   // 字典详情管理
	router.InitSysDictionaryRouter(ApiGroup)         // 字典管理
	router.InitSysOperationRecordRouter(ApiGroup)    // 操作记录
	router.InitEmailRouter(ApiGroup)                 // 邮件相关路由
	router.InitNamespacesRouter(ApiGroup)            // 命名空间相关路由
	router.InitBusinessRouter(ApiGroup)              // 一级业务线相关路由
	router.InitSubBusinessRouter(ApiGroup)           // 二级业务线相关路由
	router.InitAliasRouter(ApiGroup)                 // 别名配置相关路由
	router.InitAppsRouter(ApiGroup)                  // APP信息相关路由
	router.InitBuildHistoryRouter(ApiGroup)          // 构建信息相关路由
	router.InitConfigRouter(ApiGroup)                // 配置信息相关路由
	router.InitDeployRouter(ApiGroup)                // 部署信息相关路由
	router.InitDeployHistoryRouter(ApiGroup)         // 部署历史相关路由
	router.InitGatewayRouter(ApiGroup)               // 域名配置相关路由
	router.InitHpaRouter(ApiGroup)                   // 自动扩展配置相关路由
	router.InitLivenessRouter(ApiGroup)              // 健康检查配置相关路由
	router.InitMetricsRouter(ApiGroup)               // 监控数据配置相关路由
	router.InitReadinessRouter(ApiGroup)             // 就绪检查配置相关路由
	router.InitResourcesRouter(ApiGroup)             // 资源配置相关路由
	router.InitSysConfigRouter(ApiGroup)             // 系统变量配置相关路由
	router.InitKubeRouter(ApiGroup)             // 动态kubernetes信息相关路由

	global.GVA_LOG.Info("router register success")
	return Router
}

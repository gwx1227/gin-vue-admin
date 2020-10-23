// 自动生成模板Deploy
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type Deploy struct {
	gorm.Model
	AppId              int    `json:"appId" form:"appId" gorm:"column:app_id;uniqueIndex;not null;comment:关联应用id"`
	ArgsInfo           string `json:"argsInfo" form:"argsInfo" gorm:"column:args_info;comment:args内容"`
	ArgsSwitch         *bool  `json:"argsSwitch" form:"argsSwitch" gorm:"column:args_switch;not null;default:false;comment:是否开启args"`
	CommandInfo        string `json:"commandInfo" form:"commandInfo" gorm:"column:command_info;comment:command内容"`
	CommandSwitch      *bool  `json:"commandSwitch" form:"commandSwitch" gorm:"column:command_switch;not null;default:false;comment:是否开启command"`
	ContainerPort      int    `json:"containerPort" form:"containerPort" gorm:"column:container_port;not null;comment:容器port"`
	PullPolicy         string `json:"pullPolicy" form:"pullPolicy" gorm:"column:pull_policy;not null;default:'IfNotPresent';comment:拉取镜像策略"`
	ReplicaCountCanary int    `json:"replicaCountCanary" form:"replicaCountCanary" gorm:"column:replica_count_canary;not null;default:0;comment:canary副本数"`
	ReplicaCountOnline int    `json:"replicaCountOnline" form:"replicaCountOnline" gorm:"column:replica_count_online;not null;default:2;comment:online副本数"`
	Repository         string `json:"repository" form:"repository" gorm:"column:repository;not null;comment:docker镜像地址"`
	TagCanary          string `json:"tagCanary" form:"tagCanary" gorm:"column:tag_canary;comment:canary代码版本"`
	TagOnline          string `json:"tagOnline" form:"tagOnline" gorm:"column:tag_online;comment:online代码版本"`
	WeigitCanary       int    `json:"weigitCanary" form:"weigitCanary" gorm:"column:weigit_canary;not null;default:0;comment:canary流量权重"`
	WeigitOnline       int    `json:"weigitOnline" form:"weigitOnline" gorm:"column:weigit_online;not null;default:100;comment:online流量权重"`
}

func (Deploy) TableName() string {
	return "deploy"
}

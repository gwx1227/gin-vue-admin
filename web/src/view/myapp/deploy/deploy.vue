<template>
    <div v-if="currentAppId !== null">
       <el-form ref="ruleForm" label-position="right" :model="ruleForm" :rules="rules" :inline="true" label-width="180px" class="demo-ruleForm">
      <el-card>
        <el-form-item prop="containerPort">
          <span slot="label">
            <span slot="reference">容器端口 </span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>端口输入规范:</span><br>
                <span>数字，2-5个字符</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="ruleForm.containerPort" />
        </el-form-item> <br>
        <el-form-item label="镜像地址" prop="repository">
          <span slot="label">
            <span slot="reference">镜像地址 </span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>镜像地址输入规范:</span><br>
                <span>建议标准: NAMESPACE/APPNAME ，80个字符以内</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="ruleForm.repository" />
        </el-form-item><br>
        <el-form-item label="镜像拉取策略" prop="pullPolicy">
          <span slot="label">
            <span slot="reference">镜像拉取策略</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>镜像拉取策略说明:</span><br>
                <span>Always:镜像标签为"latest"或镜像不存在时总是从指定的仓库中获取镜像.</span><br>
                <span>IfNotPresent:仅当本地镜像缺失时方才从目标仓库下载镜像.</span><br>
                <span>Never:禁止从仓库下载镜像,即仅使用本地镜像.</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-select v-model="ruleForm.pullPolicy">
            <el-option
              v-for="item in image_pull"
              :key="item"
              :label="item"
              :value="item"
            />
          </el-select>
        </el-form-item>
      </el-card>
      <br><br>
      <el-card>
        <el-form-item label="online版本" prop="tagOnline">
          <span slot="label">
            <span slot="reference">online版本</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>正常生产部署版本,80个字符以内</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="ruleForm.tagOnline" />
        </el-form-item>
        <el-form-item label="online权重" prop="weightOnline">
          <span slot="label">
            <span slot="reference">online权重</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>正常生产流量权重,整体流量权重需要为100</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input-number v-model="ruleForm.weightOnline" :min="0" :max="100" @change="weightOnline_HandleChange" />

        </el-form-item>
        <el-form-item label="online副本数" prop="replicaCountOnline">
          <span slot="label">
            <span slot="reference">online副本数</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>正常生产版本副本数,建议最少2个</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input-number v-model="ruleForm.replicaCountOnline" :min="1" :max="10" @change="replicaCountOnline_HandleChange" />
        </el-form-item>
        <br>
        <el-form-item label="canary版本" prop="tagCanary">
          <span slot="label">
            <span slot="reference">canary版本</span>
            <el-tooltip placement="top">
              <div slot="content">
                <div slot="content">
                  <span>说明:</span><br>
                  <span>金丝雀部署版本,80个字符以内</span>
                </div>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input v-model="ruleForm.tagCanary" />
        </el-form-item>
        <el-form-item label="canary权重" prop="weightCanary">
          <span slot="label">
            <span slot="reference">canary权重</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>金丝雀流量权重,整体流量权重需要为100</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input-number v-model="ruleForm.weightCanary" :min="0" :max="100" @change="weightCanary_HandleChange" />

        </el-form-item>

        <el-form-item label="canary副本数" prop="replicaCountCanary">
          <span slot="label">
            <span slot="reference">canary副本数</span>
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>金丝雀版本副本数,建议最少1个,如果不做金丝雀发布的话可以为0,但是前提是没有金丝雀流量.</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
          </span>
          <el-input-number v-model="ruleForm.replicaCountCanary" :min="0" :max="10" @change="replicaCountCanary_HandleChange" />

        </el-form-item>

      </el-card><br><br>
      <el-card>
        <el-form-item>
          <span>启用 COMMAND 配置
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>如果仅为容器定义了command字段,那么它将覆盖镜像中定义的程序及参数,并以无参数方式运行应用程序.</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
            <el-switch
              v-model="ruleForm.commandSwitch"
              active-color="#13ce66"
              inactive-color="#ff4949"
              inactive-text="关"
              active-text="开"
            /></span>
        </el-form-item>
        <div v-if="ruleForm.commandSwitch">
          <el-form-item label="COMMAND内容" prop="commandInfo">
            <span slot="label">
              <span slot="reference">COMMAND内容</span>
              <el-tooltip placement="top">
                <div slot="content">
                  <span>说明:</span><br>
                  <span>例如: "/bin/sh"</span>
                </div>
                <i class="el-icon-question" />
              </el-tooltip>
            </span>
            <el-input v-model="ruleForm.commandInfo" />
          </el-form-item>
        </div><br>
        <el-form-item>
          <div>启用 ARGS 配置
            <el-tooltip placement="top">
              <div slot="content">
                <span>说明:</span><br>
                <span>如果仅为容器定义了args字段,那么它将作为参数传递给镜像中默认指定运行的应用程序.</span>
              </div>
              <i class="el-icon-question" />
            </el-tooltip>
            <el-switch
              v-model="ruleForm.argsSwitch"
              active-color="#13ce66"
              inactive-color="#ff4949"
              inactive-text="关"
              active-text="开"
            /></div>
        </el-form-item>
        <div v-if="ruleForm.argsSwitch">
          <el-form-item label="ARGS内容" prop="argsInfo">
            <span slot="label">
              <span slot="reference">ARGS内容</span>
              <el-tooltip placement="top">
                <div slot="content">
                  <span>说明:</span><br>
                  <span>多个参数间使用","分隔,例如: "-c","while true; do sleep 30; done"</span>
                </div>
                <i class="el-icon-question" />
              </el-tooltip>
            </span>
            <el-input v-model="ruleForm.argsInfo" />
          </el-form-item>
        </div>
      </el-card>
    </el-form><br>
    <div style="text-align:center">
      <el-button round type="info" @click="onClose">取消</el-button>
      <el-button round type="info" @click="getInfo">test</el-button>
      <el-button
        round
        type="primary"
        @click="saveDeployData()"
      >保存</el-button>
    </div>
    </div>
</template>

<script>
import {
  findDeployByAppId,
  updateDeploy, 
} from "@/api/deploy";  //  此处请自行替换地址
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Deploy",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId'])
  },
  beforeMount() {
    this.getDeployData()
  },
  data() {
    return {
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      query: {
        appId: ''
      },
      image_pull: ['Always', 'Never', 'IfNotPresent'],
      rules: {
        containerPort: [
          { required: true, message: '请输入容器端口', trigger: 'blur' },
          { required: true, pattern: /(^\d{2,5}$)/, message: '长度在 2 到 5 个数字', trigger: 'blur' }
        ],
        repository: [
          { required: true, message: '请输入镜像地址', trigger: 'blur' }
        ],
        tagOnline: [
          { required: true, message: '请输入镜像地址', trigger: 'blur' }
        ]
      },
      info: '',
      ruleForm: {
        appId: "",
        containerPort: '',
        repository: '',
        pullPolicy: '',
        weightOnline: '',
        weightCanary: '',
        replicaCountOnline: '',
        replicaCountCanary: '',
        tagOnline: '',
        tagCanary: '',
        commandSwitch: '',
        commandInfo: '',
        argsSwitch: '',
        argsInfo: ''
      },
      multipleSelection: [],formData: {
        appId:null,argsInfo:null,argsSwitch:null,commandInfo:null,commandSwitch:null,containerPort:null,pullPolicy:null,replicaCountCanary:null,replicaCountOnline:null,repository:null,tagCanary:null,tagOnline:null,weigitCanary:null,weigitOnline:null,
      }
    };
  },
  filters: {
    formatDate: function(time) {
      if (time != null && time != "") {
        var date = new Date(time);
        return formatTimeToStr(date, "yyyy-MM-dd hh:mm:ss");
      } else {
        return "";
      }
    },
    formatBoolean: function(bool) {
      if (bool != null) {
        return bool ? "是" :"否";
      } else {
        return "";
      }
    }
  },
  methods: {
   getDeployData(){
      this.query.appId = this.currentAppId
      findDeployByAppId(this.query).then( res => {
        this.ruleForm.containerPort = res.data.redeploy.containerPort
        this.ruleForm.repository = res.data.redeploy.repository
        this.ruleForm.pullPolicy = res.data.redeploy.pullPolicy
        this.ruleForm.weightOnline = res.data.redeploy.weightOnline
        this.ruleForm.weightCanary = res.data.redeploy.weightCanary
        this.ruleForm.replicaCountOnline = res.data.redeploy.replicaCountOnline
        this.ruleForm.replicaCountCanary = res.data.redeploy.replicaCountCanary
        this.ruleForm.tagOnline = res.data.redeploy.tagOnline
        this.ruleForm.tagCanary = res.data.redeploy.tagCanary
        this.ruleForm.commandSwitch = res.data.redeploy.commandSwitch
        this.ruleForm.commandInfo = res.data.redeploy.commandInfo
        this.ruleForm.argsSwitch = res.data.redeploy.argsSwitch
        this.ruleForm.argsInfo = res.data.redeploy.argsInfo
       
        }
      )
   }, 
   saveDeployData() {
      this.ruleForm.appId = this.currentAppId
      updateDeploy(this.ruleForm).then(res => {
        console.log("staus: ", res.code)
        this.getDeployData()
        this.$notify({
          title: '成功',
          message: '编辑部署配置信息完成.',
          type: 'success'
        })
      })
    },
    onClose() {
      this.getDeployData()
      this.$emit('on-close')
    },
    weightCanary_HandleChange(value) {
      this.ruleForm.weightOnline = 100-value 
    },
    weightOnline_HandleChange(value) {
      this.ruleForm.weightCanary = 100-value 
    },
    replicaCountCanary_HandleChange(value) {
      this.ruleForm.replicaCountCanary = value
    },
    replicaCountOnline_HandleChange(value) {
      this.ruleForm.replicaCountOnline = value
    },
    getInfo() {
      console.log('.....1', this.ruleForm.weightCanary)
      console.log('.....2', this.ruleForm.weightOnline)
    }
  },
}
</script>

<style>
</style>
<template>
  <div v-show="currentAppId !== null">
            <el-button type="primary" round  icon="el-icon-back"  @click="resetCurrentAppId">返回应用列表</el-button> <br>
            <!-- <el-button type="primary" round  icon="el-icon-back"  @click="getDeploymentData">定位</el-button>  -->
            <br>
    <el-card>
      <el-table
        :data="tableData"
        border
        stripe
        style="width: 100%"
      >
        <el-table-column
          prop="appName"
          label="应用名称"
          align="center"
          min-width="30%"
        />
        <el-table-column
          prop="gitUrl"
          label="代码地址"
          align="center"
          min-width="30%"
        />
        <el-table-column
          prop="language"
          label="编程语言"
          align="center"
          min-width="10%"
        />
        <el-table-column
          prop="currentEnv"
          label="应用环境"
          align="center"
          min-width="10%"
        />
      </el-table>
      <el-table
        :data="tableData"
        border
        stripe
        style="width: 100%"
      >
        <el-table-column
          prop="business"
          label="主业务线"
          align="center"
        />
        <el-table-column
          prop="subbusiness"
          label="次业务线"
          align="center"
        />
        <el-table-column
          prop="createUser"
          label="创建人员"
          align="center"
        />
        <el-table-column
          prop="CreatedAt"
          label="创建时间"
          align="center"
          :formatter="dateFormat"
        />
      </el-table>
    </el-card>
    <br><br>
    <el-card>
      <el-table
        :data="deployment_data" stripe="True" fit='True' default-expand-all='True' highlight-current-row  style="width: 100%"
      >
        <el-table-column
          prop="pod_env"
          label="流程状态"
          align="center"
           min-width="10%"
        />
        <el-table-column
          prop="deployment_name"
          label="实例名称"
          align="center"
           min-width="30%"
        />

        <el-table-column
          prop="replicas"
          label="副本数量"
          align="center"
           min-width="5%"
        />
        <el-table-column
          prop="unavailable_replicas"
          label="当前状态"
          align="center"
           min-width="5%"
        ><template slot-scope="scope"><span v-if="scope.row.unavailable_replicas==0" style="color: green">健康</span>
          <span v-else style="color:red">不健康</span>
        </template></el-table-column>
        <el-table-column
          prop="creation_timestamp"
          label="应用时间"
          align="center"
           min-width="15%"
          :formatter="dateFormat"
        />
        <el-table-column
          label="操作"
           min-width="5%"
        > <template slot-scope="scope">
          <el-button icon="el-icon-search" type="primary" circle @click="getPodsData(deployment_name=scope.row.pod_env)" />
          <el-dialog  customClass="customWidth" append-to-body  title="POD详情" :visible.sync="dialogTableVisible">
            <template>
              <el-table
                :data="pods_data"
                style="width: 100%"
              ><el-table-column type="expand">
                 <template slot-scope="prop">
                   <el-table :data="prop.row.ContainersInfo" style="width: 100%">
                     <el-table-column
                       align="center"
                       prop="name"
                       label="容器名称"
                       min-width="55%"
                     />
                     <el-table-column
                       align="center"
                       prop="ready"
                       label="就绪状态"
                       min-width="15%"
                     ><template slot-scope="scope_ready"><span v-if="scope_ready.row.unavailable_replicas='True'" style="color: green">健康</span>
                       <span v-else style="color: yellow">不健康</span>
                     </template></el-table-column>
                     <el-table-column
                       align="center"
                       prop="restart_count"
                       label="重启次数"
                       min-width="15%"
                     />
                     <el-table-column
                       label="操作"
                       min-width="15%"
                     > <template slot-scope="scope_terminal">
                       <el-button icon="el-icon-search" type="primary" circle @click="toTerminal(pod=prop.row.name,container=scope_terminal.row.name)" />  </template>
                     </el-table-column></el-table>

                 </template>
               </el-table-column>
                <el-table-column
                  prop="name"
                  label="POD名称"
                  align="center"
                  min-width="40%"
                />
                <el-table-column
                  prop="pod_ip"
                  label="POD地址"
                  align="center"
                  min-width="15%"
                />
                <el-table-column
                  prop="phase"
                  label="POD状态"
                  align="center"
                  min-width="10%"
                ><template slot-scope="scope_phase"><span v-if="scope_phase.row.unavailable_replicas='Running'" style="color: green">健康</span>
                  <span v-else style="color: yellow">不健康</span>
                </template></el-table-column>
                <el-table-column
                  prop="container_count"
                  label="容器数量"
                  align="center"
                  min-width="10%"
                />
                <el-table-column
                  prop="creation_timestamp"
                  label="应用时间"
                  align="center"
                  min-width="18%"
                  :formatter="dateFormat"
                />
                <el-table-column
                  prop="host_ip"
                  label="宿主机IP"
                  align="center"
                  min-width="15%"
                />
                <el-table-column
                  prop="node_name"
                  label="宿主机"
                  align="center"
                  min-width="15%"
                />
                <el-table-column
                  label="操作"
                  min-width="10%"
                > <template slot-scope="scope_pod">
                  <el-button type="text" circle @click="restartPodData(pod_name=scope_pod.row.name)">重启</el-button>
                </template> </el-table-column>
              </el-table></template>
          </el-dialog>

        </template></el-table-column>
      </el-table>
    </el-card> 
  </div>
</template>

<script>
import {
    findApps,
} from "@/api/apps";  //  此处请自行替换地址
import {
  getBusinessList,
} from "@/api/business";
import {
  getSubBusinessList,
} from "@/api/subBusiness";
import {
  getPodList,
  restartPod,
  getDeploymentList,
} from "@/api/kubeinfo";
import { formatTimeToStr } from "@/utils/data";
import { mapGetters, mapMutations} from 'vuex'
import moment from 'moment'
export default {
  name: "Apps",
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  data() {
    return {
      tableData: {},
      deployment_data: [],
      k8s_data: {
        deployment_data: [],
        pods_data: []
      },
      dialogTableVisible: false,
      pods_data: [],
      dialogFormVisible: false,
      visible: false,
      business: getBusinessList,
      subbusiness: getSubBusinessList,
      type: "",
      deleteVisible: false,
      query: {
        ID: '',
      },
      restart_data: {
        ID: "",
        pod_name: ''
      },
      deployment_query: {
        ID: '',
      },
      pods_query: {
        pod_env: '' ,
        ID: ''
      },
      multipleSelection: [],
      formData: {
        appName:null,businessId:null,createUser:null,currentEnv:null,gatewaySwitch:null,gitUrl:null,hpaSwitch:null,language:null,livenessSwitch:null,metricsSwitch:null,namespaceId:null,readinessSwitch:null,subbusinessId:null,updateUser:null,usable:null,
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
  watch:{
     currentAppId(val){
       this.getInfoData()
     }, 
  },
  beforeMount() {
    this.getInfoData()
    this.getDeploymentData()
  },
  methods: {
      ...mapMutations('user',['ResetCurrentAppId']),
      dingwei(){
      },
      dateFormat(row, column) {
        var date = row[column.property]
        if (date === undefined) {
          return ''
        }
        return moment(date).format('YYYY-MM-DD HH:mm:ss')
      },
      getInfoData() {
        this.query.ID = this.currentAppId;
        findApps(this.query).then(res => {
          this.tableData = res.data.reapps
        })
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
    async updateApps(row) {
      const res = await findApps({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reapps;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = { 
          appName:null,
          businessId:null,
          createUser:null,
          currentEnv:null,
          gatewaySwitch:null,
          gitUrl:null,
          hpaSwitch:null,
          language:null,
          livenessSwitch:null,
          metricsSwitch:null,
          namespaceId:null,
          readinessSwitch:null,
          subbusinessId:null,
          updateUser:null,
          usable:null,
      };
    },
    getDeploymentData() {
      this.deployment_query.ID = this.currentAppId;
      getDeploymentList(this.deployment_query).then(res => {
        this.deployment_data = res.data.list
        this.$set(this.k8s_data, 'deployment_data', res.data.list)
        // this.deployment_data = res
      })
    },
    getPodsData(pod_env) {
      this.dialogTableVisible = true
      this.pods_query.ID = this.currentAppId;
      this.pods_query.pod_env = pod_env
      getPodList(this.pods_query).then(res => {
        this.pods_data = res.data.list
      })
      const chatTimer_pod = setInterval(() => {
        getPodList(this.pods_query).then(res => {
          this.pods_data = res.data.list
        })
      }, 5000)
      this.$once('hook:beforeDestroy', () => {
        clearInterval(chatTimer_pod)
      })},
      restartPodData(pod_name) {
      this.restart_data.ID = this.currentAppId;
      console.log("...1", this.restart_data)
      this.restart_data.pod_name = pod_name
      restartPod(this.restart_data).then(res => {
        console.log("...2", res)
        this.$notify({
          title: '成功',
          message: '重启POD完成.',
          type: 'success'
        })
      })
      },
    resetCurrentAppId() {
      this.ResetCurrentAppId(null)
    },
  }
}
</script>

<style>
.customWidth{
        width:80%;
    }
</style>
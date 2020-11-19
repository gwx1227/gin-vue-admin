<template>
<div>
<div v-if="currentAppId === null">
   <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
      我的应用
    </el-row>
    <el-row>
      <span style="padding:20px">命名空间</span>
      <el-select v-model="nsvalue" filterable default-first-option placeholder="请选择命名空间" @change="changeNamespace">
        <el-option
          v-for="item in namespaces"
          :key="item.ID"
          :label="item.namespace"
          :value="item.ID"
        />
      </el-select> 
         <span style="padding:20px"></span>
       <el-button type="primary" @click="dialogFormVisibleAdd = true">创建应用</el-button>
    </el-row> 
    
    <el-tabs style="padding-left:20px;margin-top:20px;">
      <el-tab-pane label="应用">
        <el-row>
          <span style="padding:20px">应用名称</span>
          <el-select v-model="appvalue" filterable default-first-option clearable placeholder="请选择应用名称" @change="changeApp">
            <el-option
              v-for="item in tableData"
              :key="item.ID"
              :label="item.appName"
              :value="item.ID"
            />
          </el-select>
         
  <el-dialog append-to-body title="创建应用" :visible.sync="dialogFormVisibleAdd">
      <el-form ref="addData" :model="addData" label-position="top" :rules="rules" :inline="true" :label-width="formLabelWidth" class="demo-ruleForm">
        <el-card>
          <el-form-item label="namespace" :label-width="formLabelWidth">
            <el-select v-model="addData.namespaceId" placeholder="请选择命名空间" @change="changeNamespace">
              <el-option
                v-for="item in namespaces"
                :key="item.ID"
                :label="item.namespace"
                :value="item.ID"
              />
            </el-select>
          </el-form-item><br>
          <el-form-item label="应用名称" prop="app_name">
            <span slot="label">
              <span slot="reference">应用名称 </span>
              <el-tooltip placement="top">
                <div slot="content">
                  <span>格式说明:</span><br>
                  <span>例如: my-book-info</span>
                </div>
                <i class="el-icon-question" />
              </el-tooltip>
            </span>
            <el-input v-model="addData.appName" />
          </el-form-item>
          <el-form-item label="git地址" prop="git_url">
            <span slot="label">
              <span slot="reference">git地址 </span>
              <el-tooltip placement="top">
                <div slot="content">
                  <span>说明:</span><br>
                  <span>请输入ssh格式地址</span>
                </div>
                <i class="el-icon-question" />
              </el-tooltip>
            </span>
            <el-input v-model="addData.gitUrl" />
          </el-form-item><br>
          <el-form-item prop="container_port">
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
            <el-input v-model="addData.containerPort" />
          </el-form-item>
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
            <el-input v-model="addData.repository" />
          </el-form-item>
          <el-form-item label="镜像拉取策略" prop="pull_policy">
            <span slot="label">
              <span slot="reference">镜像拉取策略 </span>
            </span>
            <el-select v-model="addData.pullPolicy" placeholder="请选择镜像拉取策略">
              <el-option
                v-for="item in image_pull"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select></el-form-item><br>
          <el-form-item label="主业务线" prop="business_id">
            <span slot="label">
              <span slot="reference">主业务线 </span>
            </span>
            <el-select v-model="addData.businessId" placeholder="请选择一级业务线" @change="changeBusiness">
              <el-option
                v-for="item in business"
                :key="item.ID"
                :label="item.business"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="次业务线" prop="subbusiness_id">
            <span slot="label">
              <span slot="reference">次业务线 </span>
            </span>
            <el-select v-model="addData.subbusinessId" placeholder="请选择二级业务线" @change="changeSubBusiness">
              <el-option
                v-for="item in subbusiness"
                :key="item.ID"
                :label="item.subbusiness"
                :value="item.ID"
              />
            </el-select>
          </el-form-item>
          <el-form-item label="开发语言" prop="language">
            <span slot="label">
              <span slot="reference">开发语言 </span>
            </span>
            <el-select v-model="addData.language" placeholder="请选择应用开发语言">
              <el-option
                v-for="item in languages"
                :key="item.value"
                :label="item.label"
                :value="item.label"
              />
            </el-select>
          </el-form-item><br>
          <el-form-item label="CPU限额" prop="cpu_limit">
            <span slot="label">
              <span slot="reference">CPU限额 </span>
            </span>
            <el-select v-model="addData.cpuLimit" placeholder="请选择最大申请CPU额度">
              <el-option
                v-for="item in cpu"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select></el-form-item>
          <el-form-item label="CPU需求" prop="cpu_requests">
            <span slot="label">
              <span slot="reference">CPU需求 </span>
            </span>
            <el-select v-model="addData.cpuRequests" placeholder="请选择最小申请CPU额度">
              <el-option
                v-for="item in cpu"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select></el-form-item>
          <el-form-item label="内存限额" prop="mem_limit">
            <span slot="label">
              <span slot="reference">内存限额 </span>
            </span>
            <el-select v-model="addData.memLimit" placeholder="请选择最大申请内存额度">
              <el-option
                v-for="item in mem"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select></el-form-item>
          <el-form-item label="内存需求" prop="mem_requests">
            <span slot="label">
              <span slot="reference">内存需求 </span>
            </span>
            <el-select v-model="addData.memRequests" placeholder="请选择最小申请内存额度">
              <el-option
                v-for="item in mem"
                :key="item"
                :label="item"
                :value="item"
              />
            </el-select></el-form-item><br>
        </el-card><br>
        <el-card>
           <el-form-item label="网关域名" :required="true" :label-width="formLabelWidth"> 
               <el-switch 
                v-model="addData.gatewaySwitch"
                  active-text="开"
                inactive-text="关"
              />
                </el-form-item>
               <el-form-item label="自动伸缩" :required="true" :label-width="formLabelWidth">
              <el-switch
                v-model="addData.hpaSwitch"
                  active-text="开"
                inactive-text="关"
              />
               </el-form-item>
               <el-form-item label="健康检查" :required="true" :label-width="formLabelWidth">
              <el-switch
                v-model="addData.livenessSwitch"
                  active-text="开"
                inactive-text="关"
              />
               </el-form-item>
               <el-form-item label="就绪检查" :required="true" :label-width="formLabelWidth">
              <el-switch
                v-model="addData.readinessSwitch"
                  active-text="开"
                inactive-text="关"
              />
               </el-form-item>
               <el-form-item label="指标数据" :required="true" :label-width="formLabelWidth">
              <el-switch
                v-model="addData.metricsSwitch"
                active-text="开"
                inactive-text="关"
              /> 
               </el-form-item>
        </el-card><br>
        <el-card>
          <h3>COMMAND开关<el-tooltip placement="top">
            <div slot="content">
              <span>说明:</span><br>
              <span>如果仅为容器定义了command字段,那么它将覆盖镜像中定义的程序及参数,并以无参数方式运行应用程序.</span>
            </div>
            <i class="el-icon-question" />
          </el-tooltip></h3><br>
          <el-switch
            v-model="addData.commandSwitch"
            style="display: block"
            active-color="#13ce66"
            inactive-color="#ff4949"
            inactive-text="关"
            active-text="开"
          />
          <div v-if="addData.commandSwitch">
            <el-form-item label="COMMAND内容" prop="command_info">
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
              <el-input v-model="addData.commandInfo" />
            </el-form-item>
          </div><br>
          <h3>ARGS开关 <el-tooltip placement="top">
            <div slot="content">
              <span>说明:</span><br>
              <span>如果仅为容器定义了args字段,那么它将作为参数传递给镜像中默认指定运行的应用程序.</span>
            </div>
            <i class="el-icon-question" />
          </el-tooltip></h3><br>
          <el-switch
            v-model="addData.argsSwitch"
            style="display: block"
            active-color="#13ce66"
            inactive-color="#ff4949"
            inactive-text="关"
            active-text="开"
          />
          <div v-if="addData.argsSwitch">
            <el-form-item label="ARGS内容" prop="args_info">
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
              <el-input v-model="addData.argsInfo" />
            </el-form-item>
          </div>
        </el-card>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisibleAdd = false">取 消</el-button>
        <el-button type="primary" @click="createAppData()">确 定</el-button>
      </div>
    </el-dialog>
        </el-row>
        <el-table
          :data="tableData"
          style="width: 100%"
        >
          <el-table-column
            v-show="false"
            prop="ID"
            label="appId"
            align="center"
          />
          <el-table-column
            prop="appName"
            label="应用名称"
            align="center"
          />
          <el-table-column
            prop="currentEnv"
            label="当前环境"
            align="center"
          />
          <el-table-column
            prop="createUser"
            label="创建人"
            align="center"
          />
          <el-table-column
            prop="CreatedAt"
            label="创建时间"
            align="center"
            :formatter="dateFormat"
          />
          <el-table-column
            prop="UpdatedAt"
            label="更新时间"
            align="center"
            :formatter="dateFormat"
          />
          <el-table-column
            label="详情"
            align="center"
          ><template slot-scope="scope"><el-button icon="el-icon-search" circle @click="heanleInfo(scope.row.ID)" /></template></el-table-column>
        </el-table>
        <el-pagination
          :current-page="appListQuery.page"
          :page-size="appListQuery.pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{float:'right',padding:'20px'}"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
          layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>
      </el-tab-pane>
    </el-tabs>
</div>
<div v-else="currentAppId != null">
    <router-view></router-view>
    </div>
    </div>
</template>

<script>
import { getNamespacesList } from '@/api/namespaces'
import { getBusinessList } from '@/api/business'
import { getSubBusinessListByBusinessId } from '@/api/subBusiness'
import { createApps, getAppsListByNamespaceId } from  '@/api/apps'
import { mapGetters } from 'vuex'
import { store } from '@/store/index'
import moment from 'moment'
export default {
    name:"Myapps",
    data () {
      return {
        test: getSubBusinessListByBusinessId,
        namespaces: [],
        total: ''||0,
        namespace: '',
        // current_app_id: this.currentAppId|| null,
        nsvalue: '' || 1,
        apps: [],
        appvalue: '',
        appListQuery: {
          page: ''||1,
          pageSize: ''||10,
          namespaceId: this.nsvalue || 1,
          ID: this.appvalue || ''
        },
        businessId: '',
        business: [],
        subbusinessId: '',
        subbusiness: [],
        subbusinessQuery: {
          businessId: this.businessId
        },
        rules: {
        container_port: [
          { required: true, message: '请输入容器端口', trigger: 'blur' },
          { required: true, pattern: /(^\d{2,5}$)/, message: '长度在 2 到 5 个数字', trigger: 'blur' }
        ],
        git_url: [
          { required: true, message: '请输入git地址', trigger: 'blur' },
          { required: true, pattern: /(^git@.:.\/.\.git$)/, message: '请输入正确的SSH格式的git地址', trigger: 'blur' }
        ],
        app_name: [
          // { required: true, message: '请输入应用名称', trigger: 'blur' }
        ],
        business_id: [
          { required: true, message: '请选择对应主业务线信息', trigger: 'blur' }
        ],
        subbusiness_id: [
          { required: true, message: '请选择对应次业务线信息', trigger: 'blur' }
        ],
        cpu_limit: [
          { required: true, message: '请选择CPU限额', trigger: 'blur' }
        ],
        cpu_requests: [
          { required: true, message: '请选择CPU需求', trigger: 'blur' }
        ],
        mem_limit: [
          { required: true, message: '请选择内存限额', trigger: 'blur' }
        ],
        mem_requests: [
          { required: true, message: '请选择内存需求', trigger: 'blur' }
        ],
        language: [
          { required: true, message: '请选择应用开发语言', trigger: 'blur' }
        ],
        repository: [
          { required: true, message: '请输入镜像地址', trigger: 'blur' }
        ]
      },
        formLabelWidth: '100px',
        dialogFormVisibleAdd: false,
        tableData: [],
        addData: {
          appName:'',
          businessId:'',
          createUser:'',
          currentEnv: ''||"test",
          gatewaySwitch:'' || false,
          gitUrl:'',
          hpaSwitch:'' || false,
          language:'',
          livenessSwitch:'' || false,
          metricsSwitch:'' || false,
          namespaceId:null,
          readinessSwitch:'' || false,
          pullPolicy: '' || 'IfNotPresent',
          subbusinessId:null,
          updateUser:null,
          usable:'' || true,
          containerPort: '',
          repository: '',
          commandSwitch: '' || false,
          argsSwitch: '' || false,
          cpuLimit: '' || '1',
          cpuRequests: '' || '1',
          memLimit: '' || '1Gi',
          memRequests: '' || '1Gi',
          argsInfo: '',
          commandInfo: '',
        },
        languages: [{
          value: 'python',
          label: 'python'
        }, {
          value: 'go',
          label: 'go'
        }, {
          value: 'java',
          label: 'java'
        }, {
          value: 'vue',
          label: 'vue'
        }, {
          value: 'php',
          label: 'php'
        }, {
          value: 'other',
          label: 'other'
        }],
        cpu: ['500m', '1', '2', '3', '4'],
        mem: ['500Mi', '1Gi', '2Gi', '3Gi', '4Gi'],
        image_pull: ['Always', 'Never', 'IfNotPresent']
        }
    },
    computed: {
      ...mapGetters('user', ['userInfo','currentAppId']),
    },
    watch: {
      businessId(val) {
        console.log("业务线ID有变化: ",val);
        this.subbusiness = []
        this.addData.subbusinessId = ''
        this.subbusinessQuery.businessId = val
        this.getSubBusinessData()
      },
      subbusinessId(val) {
       console.log("二级业务线ID有变化: ",val); 
      }, 
      nsvalue(val) {
        this.appListQuery.namespaceId = val
        this.appvalue = ''
        this.getAppsListData()
      },
      appvalue(val) {
        this.appListQuery.namespaceId = this.nsvalue
        this.appListQuery.ID = val
        this.getAppsListData() 
      },
    },
    beforeMount() {
      // this.nsvalue = 1
      this.getNamespacesListData()
      this.getBusinessData()
      this.getAppsListData()
    },
    methods: {
      dateFormat(row, column) {
        var date = row[column.property]
        if (date === undefined) {
          return ''
        }
        return moment(date).format('YYYY-MM-DD HH:mm:ss')
      },
      getNamespacesListData() {
        getNamespacesList().then(res => { 
        this.namespaces = res.data.list
        })
      },
      getBusinessData() {
        getBusinessList().then(res => {
          this.business = res.data.list
        })
      },
      getSubBusinessData() {
        getSubBusinessListByBusinessId(this.subbusinessQuery).then(res => {
          this.subbusiness = res.data.list
        })
      },
      getAppsListData() {
        getAppsListByNamespaceId(this.appListQuery).then(res => {
          this.tableData = res.data.list
          this.total = res.data.total
        })
      },
      changeNamespace(_value) {
        this.nsvalue = _value
        console.log("当前NSID:",this.nsvalue);
      },
      createAppData() {
        this.dialogFormVisibleAdd = false
        this.addData.createUser = this.userInfo.userName
        this.addData.updateUser = this.userInfo.userName
        this.addData.namespaceId = this.nsvalue
        console.log("请求创建应用: ", this.addData);
        createApps(this.addData).then(res => {
          console.log(res.code);
          this.getAppsListData()
          // this.resetAddData()
          this.$notify({
            title: '成功',
            message: '创建应用成功,请进行应用配置~',
            type: 'success'
          })
        })
      }, 
      changeBusiness(_value) {
        this.businessId = _value
      },
      changeSubBusiness(_value) {
        this.subbusinessId = _value
      },
      changeApp(_value) {
        this.appvalue = _value
     
      },
      heanleInfo(val) {
        store.commit('user/setCurrentAppId',val)
      },
      handleSizeChange(val) {
        this.appListQuery.pageSize = val
        this.getAppsListData()
      },
      handleCurrentChange(val) {
        this.appListQuery.page = val
        this.getAppsListData()
      }
    },
}
</script>
<style lang="scss">
    
</style>
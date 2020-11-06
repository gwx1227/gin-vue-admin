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
            <el-form :model="addData">
              <el-form-item label="应用名称" :required="true" :label-width="formLabelWidth">
                <el-input v-model="addData.appName" autocomplete="off" placeholder="您的应用名称,例如: my-book-info" />
              </el-form-item>
              <el-form-item label="git地址" :required="true" :label-width="formLabelWidth">
                <el-input v-model="addData.gitUrl" autocomplete="off" placeholder="您应用对应的git地址(注意:是ssh方式的),例如: git@github.com:gwx1227/helm.git" />
              </el-form-item>
              <el-form-item label="主业务线" :required="true" :label-width="formLabelWidth">
                <el-select v-model="addData.businessId" placeholder="请选择一级业务线" @change="changeBusiness">
                  <el-option
                    v-for="item in business"
                    :key="item.ID"
                    :label="item.business"
                    :value="item.ID"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="次业务线" :required="true" :label-width="formLabelWidth">
                <el-select v-model="addData.subbusinessId" placeholder="请选择二级业务线" @change="changeSubBusiness">
                  <el-option
                    v-for="item in subbusiness"
                    :key="item.ID"
                    :label="item.subbusiness"
                    :value="item.ID"
                  />
                </el-select>
              </el-form-item>
              <el-form-item label="开发语言" :required="true" :label-width="formLabelWidth">
                <el-select v-model="addData.language" placeholder="请选择应用开发语言">
                  <el-option
                    v-for="item in languages"
                    :key="item.value"
                    :label="item.label"
                    :value="item.label"
                  />
                </el-select>
              </el-form-item>
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
          subbusinessId:null,
          updateUser:null,
          usable:'' || true,
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
        createApps(this.addData).then(res => {
          console.log(res.code);
          this.getAppsListData()
        })
      }, 
      changeBusiness(_value) {
        this.businessId = _value
      },
      changeSubBusiness(_value) {
        this.addData.subbusinessId = _value
      },
      changeApp(_value) {
        this.appvalue = _value
     
      },
      heanleInfo(val) {
        console.log("当前应用ID1: ", this.currentAppId);
        console.log("当前应用ID2: ", val);
        // this.$router.push({
        //   name: '/layout/myapp/appinfo',
        //   params: {
        //     ID: val
        //   }
        // })
        // this.$store.dispatch('user/UpdateCurrentAppId',val)
        store.commit('user/setCurrentAppId',val)
        store.
        console.log("当前应用ID3: ", this.currentAppId) 
        // this.current_app_id = val
        // this.$store.dispatch('app/updateNsvalue', this.nsvalue)
        // this.$store.dispatch('app/updateAppvalue', val)
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
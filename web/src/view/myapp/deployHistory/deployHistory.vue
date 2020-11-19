<template>
    <div v-if="currentAppId !== null">
       <el-button type="primary" round @click="to_deploy()">部署</el-button>
    <el-table
      :data="deploy_hisory"
      stripe
      align="center"
    >
      <el-table-column
        prop="ID"
        label="部署序号"
        align="center"
      />
      <el-table-column
        prop="CreatedAt"
        label="部署时间"
        align="center"
        :formatter="dateFormat"
      />
      <el-table-column
        prop="opsUser"
        label="部署人员"
        align="center"
      />
      <el-table-column
        prop="result"
        label="部署结果"
        align="center"
      />
      <el-table-column
        fixed="right"
        label="操作"
        align="center"
      >
        <template slot-scope="scope">
          <el-button icon="el-icon-search" type="primary" circle @click="checkInfo(id=scope.row.ID)" />
          <el-dialog append-to-body title="部署详情" :visible.sync="dialogTableVisible">
            <template >
               <json-view :data="value" theme="one-dark" />
            </template>
          </el-dialog>

        </template>
      </el-table-column>
     

    </el-table>
        <el-pagination
          :current-page="query.page"
          :page-size="query.pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :style="{float:'right',padding:'20px'}"
          :total="total"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
          layout="total, sizes, prev, pager, next, jumper"
        ></el-pagination>
    </div>
</template>
<script>
import {
  getDeployHistoryList,
  findDeployHistory,
  createDeployHistory, 
} from "@/api/deployHistory";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import jsonView from 'vue-json-views';
import { mapGetters } from 'vuex'
import moment from 'moment'
export default {
  name: "DeployHistory",
  mixins: [infoList],
  components: { jsonView },
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getDeployHistoryData()
  },
  data() {
    return {
      dialogFormVisible: false,
      visible: false,
      type: "",
      total: ''||0,
      deleteVisible: false,
      query: {
        page: ''||1,
        pageSize: ''||10,
        appId: ''
      },
      dialogTableVisible: false,
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      deploy_hisory: [],
      jsonData: [],
      value: "",
      deploy_query: {
        ID: '',
      },
      deploy_data: {
        appId: '',
        opsUser: ''
      },
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
   dateFormat(row, column) {
      var date = row[column.property]
      if (date === undefined) {
        return ''
      }
      return moment(date).format('YYYY-MM-DD HH:mm:ss')
    },
    handleSizeChange(val) {
      this.query.pageSize = val
      this.getDeployHistoryData()
    },
    handleCurrentChange(val) {
      this.query.page = val
      this.getDeployHistoryData()
    },
    getDeployHistoryData() {
      this.query.appId = this.currentAppId
      getDeployHistoryList(this.query).then(res => {
        this.deploy_hisory = res.data.list
        this.total = res.data.total
      })
    },
    checkInfo(ID) {
      this.dialogTableVisible = true
      this.deploy_query.ID = ID
      const yaml = require('js-yaml')
      findDeployHistory(this.deploy_query).then(res => {
         this.value = yaml.safeLoad(res.data.redeployHistory.deployInfo)
      })
    },
    to_deploy() {
      this.deploy_data.appId = this.currentAppId
      this.deploy_data.opsUser = this.userInfo.userName 
      createDeployHistory(this.deploy_data).then(res => {
        this.getDeployHistoryData()
        this.$notify({
          title: '成功',
          message: '开始部署,请留意服务状态.',
          type: 'success'
        })
      })
    }
    
  },
}
</script>

<style >
</style>
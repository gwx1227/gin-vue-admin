<template>
    <div v-if="currentAppId !== null">
          <el-switch
            v-model="metricsSwitch"
            style="display: block"
            active-color="#13ce66"
            inactive-color="#ff4949"
            inactive-text="关"
            active-text="开"
            @change="change_switch"
          /><br>
           <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
       <el-button type="primary" round @click="openEdit()">调整监控数据配置</el-button>
    <el-dialog append-to-body title="修改监控接入配置" :visible.sync="dialogFormVisible">
      <el-form :model="editData">
        <el-form-item :label-width="formLabelWidth">
        </el-form-item>
        <el-form-item label="监控端口" :label-width="formLabelWidth">
          <el-input
            v-model="editData.port"
            placeholder="例如: 8888 "
            clearableautocomplete="off"
            style="width:200px"
          />
        </el-form-item>
        <el-form-item label="监控路径" :label-width="formLabelWidth">
          <el-input
            v-model="editData.path"
            placeholder="例如: /metrics "
            clearableautocomplete="off"
            style="width:200px"
            disabled
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="editMetricsData()">确 定</el-button>
      </div>
    </el-dialog>
           </el-row>
    <el-card>
      <el-form
        :model="metrics"
      >
        <el-form-item label="监控端口" style="margin-bottom: 5px;">
          <el-input
            v-model="metrics.port"
            placeholder="例如: 8081"
            clearable
            style="width:200px"
            disabled
          />

          <el-form-item label="监控路径" style="margin-bottom: 5px;">
            <el-input
              v-model="metrics.path"
              placeholder="/metrics"
              clearable
              style="width:200px"
              disabled
            />
          </el-form-item>
        </el-form-item></el-form>
    </el-card>
    </div>
</template>

<script>
import {
 findMetrics,
 updateMetrics, 
} from "@/api/metrics";  //  此处请自行替换地址
import {
  findApps,
  updateAppsSwitch,
} from '@/api/apps';
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Metrics",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getMetricsData()
    this.getSwitchData()
  },
  data() {
    return {
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      switchQuery: {
        ID: ''
      },
      query: {
        appId: ''
      },
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      metricsSwitch: '',
      metrics: {
        port: '',
        path: ''
      },
      editData: {
        appId: '',
        port: '',
        path: '' || '/metrics'
      },
      switchData: {
        ID: '',
        metricsSwitch: ''
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
    change_switch(value) {
      this.switchData.metricsSwitch = value
      this.switchData.ID = this.currentAppId
      updateAppsSwitch(this.switchData).then(res => {
        console.log("staus: ", res.code)
        this.getSwitchData()
        this.$notify({
          title: '成功',
          message: '调整就绪检查开关完成.',
          type: 'success'
        })
      })
    },
    getSwitchData() {
      this.switchQuery.ID = this.currentAppId
      findApps(this.switchQuery).then(res => {
        this.metricsSwitch = res.data.reapps[0].metricsSwitch
      })
    },
    getMetricsData() {
      this.query.appId = this.currentAppId
      findMetrics(this.query).then(res => {
        this.metrics.port = res.data.remetrics.port
        this.metrics.path = res.data.remetrics.path
      })
    },
    openEdit() {
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = true
    },
    editMetricsData() {
      updateMetrics(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.getMetricsData()
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '编辑监控接入信息完成.',
          type: 'success'
        })
      })
    }
  },
}
</script>

<style>
</style>
<template>
    <div v-if="currentAppId !== null">
          <el-switch
            v-model="readinessSwitch"
            style="display: block"
            active-color="#13ce66"
            inactive-color="#ff4949"
            inactive-text="关"
            active-text="开"
            @change="change_switch"
          /><br>
           <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
       <el-button type="primary" round @click="openEdit()">调整就绪检查参数</el-button>
    <el-dialog append-to-body title="修改健康检查配置" :visible.sync="dialogFormVisible">
      <el-form :model="editData">
        <el-form-item :label-width="formLabelWidth">
        </el-form-item>
        <el-form-item label="认定失败次数" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.failureThreshold" :min="2" :max="10" @change="failure_threshold_HandleChange" />
        </el-form-item>
        <el-form-item label="初始探测延迟" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.initialDelaySeconds" :min="2" :max="60" @change="initial_delay_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="探测间隔时间" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.periodSeconds" :min="2" :max="10" @change="period_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="认定成功次数" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.successThreshold" :min="1" :max="10" @change="success_threshold_HandleChange" />
        </el-form-item>
        <el-form-item label="探测超时时间" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.timeoutSeconds" :min="1" :max="10" @change="timeout_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="应用探测路径" style="margin-bottom: 5px;">
          <el-input
            v-model="editData.path"
            placeholder="例如: /healthcheck"
            clearable
            style="width:200px"
          />

        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="editReadinessData()">确 定</el-button>
      </div>
    </el-dialog>
           </el-row>
    <el-card>
      <el-form
        :model="readiness"
      >
        <el-form-item label="认定失败次数" style="margin-bottom: 5px;">
          <el-input-number v-model="readiness.failureThreshold" :min="2" :max="10" disabled @change="failure_threshold_HandleChange" />
        </el-form-item>
        <el-form-item label="初始探测延迟" style="margin-bottom: 5px;">
          <el-input-number v-model="readiness.initialDelaySeconds" :min="2" :max="10" disabled @change="initial_delay_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="探测间隔时间" style="margin-bottom: 5px;">
          <el-input-number v-model="readiness.periodSeconds" :min="2" :max="10" disabled @change="period_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="认定成功次数" style="margin-bottom: 5px;">
          <el-input-number v-model="readiness.successThreshold" :min="2" :max="10" disabled @change="success_threshold_HandleChange" />
        </el-form-item>
        <el-form-item label="探测超时时间" style="margin-bottom: 5px;">
          <el-input-number v-model="readiness.timeoutSeconds" :min="2" :max="10" disabled @change="timeout_seconds_HandleChange" />
        </el-form-item>
        <el-form-item label="应用探测路径" style="margin-bottom: 5px;">
          <el-input
            v-model="readiness.path"
            placeholder="例如: /healthcheck"
            clearable
            style="width:200px"
            disabled
          />

        </el-form-item>
      </el-form>
    </el-card>
    </div>
</template>

<script>
import {
  updateReadiness,
  findReadiness, 
} from "@/api/readiness";  //  此处请自行替换地址
import {
  findApps,
  updateAppsSwitch,
} from '@/api/apps';
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Readiness",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getReadinessData()
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
      readinessSwitch: '',
      readiness: {
        failureThreshold: '',
        initialDelaySeconds: '',
        path: '',
        periodSeconds: '',
        successThreshold: '',
        timeoutSeconds: ''
      },
      editData: {
        appId: '',
        failureThreshold: '' || 3,
        initialDelaySeconds: '' || 10,
        path: '' || '/healthcheck',
        periodSeconds: '' || 3,
        successThreshold: '' || 1,
        timeoutSeconds: '' || 2,
      },
      switchData: {
        ID: '',
        readinessSwitch: ''
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
      this.switchData.readinessSwitch = value
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
    getReadinessData() {
      this.query.appId = this.currentAppId
      findReadiness(this.query).then(res => {
        this.readiness.failureThreshold = res.data.rereadiness.failureThreshold
        this.readiness.initialDelay_seconds = res.data.rereadiness.initialDelaySeconds
        this.readiness.path = res.data.rereadiness.path
        this.readiness.periodSeconds = res.data.rereadiness.periodSeconds
        this.readiness.successThreshold = res.data.rereadiness.successThreshold
        this.readiness.timeoutSeconds = res.data.rereadiness.timeoutSeconds
      })
    },
    getSwitchData() {
      this.switchQuery.ID = this.currentAppId
      findApps(this.switchQuery).then(res => {
        this.readinessSwitch = res.data.reapps[0].readinessSwitch
      })
    },
    openEdit() {
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = true
    },
    editReadinessData() {
      updateReadiness(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.getReadinessData()
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '编辑就绪检查信息完成.',
          type: 'success'
        })
      })
    },
    failure_threshold_HandleChange(value) {
      this.editData.failureThreshold = value
    },
    initial_delay_seconds_HandleChange(value) {
      this.editData.initialDelaySeconds = value
    },
    period_seconds_HandleChange(value) {
      this.editData.periodSeconds = value
    },
    success_threshold_HandleChange(value) {
      this.editData.successThreshold = value
    },
    timeout_seconds_HandleChange(value) {
      this.editData.timeoutSeconds = value
    },
  },
}
</script>

<style>
</style>
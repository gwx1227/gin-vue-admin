<template>
    <div v-if="currentAppId !== null">
      
         <el-switch
            v-model="hpaSwitch"
            style="display: block"
            active-color="#13ce66"
            inactive-color="#ff4949"
            inactive-text="关"
            active-text="开"
            @change="change_switch"
          /><br>
          <el-button type="primary" round @click="openEdit()">编辑信息</el-button><br>
    <el-dialog append-to-body title="修改监控接入配置" :visible.sync="dialogFormVisible">
      <el-form :model="editData">
        <el-form-item :label-width="formLabelWidth">
        </el-form-item>
        <el-form-item label="最小副本数" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.minReplicas" :min="2" :max="10" @change="minHandleChange" />
        </el-form-item>
        <el-form-item label="最大副本数" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.maxReplicas" :min="2" :max="10" @change="maxHandleChange" />
        </el-form-item>
        <el-form-item label="CPU水位线" style="margin-bottom: 5px;">
          <el-input-number v-model="editData.cpuTargetAverageUtilization" :min="50" :max="200" @change="cpuHandleChange" />
        </el-form-item>
        <el-form-item label="内存水位线" style="margin-bottom: 5px;">
          <el-input
            v-model="editData.memTargetAverageValue"
            placeholder="例如: 2Gi"
            clearable
            style="width:200px"
          />

        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="editHpaData()">确 定</el-button>
      </div>
    </el-dialog>
    <el-card>
      <el-form
        :model="hpa"
      >
        <el-form-item label="最小副本数" style="margin-bottom: 5px;">
          <el-input-number v-model="hpa.minReplicas" :min="2" :max="10" disabled @change="minHandleChange" />
        </el-form-item>
        <el-form-item label="最大副本数" style="margin-bottom: 5px;">
          <el-input-number v-model="hpa.maxReplicas" :min="2" :max="10" disabled @change="maxHandleChange" />
        </el-form-item>
        <el-form-item label="CPU水位线" style="margin-bottom: 5px;">
          <el-input-number v-model="hpa.cpuTargetAverageUtilization" :min="50" :max="200" disabled @change="cpuHandleChange" />
        </el-form-item>
        <el-form-item label="内存水位线" style="margin-bottom: 5px;">
          <el-input
            v-model="hpa.memTargetAverageValue"
            placeholder="例如: 2Gi"
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
 updateHpa,
 findHpa, 
} from "@/api/hpa";  //  此处请自行替换地址
import {
  findApps,
  updateAppsSwitch,
} from '@/api/apps';
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Hpa",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getHpaData()
    this.getSwitchData()
  },
  data() {
    return {
      dialogFormVisible: false,
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      visible: false,
      type: "",
      deleteVisible: false,
      switchQuery: {
        ID: ''
      },
      query: {
        appId: ''
      },
      hpa: {
        minReplicas: '',
        maxReplicas: '',
        cpuTargetAverageUtilization: '',
        memTargetAverageValue: ''
      },
      hpaSwitch: '',
      editData: {
        appId: '',
        minReplicas: '' || 2,
        maxReplicas: '' || 10,
        cpuTargetAverageUtilization: '' || 70,
        memTargetAverageValue: ''
      },
      switchData: {
        ID: '',
        hpaSwitch: ''
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
      this.switchData.hpaSwitch = value
      this.switchData.ID = this.currentAppId
      updateAppsSwitch(this.switchData).then(res => {
        console.log("staus: ", res.code)
        this.getSwitchData()
        this.$notify({
          title: '成功',
          message: '调整自动扩展开关完成.',
          type: 'success'
        })
      })
    },
    getSwitchData() {
      this.switchQuery.ID = this.currentAppId
      findApps(this.switchQuery).then(res => {
        this.hpaSwitch = res.data.reapps[0].hpaSwitch
      })
    },
    getHpaData() {
      this.query.appId = this.currentAppId
      findHpa(this.query).then(res => {
        this.hpa.minReplicas = res.data.rehpa.minReplicas
        this.hpa.maxReplicas = res.data.rehpa.maxReplicas
        this.hpa.cpuTargetAverageUtilization = res.data.rehpa.cpuTargetAverageUtilization
        this.hpa.memTargetAverageValue = res.data.rehpa.memTargetAverageValue
      })
    },
    openEdit() {
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = true
    },
    editHpaData() {
      updateHpa(this.editData).then(res => {
        this.getHpaData()
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '编辑动态扩展信息完成.',
          type: 'success'
        })
      })
    },
    minHandleChange(value) {
      this.editData.minReplicas = value
    },
    maxHandleChange(value) {
      this.editData.maxReplicas = value
    },
    cpuHandleChange(value) {
      this.editData.maxReplicas = value
    },
  },
}
</script>

<style>
</style>
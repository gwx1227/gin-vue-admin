<template>
    <div v-if="currentAppId !== null">
      <el-switch
      v-model="gatewaySwitch"
      style="display: block"
      active-color="#13ce66"
      inactive-color="#ff4949"
      inactive-text="关"
      active-text="开"
      @change="change_gateway_switch"
    />
    <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
      <el-button type="primary" :disabled="!gatewaySwitch" round @click="dialogFormVisibleAdd = true">添加域名</el-button>
      <el-dialog append-to-body title="添加域名" :visible.sync="dialogFormVisibleAdd">
        <el-form :model="addData">
          <el-form-item label="域名" :label-width="formLabelWidth">
            <el-input v-model="addData.hosts" autocomplete="off" />
          </el-form-item>
          <el-form-item label="路径" :label-width="formLabelWidth">
            <el-input v-model="addData.paths" autocomplete="off" />
          </el-form-item>
          <el-form-item label="ONLINE权重" :label-width="formLabelWidth">
            <el-input-number v-model="addData.weightOnline" :min="0" :max="100" @change="weight_online_HandleChange" />
          </el-form-item>
          <el-form-item label="CANARY权重" :label-width="formLabelWidth">
            <el-input-number v-model="addData.weightCanary" :min="0" :max="100" @change="weight_canary_HandleChange" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisibleAdd = false">取 消</el-button>
          <el-button type="primary" @click="addGatewayData()">确 定</el-button>
        </div>
      </el-dialog>
    </el-row>
    <el-table
      :data="tableData"
      style="width: 100%"
    >
      <el-table-column
        prop="hosts"
        label="域名"
        align="center"
      />
      <el-table-column
        prop="paths"
        label="路径"
        align="center"
      />
      <el-table-column
        prop="weightOnline"
        label="ONLINE权重"
        align="center"
      />
      <el-table-column
        prop="weightCanary"
        label="CANARY权重"
        align="center"
      />
      <el-table-column
        fixed="right"
        label="操作"
        align="center"
      >
        <template slot-scope="scope">
          <template>
            <el-popconfirm
              confirm-button-text="好的"
              cancel-button-text="不用了"
              icon="el-icon-info"
              icon-color="red"
              :title="`确认要删除域名 ${scope.row.hosts} 下路径 ${scope.row.paths} 的配置?`"
              @confirm="deleteGatewayData(ID=scope.row.ID)"
            >
              <el-button slot="reference" type="danger" :disabled="!gatewaySwitch" icon="el-icon-delete" circle />
            </el-popconfirm>
          </template>

          <el-button icon="el-icon-edit" :disabled="!gatewaySwitch" type="primary" circle @click="openEdit(ID=scope.row.ID,hosts=scope.row.hosts,paths=scope.row.paths, weightOnline=scope.row.weightOnline, weightCanary=scope.row.weightCanary)" />

          <el-dialog append-to-body title="修改配置" :visible.sync="dialogFormVisible">
            <el-form :model="editData">
              <el-form-item label="HOST" :label-width="formLabelWidth">
                <el-input v-model="editData.hosts" disabled autocomplete="off" />
              </el-form-item>
              <el-form-item label="PATH" :label-width="formLabelWidth">
                <el-input v-model="editData.paths" autocomplete="off" />
              </el-form-item>
              <el-form-item label="ONLINE权重" :label-width="formLabelWidth">
                <el-input-number v-model="editData.weightOnline" :min="0" :max="100" @change="weight_online_HandleChange" />
              </el-form-item>
              <el-form-item label="CANARY权重" :label-width="formLabelWidth">
                <el-input-number v-model="editData.weightCanary" :min="0" :max="100" @change="weight_canary_HandleChange" />
              </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="editGatewayData()">确 定</el-button>
            </div>
          </el-dialog>

        </template>
      </el-table-column>
    </el-table>
    </div>
</template>

<script>
import {
 updateGateway,
 getGatewayList,
 deleteGateway,
 createGateway,
} from "@/api/gateway";  //  此处请自行替换地址
import {
  findApps,
  updateAppsSwitch,
} from '@/api/apps';
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Gateway",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getGatewayData()
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
      gatewaySwitch: '',
      query: {
        appId: ''
      },
      deleteData: {
        ID:'',
      },
      addData: {
        appId: '',
        weightCanary: '',
        weightOnline: '',
        hosts: '',
        paths: ''
      },
      editData: {
        appId: '',
        ID: '',
        weightCanary: '',
        weightOnline: '',
        hosts: '',
        paths: ''
      },
      switchData: {
        ID: '',
        gatewaySwitch: ''
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
    change_gateway_switch(value) {
      this.switchData.gatewaySwitch = value
      this.switchData.ID = this.currentAppId
      updateAppsSwitch(this.switchData).then(res => {
        console.log("staus: ", res.code)
        this.getSwitchData()
        this.$notify({
          title: '成功',
          message: '调整域名开关完成.',
          type: 'success'
        })
      })
    },
    getSwitchData() {
      this.switchQuery.ID = this.currentAppId
      findApps(this.switchQuery).then(res => {
        this.gatewaySwitch = res.data.reapps[0].gatewaySwitch
      })
    },
    getGatewayData() {
      this.query.appId = this.currentAppId
      getGatewayList(this.query).then(res => {
        this.tableData = res.data.list
      })
    },
    openEdit(ID,hosts, paths, weightOnline, weightCanary) {
      this.editData.ID = ID
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = true
      this.editData.hosts = hosts
      this.editData.paths = paths
      this.editData.weightOnline = weightOnline
      this.editData.weightCanary = weightCanary
    },
    resetGatewayData() {
      this.editData.ID = ''
      this.editData.hosts = ''
      this.editData.paths = ''
      this.editData.weightOnline = ''
      this.editData.weightCanary = ''
    },
    resetAddData() {
      this.addData.hosts = ''
      this.addData.paths = ''
      this.addData.weightCanary = ''
      this.addData.weightOnline = ''
    },
    addGatewayData() {
      this.dialogFormVisibleAdd = false
      this.addData.appId = this.currentAppId
      createGateway(this.addData).then(res => {
        console.log("staus: ", res.code)
        this.getGatewayData()
        this.resetAddData()
        this.$notify({
          title: '成功',
          message: '添加域名完成.',
          type: 'success'
        })
      })
    },
    editGatewayData() {
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = false
      updateGateway(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.resetGatewayData()
        this.getGatewayData()
        this.$notify({
          title: '成功',
          message: '编辑域名完成.',
          type: 'success'
        })
      })
    },
    weight_canary_HandleChange(value) {
      this.editData.weightOnline = 100-value
      this.addData.weightOnline = 100-value  
    },
    weight_online_HandleChange(value) {
      this.editData.weightCanary = 100-value 
      this.addData.weightCanary = 100-value 
    },
    deleteGatewayData(ID) {
      this.deleteData.ID = ID
      console.log("......",this.deleteData)
      deleteGateway(this.deleteData).then(res => {
        console.log("staus: ", res.code)
        this.getGatewayData()
        this.$notify({
          title: '成功',
          message: '删除域名下路径配置完成.',
          type: 'success'
        })
      })
    },
  }
}
</script>

<style>
</style>
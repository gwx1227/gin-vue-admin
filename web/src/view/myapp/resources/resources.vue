<template>
    <div v-if="currentAppId !== null">
       <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
         <el-button type="primary" round @click="openEdit()">调整资源使用配额</el-button>

    <el-dialog append-to-body title="修改资源配置" :visible.sync="dialogFormVisible">
      <el-form :model="editData">
        <el-form-item label="CPU限额" :label-width="formLabelWidth">
          <el-input
            v-model="editData.cpuLimit"
            placeholder="例如: 300m, 1 "
            clearableautocomplete="off"
            style="width:200px"
          />
        </el-form-item>
        <el-form-item label="CPU需求" :label-width="formLabelWidth">
          <el-input
            v-model="editData.cpuRequests"
            placeholder="例如: 300m, 1 "
            style="width:200px"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="内存限额" :label-width="formLabelWidth">
          <el-input
            v-model="editData.memLimit"
            placeholder="例如: 300Mi, 1Gi "
            style="width:200px"
            autocomplete="off"
          />
        </el-form-item>
        <el-form-item label="内存需求" :label-width="formLabelWidth">
          <el-input
            v-model="editData.memRequests"
            placeholder="例如: 300Mi, 1Gi"
            style="width:200px"
            autocomplete="off"
          />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="editResourcesData()">确 定</el-button>
      </div>
    </el-dialog>
       </el-row>
    <el-card>
      <el-form
        :model="resources"
      >
        <el-form-item label="CPU限额" style="margin-bottom: 5px;">
          <el-input
            v-model="resources.cpuLimit"
            style="width:200px"
            :disabled="resources_switch"
          />
        </el-form-item>
        <el-form-item label="CPU需求" style="margin-bottom: 5px;">
          <el-input
            v-model="resources.cpuRequests"
            style="width:200px"
            :disabled="resources_switch"
          />
        </el-form-item>
        <el-form-item label="内存限额" style="margin-bottom: 5px;">
          <el-input
            v-model="resources.memLimit"
            style="width:200px"
            :disabled="resources_switch"
          />
        </el-form-item>
        <el-form-item label="内存需求" style="margin-bottom: 5px;">
          <el-input
            v-model="resources.memRequests"
            style="width:200px"
            :disabled="resources_switch"
          />
        </el-form-item>
      </el-form>
    </el-card>
    </div>
</template>

<script>
import {
 findResources,
 updateResources, 
} from "@/api/resources";  //  此处请自行替换地址
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Resources",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getResourcesData()
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
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      resources_switch: true,
      resources: {
        cpuLimit: '',
        cpuRequests: '',
        memLimit: '',
        memRequests: ''
      },
      editData: {
        appId: '',
        cpuLimit: '',
        cpuRequests: '',
        memLimit: '',
        memRequests: ''
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
    getResourcesData() {
      this.query.appId = this.currentAppId 
      findResources(this.query).then(res => {
        this.resources.cpuLimit = res.data.reresources.cpuLimit
        this.resources.cpuRequests = res.data.reresources.cpuRequests
        this.resources.memLimit = res.data.reresources.memLimit
        this.resources.memRequests = res.data.reresources.memRequests
      })
      console.log(this.resources)
    },
    editResourcesData() {
      this.editData.appId = this.currentAppId
      updateResources(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.getResourcesData()
        this.dialogFormVisible = false
        this.$notify({
          title: '成功',
          message: '编辑资源信息完成.',
          type: 'success'
        })
      })
    },
    openEdit() {
      this.dialogFormVisible = true
    }, 
  },
}
</script>

<style>
</style>
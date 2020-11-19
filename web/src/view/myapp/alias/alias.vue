<template>
  <div v-if="currentAppId !== null">
 <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
      <el-button type="primary" round @click="dialogFormVisibleAdd = true">添加域名劫持</el-button>
      <el-dialog append-to-body title="添加配置" :visible.sync="dialogFormVisibleAdd">
        <el-form :model="addData">
          <el-form-item label="IP" :label-width="formLabelWidth">
            <el-input v-model="addData.ip" autocomplete="off" />
          </el-form-item>
          <el-form-item label="HOSTNAME" :label-width="formLabelWidth">
            <el-input v-model="addData.hostname" autocomplete="off" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisibleAdd = false">取 消</el-button>
          <el-button type="primary" @click="addAliasData()">确 定</el-button>
        </div>
      </el-dialog>
    </el-row>
    <el-table
      :data="alias_list"
      stripe
      align="center"
    >
      <el-table-column
        prop="ip"
        label="IP"
        align="center"
      />
      <el-table-column
        prop="hostname"
        label="HOSTNAME"
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
              :title="`确认要删除对域名 ${scope.row.hostname} 的劫持?`"
              @confirm="deleteAliasData(id=scope.row.ID)"
            >
              <el-button slot="reference" type="danger" icon="el-icon-delete" circle />
            </el-popconfirm>
          </template>

          <el-button icon="el-icon-edit" type="primary" circle @click="openEdit(ip=scope.row.ip,hostname=scope.row.hostname)" />

          <el-dialog append-to-body title="修改劫持" :visible.sync="dialogFormVisible">
            <el-form :model="editData">
              <el-form-item label="HOSTNAME" :label-width="formLabelWidth">
                <el-input v-model="editData.hostname" disabled autocomplete="off" />
              </el-form-item>
              <el-form-item label="IP" :label-width="formLabelWidth">
                <el-input v-model="editData.ip" autocomplete="off" />
              </el-form-item>

            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="editAliasData()">确 定</el-button>
            </div>
          </el-dialog>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
import {
    createAlias,
    deleteAlias,
    updateAlias,
    getAliasListByAppId,
} from "@/api/alias";  //  此处请自行替换地址
import { mapGetters } from 'vuex'
import { formatTimeToStr } from "@/utils/data";

export default {
  name: "Alias",
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  data() {
    return {
      alias_list: [],
      dialogFormVisible: false,
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      visible: false,
      type: "",
      addData: {
        ip: '',
        hostname: '',
        appId: '',
      },
      editData: {
        appId: '',
        ip: '',
        hostname: ''
      },
      deleteData: {
        id: ''|| 1000,
      },
      query: {
        appId: ''
      },
      deleteVisible: false,
      multipleSelection: [],formData: {
        appId:1,hostname:null,ip:null,
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
  beforeMount() {
    this.getAliasData()
  },
  methods: {
    getAliasData() {
      this.query.appId = this.currentAppId
      getAliasListByAppId(this.query).then(res => {
        this.alias_list = res.data.list
      })
    },
    deleteAliasData(id) {
      this.deleteData.id = id 
      
      deleteAlias(this.deleteData).then(res => {
        console.log("staus: ", res.code)
        this.getAliasData()
        // this.deleteData.id = ''
        this.$notify({
          title: '成功',
          message: '删除域名劫持信息完成.',
          type: 'success'
        })
      })
    }, 
    resetAddData() {
      this.addData.ip = ''
      this.addData.hostname = ''
    },
    resetEditData() {
      this.editData.ip = ''
      this.editData.hostname = ''
    },
    editAliasData() {
      this.editData.appId = this.currentAppId
      this.dialogFormVisible = false
      updateAlias(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.resetEditData()
        this.getAliasData()
        this.$notify({
          title: '成功',
          message: '编辑域名劫持信息完成.',
          type: 'success'
        })
      })
    },
    openEdit(ip, hostname) {
      this.dialogFormVisible = true
      this.editData.ip = ip
      this.editData.hostname = hostname
    },
    addAliasData() {
      this.addData.appId = this.currentAppId
      this.dialogFormVisibleAdd = false
      createAlias(this.addData).then(res => {
        console.log("staus: ", res.code)
        this.resetAddData()
        this.getAliasData()
        this.$notify({
          title: '成功',
          message: '添加域名劫持信息完成.',
          type: 'success'
        })
      })
    },
  },  
};
</script>

<style>
</style>
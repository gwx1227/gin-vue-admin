<template>
    <div v-if="currentAppId !== null">
      <el-row style="font-size: 26px;line-height: 26px;margin-bottom: 20px;padding:20px;">
      <el-button type="primary" round @click="dialogFormVisibleAdd = true">添加配置</el-button>
      <el-dialog append-to-body title="添加配置" :visible.sync="dialogFormVisibleAdd">
        <el-form :model="addData">
          <el-form-item label="KEY" :label-width="formLabelWidth">
            <el-input v-model="addData.key" autocomplete="off" />
          </el-form-item>
          <el-form-item label="VALUE" :label-width="formLabelWidth">
            <el-input v-model="addData.value" autocomplete="off" />
          </el-form-item>
        </el-form>
        <div slot="footer" class="dialog-footer">
          <el-button @click="dialogFormVisibleAdd = false">取 消</el-button>
          <el-button type="primary" @click="addConfigData()">确 定</el-button>
        </div>
      </el-dialog>
    </el-row>
    <el-table
      :data="tableData"
      stripe
      align="center"
    >
      <el-table-column
        prop="key"
        label="配置项"
        align="center"
      />
      <el-table-column
        prop="value"
        label="配置值"
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
              :title="`确认要删除配置 ${scope.row.key} ?`"
              @confirm="deleteConfigData(id=scope.row.ID)"
            >
              <el-button slot="reference" type="danger" icon="el-icon-delete" circle />
            </el-popconfirm>
          </template>

          <el-button icon="el-icon-edit" type="primary" circle @click="openEdit(id=scope.row.ID,key=scope.row.key,value=scope.row.value)" />

          <el-dialog append-to-body title="修改配置" :visible.sync="dialogFormVisible">
            <el-form :model="editData">
              <el-form-item label="KEY" :label-width="formLabelWidth">
                <el-input v-model="editData.key" disabled autocomplete="off" />
              </el-form-item>
              <el-form-item label="VALUE" :label-width="formLabelWidth">
                <el-input v-model="editData.value" autocomplete="off" />
              </el-form-item>
            </el-form>
            <div slot="footer" class="dialog-footer">
              <el-button @click="dialogFormVisible = false">取 消</el-button>
              <el-button type="primary" @click="editConfigData()">确 定</el-button>
            </div>
          </el-dialog>

        </template>
      </el-table-column>
    </el-table>
    </div>
</template>

<script>
import {
    createConfig,
    deleteConfig,
    updateConfig,
    getConfigListByAppId,
} from "@/api/config";  //  此处请自行替换地址
import infoList from "@/components/mixins/infoList";
import { mapGetters } from 'vuex'
export default {
  name: "Config",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  data() {
    return {
     visible: false,
      tableData: [],
      dialogFormVisible: false,
      dialogFormVisibleAdd: false,
      formLabelWidth: '120px',
      query: {
        appId: ''
      },
      deleteData: {
        id: '',
      },
      addData: {
        app_id: '',
        key: '',
        value: ''
      },
      editData: {
        appId: '',
        id: '',
        key: '',
        value: ''
      },
      };
  },
  beforeMount() {
    this.getConfigData()
  },
  methods: {
    getConfigData() { 
        this.query.appId = this.currentAppId
        getConfigListByAppId(this.query).then(res => {
        this.tableData = res.data.list
        })
    },
    openEdit(id, key, value) {
      this.dialogFormVisible = true
      this.editData.id = id
      this.editData.key = key
      this.editData.value = value
    },
    editConfigData() {
      this.dialogFormVisible = false
      this.editData.appId = this.currentAppId
      console.log("..",this.editData); 
      updateConfig(this.editData).then(res => {
        console.log("staus: ", res.code)
        this.resetEditData()
        this.getConfigData()
        this.$notify({
          title: '成功',
          message: '编辑配置项完成.',
          type: 'success'
        })
      })
    }, 
    deleteConfigData(id) {
      this.deleteData.id = id   
      console.log(this.deleteData)
      deleteConfig(this.deleteData).then(res => {
        console.log("staus: ", res.code)
        this.getConfigData()
        this.$notify({
          title: '成功',
          message: '删除配置项完成.',
          type: 'success'
        })
      })
    },
    resetAddData() {
      this.addData.appId = ''
      this.addData.key = ''
      this.addData.value = ''
    },
    resetEditData() {
      this.editData.id = ''
      this.editData.key = ''
      this.editData.value = ''
    },
    addConfigData() {
      this.dialogFormVisibleAdd = false
      this.addData.appId = this.currentAppId
      console.log('ADD DATA : ', this.addData)
      createConfig(this.addData).then(res => {
        console.log("staus: ", res.code)
        this.resetAddData()
        this.getConfigData() 
        this.$notify({
          title: '成功',
          message: '添加配置项完成.',
          type: 'success'
        })
      })
    },
  }
};
</script>

<style>
</style>
<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                              
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增apps表</el-button>
        </el-form-item>
        <el-form-item>
          <el-popover placement="top" v-model="deleteVisible" width="160">
            <p>确定要删除吗？</p>
              <div style="text-align: right; margin: 0">
                <el-button @click="deleteVisible = false" size="mini" type="text">取消</el-button>
                <el-button @click="onDelete" size="mini" type="primary">确定</el-button>
              </div>
            <el-button icon="el-icon-delete" size="mini" slot="reference" type="danger">批量删除</el-button>
          </el-popover>
        </el-form-item>
      </el-form>
    </div>
    <el-table
      :data="tableData"
      @selection-change="handleSelectionChange"
      border
      ref="multipleTable"
      stripe
      style="width: 100%"
      tooltip-effect="dark"
    >
    <el-table-column type="selection" width="55"></el-table-column>
    <el-table-column label="日期" width="180">
         <template slot-scope="scope">{{scope.row.CreatedAt|formatDate}}</template>
    </el-table-column>
    
    <el-table-column label="应用名称" prop="appName" width="120"></el-table-column> 
    
    <el-table-column label="businessId字段" prop="businessId" width="120"></el-table-column> 
    
    <el-table-column label="创建人" prop="createUser" width="120"></el-table-column> 
    
    <el-table-column label="当前应用环境" prop="currentEnv" width="120"></el-table-column> 
    
    <el-table-column label="域名开关" prop="gatewaySwitch" width="120">
         <template slot-scope="scope">{{scope.row.gatewaySwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="代码地址" prop="gitUrl" width="120"></el-table-column> 
    
    <el-table-column label="自动扩展开关" prop="hpaSwitch" width="120">
         <template slot-scope="scope">{{scope.row.hpaSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="应用语言" prop="language" width="120"></el-table-column> 
    
    <el-table-column label="健康检查开关" prop="livenessSwitch" width="120">
         <template slot-scope="scope">{{scope.row.livenessSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="监控接入开关" prop="metricsSwitch" width="120">
         <template slot-scope="scope">{{scope.row.metricsSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="namespaceId字段" prop="namespaceId" width="120"></el-table-column> 
    
    <el-table-column label="就绪检查开关" prop="readinessSwitch" width="120">
         <template slot-scope="scope">{{scope.row.readinessSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="subbusinessId字段" prop="subbusinessId" width="120"></el-table-column> 
    
    <el-table-column label="更新人" prop="updateUser" width="120"></el-table-column> 
    
    <el-table-column label="服役状态标志位" prop="usable" width="120">
         <template slot-scope="scope">{{scope.row.usable|formatBoolean}}</template>
    </el-table-column>
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateApps(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteApps(scope.row)">确定</el-button>
            </div>
            <el-button type="danger" icon="el-icon-delete" size="mini" slot="reference">删除</el-button>
          </el-popover>
        </template>
      </el-table-column>
    </el-table>

    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[10, 30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :visible.sync="dialogFormVisible" title="弹窗操作">
      此处请使用表单生成器生成form填充 表单默认绑定 formData 如手动修改过请自行修改key
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
    createApps,
    deleteApps,
    deleteAppsByIds,
    updateApps,
    findApps,
    getAppsList
} from "@/api/apps";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Apps",
  mixins: [infoList],
  data() {
    return {
      listApi: getAppsList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        appName:null,businessId:null,createUser:null,currentEnv:null,gatewaySwitch:null,gitUrl:null,hpaSwitch:null,language:null,livenessSwitch:null,metricsSwitch:null,namespaceId:null,readinessSwitch:null,subbusinessId:null,updateUser:null,usable:null,
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
      //条件搜索前端看此方法
      onSubmit() {
        this.page = 1
        this.pageSize = 10          
        if (this.searchInfo.gatewaySwitch==""){
          this.searchInfo.gatewaySwitch=null
        }         
        if (this.searchInfo.hpaSwitch==""){
          this.searchInfo.hpaSwitch=null
        }         
        if (this.searchInfo.livenessSwitch==""){
          this.searchInfo.livenessSwitch=null
        }        
        if (this.searchInfo.metricsSwitch==""){
          this.searchInfo.metricsSwitch=null
        }         
        if (this.searchInfo.readinessSwitch==""){
          this.searchInfo.readinessSwitch=null
        }          
        if (this.searchInfo.usable==""){
          this.searchInfo.usable=null
        }      
        this.getTableData()
      },
      handleSelectionChange(val) {
        this.multipleSelection = val
      },
      async onDelete() {
        const ids = []
        this.multipleSelection &&
          this.multipleSelection.map(item => {
            ids.push(item.ID)
          })
        const res = await deleteAppsByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateApps(row) {
      const res = await findApps({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.reapps;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          appName:null,
          businessId:null,
          createUser:null,
          currentEnv:null,
          gatewaySwitch:null,
          gitUrl:null,
          hpaSwitch:null,
          language:null,
          livenessSwitch:null,
          metricsSwitch:null,
          namespaceId:null,
          readinessSwitch:null,
          subbusinessId:null,
          updateUser:null,
          usable:null,
      };
    },
    async deleteApps(row) {
      this.visible = false;
      const res = await deleteApps({ ID: row.ID });
      if (res.code == 0) {
        this.$message({
          type: "success",
          message: "删除成功"
        });
        this.getTableData();
      }
    },
    async enterDialog() {
      let res;
      switch (this.type) {
        case "create":
          res = await createApps(this.formData);
          break;
        case "update":
          res = await updateApps(this.formData);
          break;
        default:
          res = await createApps(this.formData);
          break;
      }
      if (res.code == 0) {
        this.$message({
          type:"success",
          message:"创建/更改成功"
        })
        this.closeDialog();
        this.getTableData();
      }
    },
    openDialog() {
      this.type = "create";
      this.dialogFormVisible = true;
    }
  },
  async created() {
    await this.getTableData();}
};
</script>

<style>
</style>
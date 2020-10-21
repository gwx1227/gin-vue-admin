<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">                            
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog" type="primary">新增deploy表</el-button>
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
    
    <el-table-column label="关联应用id" prop="appId" width="120"></el-table-column> 
    
    <el-table-column label="args内容" prop="argsInfo" width="120"></el-table-column> 
    
    <el-table-column label="是否开启args" prop="argsSwitch" width="120">
         <template slot-scope="scope">{{scope.row.argsSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="command内容" prop="commandInfo" width="120"></el-table-column> 
    
    <el-table-column label="是否开启command" prop="commandSwitch" width="120">
         <template slot-scope="scope">{{scope.row.commandSwitch|formatBoolean}}</template>
    </el-table-column>
    
    <el-table-column label="容器port" prop="containerPort" width="120"></el-table-column> 
    
    <el-table-column label="拉取镜像策略" prop="pullPolicy" width="120"></el-table-column> 
    
    <el-table-column label="canary副本数" prop="replicaCountCanary" width="120"></el-table-column> 
    
    <el-table-column label="online副本数" prop="replicaCountOnline" width="120"></el-table-column> 
    
    <el-table-column label="docker镜像地址" prop="repository" width="120"></el-table-column> 
    
    <el-table-column label="canary代码版本" prop="tagCanary" width="120"></el-table-column> 
    
    <el-table-column label="online代码版本" prop="tagOnline" width="120"></el-table-column> 
    
    <el-table-column label="canary流量权重" prop="weigitCanary" width="120"></el-table-column> 
    
    <el-table-column label="online流量权重" prop="weigitOnline" width="120"></el-table-column> 
    
      <el-table-column label="按钮组">
        <template slot-scope="scope">
          <el-button @click="updateDeploy(scope.row)" size="small" type="primary">变更</el-button>
          <el-popover placement="top" width="160" v-model="scope.row.visible">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin: 0">
              <el-button size="mini" type="text" @click="scope.row.visible = false">取消</el-button>
              <el-button type="primary" size="mini" @click="deleteDeploy(scope.row)">确定</el-button>
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
    createDeploy,
    deleteDeploy,
    deleteDeployByIds,
    updateDeploy,
    findDeploy,
    getDeployList
} from "@/api/deploy";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";

export default {
  name: "Deploy",
  mixins: [infoList],
  data() {
    return {
      listApi: getDeployList,
      dialogFormVisible: false,
      visible: false,
      type: "",
      deleteVisible: false,
      multipleSelection: [],formData: {
        appId:null,argsInfo:null,argsSwitch:null,commandInfo:null,commandSwitch:null,containerPort:null,pullPolicy:null,replicaCountCanary:null,replicaCountOnline:null,repository:null,tagCanary:null,tagOnline:null,weigitCanary:null,weigitOnline:null,
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
        if (this.searchInfo.argsSwitch==""){
          this.searchInfo.argsSwitch=null
        }         
        if (this.searchInfo.commandSwitch==""){
          this.searchInfo.commandSwitch=null
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
        const res = await deleteDeployByIds({ ids })
        if (res.code == 0) {
          this.$message({
            type: 'success',
            message: '删除成功'
          })
          this.deleteVisible = false
          this.getTableData()
        }
      },
    async updateDeploy(row) {
      const res = await findDeploy({ ID: row.ID });
      this.type = "update";
      if (res.code == 0) {
        this.formData = res.data.redeploy;
        this.dialogFormVisible = true;
      }
    },
    closeDialog() {
      this.dialogFormVisible = false;
      this.formData = {
        
          appId:null,
          argsInfo:null,
          argsSwitch:null,
          commandInfo:null,
          commandSwitch:null,
          containerPort:null,
          pullPolicy:null,
          replicaCountCanary:null,
          replicaCountOnline:null,
          repository:null,
          tagCanary:null,
          tagOnline:null,
          weigitCanary:null,
          weigitOnline:null,
      };
    },
    async deleteDeploy(row) {
      this.visible = false;
      const res = await deleteDeploy({ ID: row.ID });
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
          res = await createDeploy(this.formData);
          break;
        case "update":
          res = await updateDeploy(this.formData);
          break;
        default:
          res = await createDeploy(this.formData);
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
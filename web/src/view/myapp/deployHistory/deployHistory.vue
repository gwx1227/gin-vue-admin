<template>
    <div v-if="currentAppId !== null">

    </div>
</template>
<script>
import {
   
} from "@/api/deployHistory";  //  此处请自行替换地址
import { formatTimeToStr } from "@/utils/data";
import infoList from "@/components/mixins/infoList";
import { mapGetters } from 'vuex'
export default {
  name: "DeployHistory",
  mixins: [infoList],
  computed: {
    ...mapGetters('user', ['userInfo','currentAppId']),
  },
  beforeMount() {
    this.getDeployData()
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
   getDeployData(){
     this.query.appId = this.currentAppId
     

   }, 
  },
}
</script>

<style>
</style>
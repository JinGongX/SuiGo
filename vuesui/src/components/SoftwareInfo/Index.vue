<template>
<el-row class="dv-software" :gutter="24" style="width: 86%;margin: 10px auto;">
    <template>
  <el-skeleton  :loading="loading" animated :count="6" >
 <el-col :span="4"  slot="template">
   <el-skeleton-item variant="image" rows="4" style="width: 40px; height: 40px;    border-radius: 50%;" />
   <div style="padding: 4px;line-height: 30px;">
     <el-skeleton-item variant="text" style="width: 30%" />
     <div
       style="display: flex; align-items: center; justify-items: space-between;"
     >
       <el-skeleton-item variant="text" style="width: 30%;" />
     </div>
   </div>
 </el-col>
</el-skeleton> </template>
    <el-col v-if="items.length==0"><el-empty description="暂无相关信息"></el-empty></el-col>
   <!-- <el-col :span="4" :key="index" v-for="(item,index) of items"><div class="grid-content bg-purple">
    <div class="dv-software-img" >
      <el-skeleton-item v-if="item.photoimg==null" variant="image" style="width: 100%; height: 144px" /> 
      <img v-else-if="item.photoimg!=null" :src="item.photoimg" style="width: 100%; " />
   </div>
   <div  class="dv-software-title"> <span class="sp-label">{{item.title}}</span></div>
    <el-button size="mini" round @click="downfiles(item.id)" v-if="item.filesimg!=null" >下载</el-button>
    <el-button size="mini" round v-else-if="item.filesimg==null" disabled>下载</el-button></div></el-col>
  </el-row> -->

  <!-- <el-row :gutter="24" style="width: 86%;margin: 10px auto;"> -->
    <!-- <el-col :span="6" style="margin-top: 10px;">
      <el-card shadow="always"> Always </el-card>
    </el-col> -->
    <el-col :span="6"  :key="index"  v-for="(item, index) in items"
       style="margin-top: 16px">
      <el-card shadow="hover" :body-style="{ padding: '0px' }">
        <div class="sf_card_title_img">
          <img
          :src="item.photoimg"
          class="image"
          style="min-height: 65px;"
        />
        </div>
        <div class="sf_card_title_conter">
        <div class="sf_card_title_name">
          <!-- <el-badge value="new" :hidden="item.label==0?true:false" class="item" type="warning">
             {{item.title}}
         </el-badge> -->
         {{item.title}}
         <el-tag class="ml-2" type="info" size="small" round>{{ item.label }}</el-tag>
        </div>
        <el-tooltip
        class="box-item"
        effect="dark"
        :content="item.desc"
        placement="bottom"
      >
        <div class="sf_card_title_desc">  
            {{item.desc}} 
        </div>
      </el-tooltip>
      </div>
      <div class="sf_card_title_down" @click="downfiles(item.id)" title="下载"> 
        <el-icon :size="24" ><Download />
        </el-icon>
      </div>
         </el-card>
    </el-col> 
  </el-row>
</template>

<script> 
import axios from 'axios' 
import moment from "moment"
moment.locale('zh-cn')

export default {
    data () {
    return {
      isCollapse: true,
      input: '',
      items: [],
      totals: 1,
      pagesize: 30,
      page: 1,
      menuid: 0,
      loading: true
    }
  },
  mounted() {
    axios({
         method: "get",
         url: this.$globalHost+"api/v1/softinfo", //"http://127.0.0.1:2023/api/v1/softinfo",
         headers:{
            "token":sessionStorage.token
         }
        })
        .then((res) => {
          //console.log(res)
          console.log(res.data.data)
          this.loading = false
          this.items = res.data.data.lists
          console.log(this.items)
          //this.totals = response.body.count // this.items.length
          //this.page = response.body.totalpages
          //console.log(response.body.TotalPages)
        }) 
  },
  methods: {
    
    // handselect (key, keyPath) {
    //   console.log(key, keyPath)
    //   this.menuid = key
    //   this.page = 1
    //   this.getdata()
    // },
    // handleSizeChange (val) {
    //   console.log('ddd当前页' + val)
    // },
    // handleCurrentChange (val) {
    //   console.log('当前页' + val)
    //   this.$http.get('/api/v1/softinfo?pageIndex=' + val + '&pageSize=' + this.pagesize).then((response) => {
    //     console.log(response)
    //     this.items = response.body.data
    //     this.totals = response.body.count // this.items.length
    //     this.page = response.body.totalpages
    //     console.log(response.body.TotalPages)
    //   }, (response) => {
    //     console.error(response)
    //   })
    // },
    getdata () {
      this.$http.get('/api/v1/softinfo?pageIndex=' + this.page + '&pageSize=' + this.pagesize + '&menuid=' + this.menuid).then((response) => {
        console.log(response)
        this.loading = false
        this.items = response.body.data
        this.totals = response.body.count // this.items.length
        this.page = response.body.totalpages
        console.log(response.body.TotalPages)
      }, (response) => {
        console.error(response)
      })
    },
    downfiles (uid) {
      var newWeb = window.open('_blank')
      newWeb.location.href =this.$globalHost+'downSoftfile?id='+ uid  //'http://127.0.0.1:2023/downSoftfile?id='+ uid //'/api/OptionTakes/' + uid
      // this.$http.get('/api/OptionTakes/' + uid).then((response) => {
      //   console.log(response.body)
      //   this.download(response.body, '66.jpg')
      // }, (response) => {
      //   console.error(response)
      // })
    },
    download (data, fileName) {
      if (!data) {
        return
      }
      let url = window.URL.createObjectURL(new Blob([data]))
      let link = document.createElement('a')
      link.style.display = 'none'
      link.href = url
      link.setAttribute('download', fileName)
      document.body.appendChild(link)
      link.click()
    },
    downfile (fileurl) {
      if (fileurl != null) {
        console.log('test down' + fileurl)
        var newWeb = window.open('_blank')
        newWeb.location.href = fileurl
      } else {
        this.open4()
      }
    },
    open4 () {
      this.$message({
        showClose: true,
        message: '错了哦，这是一条错误消息',
        type: 'error'
      })
    } 
  },
  // mounted () {
  //   //this.$http.get('/api/v1/softinfo?pageIndex=' + this.page + '&pageSize=' + this.pagesize + '&menuid=' + this.menuid).then((response) => {
    
  //     this.$http.get('/api/v1/softinfo').then((response) => {  console.log(response)
  //     this.loading = false
  //     this.items = response.body.data
  //     this.totals = response.body.count // this.items.length
  //     this.page = response.body.totalpages
  //     console.log(response.body.TotalPages)
  //   }, (response) => {
  //     console.error(response)
  //   })
  // }
}
 
</script>
 
<style>
/* .dv-software-img{
  width:60px;
  height: 60px;
  border-radius: 50%;
  border: 1px solid #6ab985;
  text-align: center;
  margin:0 auto;
} */
.dv-software-title{
  text-align: center;
}
.grid-content{
  text-align: center;
}
.sf_card_title_img{
  float:left;
  width:65px;
  min-height:65px;
  max-height:65px;
  border-right: 1px solid #e1dfdf;
}
.sf_card_title_conter{
  float:left;
  padding-left:10px;
  width: 60%;
}
.sf_card_title_name{
  float:left;
  /* padding-left:10px; */
  width: 70%;
  min-width:200px;
  font-size: 18px;
  font-weight: 600;
  /* padding-top:10px; */
  padding-top:4px;
}
.sf_card_title_desc{
  float:left;
  /* padding-left:10px; */
  /* width: 80%; */
  width: 70%;
  min-width:200px;
  font-size: 13px;
  overflow: hidden;
  text-overflow: ellipsis;
  word-break: break-all;
  display: -webkit-box;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;
}
.sf_card_title_down{
  float:left;
  padding-left:10px;
  padding-top: 14px;
  padding-right: 4px;
  display: block;
  cursor: pointer;
}
</style>
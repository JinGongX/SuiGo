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
        
        <el-col  :key="index"  v-for="(item, index) in items"
           style="margin-top: 16px">
          <el-card shadow="hover" :body-style="{ padding: '0px' }" style="height: 80px;">
            <div class="bksf_card_title_img">
              <img
              :src="item.photoimg"
              class="image"
              style="min-height: 65px;"
            />
            </div>
            <div class="bksf_card_title_conter">
            <div class="bksf_card_title_name">
              <!-- <el-badge value="new" :hidden="item.label==0?true:false" class="item" type="warning">
                 {{item.title}}
             </el-badge> -->
             {{item.title}}
             <el-tag class="ml-2" type="info" size="small" round>作者：{{ item.label }}</el-tag>
            </div>
            <el-tooltip
            class="box-item"
            effect="dark"
            :content="item.desc"
            placement="bottom"
          >
            <div class="bksf_card_title_desc">  
                {{item.desc}} 
            </div>
          </el-tooltip>
          </div>
          <div class="bksf_card_title_read" title="查看" >  
            <router-link :to="{path:'/Openpdf',query:{namepdf:item.filesbook}}"  target="_blank" rel="opener">
              <el-icon :size="30" ><Reading />
              </el-icon>
             </router-link>
                 
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
             url: this.$globalHost+"api/v1/Bookinfo", //"http://127.0.0.1:2023/api/v1/Bookinfo",
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
          this.$http.get('/api/v1/Bookinfo?pageIndex=' + this.page + '&pageSize=' + this.pagesize + '&menuid=' + this.menuid).then((response) => {
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
          newWeb.location.href ='http://127.0.0.1:2023/downSoftfile?id='+ uid //'/api/OptionTakes/' + uid
          // this.$http.get('/api/OptionTakes/' + uid).then((response) => {
          //   console.log(response.body)
          //   this.download(response.body, '66.jpg')
          // }, (response) => {
          //   console.error(response)
          // })
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
    .bksf_card_title_img{
      float:left;
      width:80px;
      min-height:80px;
      max-height:80px;
      border-right: 1px solid #e1dfdf;
    }
    .bksf_card_title_conter{
      float:left;
      padding-left:10px;
      width: 60%;
      padding-top: 4px;
    }
    .bksf_card_title_name{
      float:left;
      /* padding-left:10px; */
      width: 70%;
      min-width:200px;
      font-size: 18px;
      font-weight: 600;
      /* padding-top:10px; */
      padding-top:4px;
    }
    .bksf_card_title_desc{
      padding-top:6px;
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
    .bksf_card_title_read{
      float:right;
      padding-left:10px;
      padding-top: 24px;
      padding-right: 14px;
      display: block;
      cursor: pointer;
    }
    </style>
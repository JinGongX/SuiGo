<template>
  <el-row style="width: 86%; margin: -10px auto 36px auto;text-align: center;">
     <div class="my_bk_head">
       <img src="../../assets/myinfo/66d.jpg" />
     </div>
     <div class="my_bk_body">
        <div class="my_bk_body_photo">
          <img alt="photo"  :src=suiphoto />
        </div>
        <div class="my_bk_body_name">
          <span id="lbl_nameb">{{ suiname }}</span>
        </div>
        <div class="my_bk_body_left">
          <router-link to="./Blogs">
             <div class="idx_mymenu_top_logo" >
               <svg t="1724676707084" class="icon" viewBox="0 0 1152 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1475" width="32" height="32"><path d="M583.807877 778.756066c-84.099572 0-152.514558-70.908769-152.514558-158.092758s68.414986-158.059945 152.514558-158.059945 152.481745 70.908769 152.481744 158.092758-68.414986 158.059945-152.481744 158.059945z m0-261.453527c-53.944478 0-97.848195 46.364688-97.848196 103.360769s43.903717 103.393582 97.848196 103.393583 97.68413-46.364688 97.68413-103.36077-43.772466-103.393582-97.68413-103.393582z" fill="#ffffff" p-id="1476"></path><path d="M234.284411 1023.967187a109.16866 109.16866 0 0 1-109.16866-109.004596v-607.03944H179.814926v607.03944a54.403859 54.403859 0 0 0 54.338233 54.338234l685.462373-3.084417a54.436672 54.436672 0 0 0 54.469485-54.371046V311.073194h54.699176v600.772168a109.16866 109.16866 0 0 1-109.037409 109.16866z" fill="#ffffff" p-id="1477"></path><path d="M27.365994 419.060588a27.333181 27.333181 0 0 1-15.258018-50.006924L530.191528 18.873939a108.676466 108.676466 0 0 1 123.409478 0.689072l487.469077 337.251425a27.346306 27.346306 0 0 1-31.10667 44.986544L622.42871 64.516742a54.371046 54.371046 0 0 0-61.688332-0.328129L42.656826 414.368337a27.365994 27.365994 0 0 1-15.290832 4.692251z" fill="#ffffff" p-id="1478"></path></svg>          
             </div> 
          </router-link> 
        </div>
        <div class="my_bk_body_right">
          <el-tag   type="success"  effect="plain"   size="large" round> <svg t="1724769599336" class="icon" viewBox="0 0 1026 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4261" width="12" height="12"><path d="M1019.109859 384c-11.286261-32.01113-39.713391-55.05113-74.195478-60.126609L701.69212 288.233739l-105.627826-216.019478c-15.270957-31.276522-48.261565-51.489391-84.057043-51.489391-35.706435 0-68.652522 20.21287-83.968 51.489391l-105.672348 216.041739L79.166902 323.940174c-34.504348 4.964174-62.953739 27.981913-74.24 60.17113-11.264 32.50087-2.871652 67.806609 21.882435 92.137739l178.509913 175.638261-41.405217 243.378087c-5.810087 33.925565 9.282783 68.719304 38.555826 88.687304 28.627478 19.255652 67.005217 21.370435 97.836522 5.164522l211.745391-112.39513 211.878957 112.417391c13.712696 7.234783 29.094957 11.063652 44.521739 11.063652 19.010783 0 37.420522-5.609739 53.337043-16.317217 29.139478-19.878957 44.210087-54.650435 38.4-88.576l-41.382957-243.400348 178.532174-175.638261C1022.048294 451.917913 1030.440641 416.589913 1019.109859 384zM966.062207 444.527304l-195.094261 191.955478 45.278609 266.329043c2.938435 17.096348-4.585739 34.05913-19.478261 44.232348-15.248696 10.24-35.817739 11.330783-52.045913 2.782609L512.00725 826.323478l-232.537043 123.436522c-16.406261 8.637217-36.997565 7.479652-52.045913-2.671304-15.048348-10.262261-22.572522-27.247304-19.634087-44.343652l45.30087-266.284522-195.072-191.955478c-12.377043-12.154435-16.606609-29.718261-11.063652-45.723826 5.765565-16.384 20.524522-28.182261 38.622609-30.786783l266.48487-39.112348 115.97913-237.122783c7.880348-16.11687 25.154783-26.534957 43.987478-26.534957 18.899478 0 36.173913 10.395826 44.054261 26.512696l115.95687 237.122783 266.418087 39.023304c18.075826 2.671304 32.901565 14.514087 38.64487 30.809043C982.668815 414.786783 978.43925 432.328348 966.062207 444.527304zM509.046555 376.898783c-45.590261 0-82.320696 13.913043-109.122783 41.316174-44.744348 45.746087-43.78713 112.194783-43.720348 115.021913 0.26713 12.109913 10.173217 21.726609 22.238609 21.726609 0.133565 0 0.289391 0 0.422957 0 12.265739-0.222609 22.016-10.373565 21.837913-22.639304 0-0.512-0.400696-51.066435 31.254261-83.18887 18.098087-18.387478 44.054261-27.692522 77.06713-27.692522 12.288 0 22.26087-9.97287 22.26087-22.26087S521.334555 376.898783 509.046555 376.898783z" fill="#f4ea2a" p-id="4262"></path></svg>
             总发表：{{this.total}} 篇随记 </el-tag>
            <router-link  to="./Mkdown" target="_blank" rel="opener"> 
             <div class="my_bk_mkdown">
               <el-icon title="发表随记"  :size="22" style="cursor:pointer"><Promotion /></el-icon> 
             </div>
          </router-link> 
        </div>
     </div>
  </el-row>
    <!-- <el-row style="width: 86%;margin: 6px auto;text-align: center;">
      <el-col :span="24"><el-button-group class="ml-4">
        <router-link to="./Blogs"><el-button type="success" :icon="Files" >全部随记</el-button></router-link>
         <router-link to="./myBlogs"> <el-button type="success"   >我的随记</el-button></router-link>
       </el-button-group>
      </el-col>
    </el-row> -->
    <el-row style="width: 86%;margin: 0 auto;">
    
  <el-skeleton style="width: 240px" :loading="loading" animated :count="3" >
    <template #template> 
      <el-skeleton-item  variant="image" style="width: 240px; height: 240px" />
      <div style="padding: 14px">
        <el-skeleton-item variant="p" style="width: 50%" />
        <div
          style="
            display: flex;
            align-items: center;
            justify-items: space-between;
          "
        >
          <el-skeleton-item variant="text" style="margin-right: 16px" />
          <el-skeleton-item variant="text" style="width: 30%" />
        </div>
      </div> 
    </template>
  </el-skeleton> 
  <el-space wrap   style="margin-bottom: 0;width: 100%;display: grid;grid-template-columns: repeat(auto-fill, minmax(260px, 1fr));gap: 30px;" >
 
    <el-card  class="box-card" v-for="(o, index) in  blogData "   :key="o" :offset="index > 0 ? index%4 == 0? 0:1 :  0"   :body-style="{ padding: '0px' }"> 
      
          <img v-if="blogData[index].imgcover!=''"
            :src=blogData[index].imgcover
            class="image"
          />
          <el-skeleton-item v-else variant="image" style="width: 100%; height: 144px" />
          <!-- <img v-else class="image"  
             src="../../assets/sysimg/Emay.png"
             
          /> -->
          <!-- "https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" -->
          <div style="padding: 14px">
            <h3 class="blog-title" :title=blogData[index].title > <!--{{ blogData[index].account.name }} -->
              <router-link :to="{path:'/MyDetail',query:{id:blogData[index].id}}">{{ blogData[index].title }}</router-link>
            </h3>
            <span class="blog-overview">
              <!-- {{ message.substring(0,51)+"..." }} -->
              {{ blogData[index].desc}}
            </span>
           
            <div class="bottom">
              <img alt="photo" class="photo" :title=blogData[index].account.name :src=blogData[index].account.photo />
              
              <time class="time">{{ this.moments(blogData[index].created_on*1000).format("YYYY-MM-DD HH:mm") }}</time>
              <span style="font-size: 13px;color:#999;">阅读 {{ blogData[index].views }} </span>
              <!-- <router-link :to="{path:'/MyDetail',query:{id:blogData[index].id}}"  ><el-button text class="button"><el-icon><View /></el-icon>查看详细</el-button></router-link> -->
            </div>
            <div class="foot_btn"> 
              <div class="crard_tab"><el-tag effect="plain" type="warning" style="margin-right: 4px">{{ blogData[index].tag.name }}</el-tag>
                <el-tag   type="success"  effect="plain"   v-if="blogData[index].state==2" > <el-icon><Lock /></el-icon>仅我可见</el-tag>
                <el-tag   type="success"  effect="plain"   v-else> <el-icon><Unlock /></el-icon>公开</el-tag></div>
              
                <!-- <router-link :to="{path:'/EditBlogs',query:{id:blogData[index].id}}" target="_blank"  rel="opener"><el-icon size="25" :color="red" circle><Edit /></el-icon></router-link>&nbsp;&nbsp;&nbsp;
                <el-icon size="25" class="btndel" @click="deleteItem(blogData[index].id)" > <Delete /></el-icon>  -->
                <!-- <router-link :to="{path:'/MyDetail',query:{id:blogData[index].id}}"  ><el-button type="success"  icon="View" circle title="查看详细" /></router-link>&nbsp;&nbsp; -->
                <router-link :to="{path:'/EditBlogs',query:{id:blogData[index].id}}" target="_blank"  rel="opener"><el-button type="primary" icon="Edit" circle title="编辑" /></router-link>&nbsp;&nbsp;
                <!-- <el-button type="danger" icon="Delete" @click="deleteItem(blogData[index].id)" circle title="删除" /> -->
                <el-popconfirm title="确定删除?" @confirm="deleteItem(blogData[index].id)">
                  <template #reference>
                    <el-button type="danger" icon="Delete"  circle ></el-button>
                  </template>
                </el-popconfirm>
              </div>
          </div>
        </el-card>
      </el-space>
      <el-col align="middle" style="margin-top: 10px">
        <el-button @click="GetDatas()" v-if="btdisabled==false" >加载数据..</el-button>
        <el-button  v-else  disabled>-显示完毕-</el-button> 
       
      </el-col>
    </el-row>
  </template>

  <script>
  import { ref} from 'vue' 
 
 
  import axios from 'axios'
  
  import moment from "moment"

  import 'element-plus/dist/index.css'
  
  
  moment.locale('zh-cn')
  
  export default {
    data() {
      return {
        form: {
          desc: ''
        },
        blogData:[],
        total:0,
        currentDate : ref(new Date()),
        message: "",
        page:1,
        btvalue:'加载数据..',
        btdisabled:false,
        moments:moment,
        loading: true,
        suiname:sessionStorage.name,
        suiphoto:sessionStorage.photo,
      }
    },
    mounted() { 
      axios({
           method: "get",
           url: this.$globalHost+"api/v1/articles?created_by="+sessionStorage.id,//"http://127.0.0.1:2023/api/v1/articles?created_by="+sessionStorage.id,
           headers:{
              "token":sessionStorage.token
           }
          })
          .then((res) => {
            this.loading = false
           //test=res.data.data.content 
           console.log(res.data)
          this.blogData=res.data.data.lists
          this.total=res.data.data.total
           console.log("userid:"+sessionStorage.id)
           //datas.blogData=JSON.parse(JSON.stringify(res.data.data))  //eval(res.data.data).lists[7]
           if(this.blogData.length==this.total){
              this.btvalue="-显示完毕-"
              this.btdisabled=true
             }
            //  let tds=this.moments().unix(1690728099*1000)
            //  console.log("test:"+ this.moments(1691674049*1000).format("YYYY-MM-DD HH:mm:ss"))
           
      }) 
    },
    methods: {
       GetDatas(){ 
         this.page=this.page+1
          //alert(this.page)
          axios({
             method: "get",
             url:this.$globalHost+"api/v1/articles?page="+this.page+"&created_by="+sessionStorage.id, //"http://127.0.0.1:2023/api/v1/articles?page="+this.page+"&created_by="+sessionStorage.id,
             headers:{
                "token":sessionStorage.token
             }
            })
            .then((res) => {
              this.loading = false
             //test=res.data.data.content 
             console.log(res.data)
             this.blogData.push(...res.data.data.lists)
             this.total=res.data.data.total
             if(this.blogData.length==this.total){
              this.btvalue="-显示完毕-"
              this.btdisabled=true
             }
             console.log(this.total) 
             
             }) 
      },
      Delbolg(bgid){
           alert(bgid)
      },
      deleteItem(id) { 
        // if(!confirm("确定删除?"))
        // {
        //    return
        // }
           axios({
             method: "delete",
             url: this.$globalHost+"api/v1/articles/"+id,//"http://127.0.0.1:2023/api/v1/articles/"+id,
             headers:{
                "token":sessionStorage.token
             }
            })
            .then((res) => {
              console.log('删除成功', res);
               alert('删除成功')
               this.page=0
               this.blogData=[]
               this.GetDatas()
             }) 
            //  .catch(error => {
            //     // 处理错误响应
            //     console.error('删除失败', error);
            //     alert('删除失败')
            //   });
         
      }

    }
  }
   
  </script>
  
  <style>
  .time {
    font-size: 14px;
    color: #999;
  }
  
  .bottom {
    margin-top: 13px;
    line-height: 12px;
    display: flex;
    justify-content: space-between;
    align-items: center;
  }
  
  .button {
    padding: 0;
    min-height: auto;
  }
  
  .image {
    width: 100%;
    display: block;
    max-height:144px;
    object-fit: scale-down;
  }
  .blog-overview{
    color:#202124;
    font-size:14px;
    line-height: 22px;
    word-break: break-all;
    display: -webkit-box;
    overflow: hidden;
    text-overflow: ellipsis;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 3;
    margin-top:12px;
    min-height:66px;
  }
  .foot_btn{
    text-align:right;
  }
  .btndel{
    cursor: pointer;
  }
  .my_bk_head{
    width: 100%;
    height: auto;
    overflow: hidden;
    max-height: 358px;
  }
  .my_bk_head img{
    width:100%;
  }
  .my_bk_body{
    width: 100%;
    height: 42px;
    border: 1px solid #D3D6DB;
    background: #fff;
    float: left;
    margin-top: -42px;
    border-bottom-right-radius: 15px;
    border-bottom-left-radius: 15px;
  }
  .my_bk_body_photo{
    width: 86px;
    height: 86px;
    border: 1px solid #6B7878;
    border-radius: 4px;
    float: right;
    margin-top: -68px;
    margin-right: 46%;
    background: #fff;
    text-align: center;
    border-radius: 50%;
  }
  .my_bk_body_photo img{
    width: 80px;
    height: 80px;
    margin-top: 3px;
    border-radius: 50%;
  }
  .my_bk_body_name{
        /* text-align: right; */
    /* width: 160px; */
    height: 30px;
    float: right;
    margin-top: 16px;
    margin-right: -70px;
    font-size: 20px;
    color: #fff;
    text-rendering: optimizelegibility;
    text-shadow: 0 0 3px rgba(0, 0, 0, .8);
    line-height: 30px;
    font-family: Helvetica, Arial, sans-serif;
    font-weight: bold;
    letter-spacing: 0;
  }
  .my_bk_body_right{
    /* width: 70%; */
    float: right;
    /* right: 0; */
    margin-right: 28%;
    text-align: right;
    margin-top: -34px;
  }
  .my_bk_body_left{
    width: 35px;
    height: 35px; 
    text-align: center;
    margin-left: 38%;
    margin-top: 5px;
    display: block; 
  }
  .idx_mymenu_top_logo {
    width: 35px;
    height: 35px;
    display: flex;
    flex-wrap: wrap;
    justify-content: center;
    border-radius: 50%;
    background-color: #6bdd88;
    cursor: pointer;
}
.my_bk_mkdown{
  width: 22px;
  height: 22px;
  float: right;
  margin-top: 6px;
  padding-left: 10px;
}
  </style>
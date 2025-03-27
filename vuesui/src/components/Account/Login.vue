<template> 
    <div class="loginBody"><div id="particles-js"></div>
        <div class="loginDiv">
            <div class="login-content">
                <h2 class="login-title">登&nbsp;&nbsp;录</h2>
                <hr/>
                <div class="login-body">
                <el-form :model="loginForm"  :rules="rules" ref="loginForm">
                    <el-form-item label="名称"  class="redlbl" style="color:#fff;  padding-top: 20px;padding-right: 30px;" prop="username">
                        <el-input style="width:160px" type="text" v-model="loginForm.username" placeholder="请输入用户名"
                        autocomplete="off" size="default" >
                        </el-input>
                    </el-form-item>
                    <el-form-item label="密码"  class="redlbl"  label-color="#fff" prop="password" style="color:#fff;  padding-top: 20px;padding-right: 30px;" >
                        <el-input style="width:160px" type="password" v-model="loginForm.password"
                        show-password autocomplete="off" size="default"  placeholder="请输入密码"  @keyup.enter="confirm('loginForm')">
                        </el-input>
                    </el-form-item>
                    <el-form-item label="  " class="redlbl"  style="padding-top: 20px;">
                        <el-button type="primary" style="width:160px" @click="confirm('loginForm')">登 录</el-button>
                    </el-form-item>
                    <el-form-item class="redlbl"  style="padding-top: 10px;padding-right:5px;">
                        &nbsp;暂无账号?&nbsp;<a href="#">注册新账号</a>
                    </el-form-item>
                </el-form>
            </div>
         
            </div>
        </div>
           <div style=" position: absolute;   width: 100%; bottom: 0;  text-align: center;    color: #ccc;">
      
  </div>
    </div>
</template>
<script>
//import {validUserName} from  '@/utils/validate'
//import {Mkdown} from './components/Mkdown.vue'
//import Router from 'vue-router'
import particlesConfig from '../../particles-config'
import { h } from 'vue'
import { ElNotification } from 'element-plus'
import axios from 'axios'


export default{
    name: 'Login',
    data(){
         return{
            loginForm:{
                username:'',
                password:''
            },
            rules:{

            }
         }         
    },mounted(){
      const script = document.createElement('script');
      script.src =  'https://cdn.jsdelivr.net/npm/particles.js';
      script.onload = () => {
        this.initParticles();
      };
      document.body.appendChild(script);
    },
    methods:{
      initParticles() {
      // 确保 particlesJS 是全局函数
         if (window.particlesJS) {
           window.particlesJS('particles-js', particlesConfig);
         } else {
           console.error('particlesJS is not available.');
         }
      },
        confirm(formName){
            this.$refs[formName].validate((valid) => {
                if (valid) {
                    
                    console.log(this.loginForm)
                    //http://localhost:2024/auth 
                   axios({
                      method: "post",
                      url: this.$globalHost+"auth",//"http://localhost:2023/auth",//?username="+this.loginForm.username+"&password="+this.loginForm.password,
                      data: {
                        username: this.loginForm.username,
                        password:this.loginForm.password
                      }
                    }).then((res) => {
                      console.log(res);
                      if(res.data.msg=="ok"){
                        sessionStorage.token =res.data.data.token
                        sessionStorage.id =res.data.data.id
                        sessionStorage.name=res.data.data.name
                        sessionStorage.photo=res.data.data.photo
                        this.$router.push("Home")
                      }else{
                        ElNotification({
                            title: '登录失败',
                            message: h('i', { style: 'color: teal' }, '用户名或密码错误,请重新登录!'),
                            type:'error'
                          })
                      }
                      
                    })
                }
           })
        }
    }
}
</script>

<style scoped>
  .loginBody{
    width: 100vw;
  height: 100vh;
  overflow: hidden;
    /* width:100%;
    height:100%; */
    /* min-height: 100vh; */
    /* background-color: #B3C0D1; */
    background-image: url(../../assets/sysimg/MusicSC1.jpg);
    background-repeat: no-repeat;
    background-size: cover;
  }
  .loginDiv{
    /* position: absolute;
    top:50%;
    left:50%;
    margin-top:-200px;
    margin-left:-250px;
    width:450px;
    height:330px;
    background:#fff;
    border-radius: 5%; */
    position: absolute;
    top: 40%;
    left: 50%;
    transform: translate(-50%, -50%);
    text-align: center;
    color: #ffffff;
  }
  /* .login-title{
    margin: 20px 0 0 0;
    text-align: center;
  } */
  .login-content{
    background: rgb(39 39 39 / 50%);
    padding: 20px;
    border-radius: 10px;
  }
  /*.login-context{
     width:400px;
    height:250px;
    position: absolute;
    top:25px;
    left:25px;
    
  } */
   /* .login-body{
    margin-top:30px;
    margin-left: 26px;
   } */

 
</style>
<style>
.redlbl{
  margin-bottom: 0;
}
   .redlbl .el-form-item__label {
  color:#fff;
}
 
.redlbl .el-form-item__content{
  justify-content: center;
}
</style>

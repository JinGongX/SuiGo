//import Account from '../components/Account/Login.vue'
//import Mkdown from '../components/Mkdown.vue'
//import Home from '../components/Home/Index.vue'
//import Chatsui from '../components/Chatsui/Index.vue'
//import Blogs from '../components/Blogs/Index.vue'
//import Testing from '../components/Test/Testing.vue'
//
//import PushMsg from '../components/Test/PushMsg.vue'
//import ChatMsg from '../components/Test/ChatMsg.vue'
//import myBlogs from '../components/Blogs/myBlogs.vue'
//import MyDetail from '../components/BlogsDetail/MyDetail.vue'
//import EditBlogs from '../components/Blogs/EditBlogs.vue'
//
//import SoftInfo from '../components/SoftwareInfo/Index.vue'
//
//import BlogDeail from '../components/Test/Testing.vue'
//
//import MsgInfo from '../components/Test/PChatMsg.vue'
//import SeePdf from '../components/FilePdf/Index.vue'
//import OpenPdf from '../components/FilePdf/OpenPdf.vue'
//
//import ExcalidrawWrapper from '../components/Excalidrawcp/ExcalidrawWrapper.vue'
//import BookInfo from '../components/BookInfo/Index.vue'
//import BlogRead from '../components/BlogsDetail/Read.vue'
////test
//import RedBlogs from '../components/Test/RedBlogs.vue'
//import TestOutLine from '../components/Test/TestOutLine.vue'
//
//import Pushkimi from '../components/Blogs/pushKimi.vue'
//test

const routes = [
     {
         name: '',
         path: '/',
         component:()=> import('../components/Home/Index.vue'),// Home,
         meta: { requiresAuth: true }
     },
    {
        name: 'Account',
        path: '/Account',
        component:()=>import('../components/Account/Login.vue'), //Account
    },
    {
        name: 'Pushmsg',
        path: '/PushMsg',
        component:()=>import('../components/Test/PushMsg.vue'), //PushMsg,
        meta: { requiresAuth: true }
    },
    {
        name: 'ChatMsg',
        path: '/ChatMsg',
        component:()=>import('../components/Test/ChatMsg.vue'), //ChatMsg,
        meta: { requiresAuth: true }
    },
    {
        name: 'Home',
        path: '/Home',
        component:()=> import('../components/Home/Index.vue'),//Home,
        redirect:'/Blogs',
        meta: { requiresAuth: true },
        children:[{
            name: 'Mkdown',
            path: '/Mkdown',
            component: ()=> import('../components/Mkdown.vue'),//Mkdown,
            meta:{
                showNav:true
            },
            children:[{
                name: 'Chatsui',
                path: '/Chatsui',
                component:()=> import('../components/Chatsui/Index.vue'),// Chatsui,
            }]
        },{
            name: 'ChatMsg',
            path: '/ChatMsg',
            component: ()=> import('../components/Test/ChatMsg.vue'),//ChatMsg,
            children:[{
                name: 'MsgInfo',
                path: '/MsgInfo',
                component: ()=> import('../components/Test/PChatMsg.vue'),//MsgInfo,
            }]
        },{
            name: 'Chatsui',
            path: '/Chatsui',
            component: ()=> import('../components/Chatsui/Index.vue'),//Chatsui,
        },{
            name: 'Blogs',
            path: '/Blogs',
            component: ()=> import('../components/Blogs/Index.vue'),//Blogs, 
            children:[{
                name: 'Testing',
                path: '/Testing',
                component: ()=> import('../components/Test/Testing.vue'),//Testing,
                props:(route)=>{
                    return route.query.id
                }
            }]
        },{
            name: 'myBlogs',
            path: '/myBlogs',
            component: ()=> import('../components/Blogs/myBlogs.vue'),//myBlogs,
            children:[{
                name: 'MyDetail',
                path: '/MyDetail',
                component: ()=> import('../components/BlogsDetail/MyDetail.vue'),//MyDetail,
                props:(route)=>{
                    return route.query.id
                }
            },{
                name: 'EditBlogs',
                path: '/EditBlogs',
                component: ()=> import('../components/Blogs/EditBlogs.vue'),//EditBlogs,
                props:(route)=>{
                    return route.query.id
                },
                meta:{
                    showNav:true
                }
            }] 
        },{
            name: 'MyDetail',
            path: '/MyDetail',
            component: ()=> import('../components/BlogsDetail/MyDetail.vue'),//MyDetail,
        },{
            name: 'EditBlogs',
            path: '/EditBlogs',
            component: ()=> import('../components/Blogs/EditBlogs.vue'),//EditBlogs,
            meta:{
                showNav:true
            }
        },{
            name: 'SoftInfo',
            path: '/SoftInfo',
            component: ()=> import('../components/SoftwareInfo/Index.vue'),//SoftInfo,
        },{
            name: 'Testing',
            path: '/Testing',
            component: ()=> import('../components/Test/Testing.vue'),//Testing,
            props:(route)=>{
                return route.query.id
            }
        },{
            name: 'ExcalidrawWrapper',
            path: '/ExcalidrawWrapper',
            component: ()=> import('../components/Excalidrawcp/ExcalidrawWrapper.vue'),//ExcalidrawWrapper, 
            meta:{
                showNav:true
            },
        },{
            name: 'BookInfo',
            path: '/BookInfo',
            component:()=> import('../components/BookInfo/Index.vue'),// BookInfo
        }]
    },{
        name: 'BlogDeail',
        path: '/BlogDeail',
        component: ()=> import('../components/Test/Testing.vue'),//BlogDeail,
        props:(route)=>{
            return route.query.id
        }
    },{
        name: 'SeePdf',
        path: '/SeePdf',
        component: ()=> import('../components/FilePdf/Index.vue'),//SeePdf,
        children:[{
            name: 'OpenPdf',
            path: '/OpenPdf',
            component: ()=> import('../components/FilePdf/OpenPdf.vue'),//OpenPdf,
            meta:{
                showNavOpen:true
            }
            // props:(route)=>{
            //     return route.query.namepdf
            // }
        }] 
    },{
        name: 'RedBlogs',
        path: '/RedBlogs',
        component: ()=> import('../components/Test/RedBlogs.vue'),//RedBlogs
    },{
        name: 'BlogRead',
        path: '/BlogRead',
        component: ()=> import('../components/BlogsDetail/Read.vue'),//BlogRead
    },{
        name: 'testd',
        path: '/testd',
        component: ()=> import('../components/Test/Blogtest.vue'),//BlogRead
    }
    
    
];

export default routes
import {createRouter, createWebHistory} from 'vue-router'
import routes from './routes'    // 导入 router 目录下的 router.js

const router = createRouter({
	history: createWebHistory(),
	routes
})
function isTokenExpired(token) {
    if (!token) {
        return true; // 没有令牌则视为已过期
    }

    const tokenData = JSON.parse(atob(token.split('.')[1])); // 解码令牌的payload部分
    const expirationTime = tokenData.exp * 1000; // 转换为毫秒
	console.log(Date.now()+"==="+expirationTime)
    return Date.now() >= expirationTime;
}

 router.beforeEach((to,from,next)=>{
 	//if(to.path=='/Login') return next()

 	//const tokenStr =sessionStorage.token //window.sessionStorage.getItem('token')
 	//if(!tokenStr) return next('/Login')
 	//next()
	
	const isLoggedIn=sessionStorage.getItem('token')
	console.log(isLoggedIn)
	if((isLoggedIn==null&& to.meta.requiresAuth)){
		console.log("no ")
		next('/Account')
	}else{ 
		if(isTokenExpired(isLoggedIn)){
			sessionStorage.removeItem('token')
			sessionStorage.removeItem('id')
		}
		console.log("yes ")
		next() 
		//console.log("yes ")
		//	next()
	}
 })
export default router;
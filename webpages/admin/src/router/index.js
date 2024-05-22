import Vue from 'vue'
import VueRouter from 'vue-router'
import Login from '../views/Login.vue'
import Admin from '../views/Admin.vue'

// 页面路由
import Dashboard from '../components/admin/Dashboard.vue'
import AddArticle from '../components/admin/article/AddArticle.vue'
import ArticleList from '../components/admin/article/ArticleList.vue'
import CategoryList from '../components/admin/category/CategoryList.vue'
import UserList from '../components/admin/user/UserList.vue'

Vue.use(VueRouter)
const routes = [
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/admin',
    name: 'admin',
    component: Admin,
    children:[{path:'dashboard',component:Dashboard},
              {path:'addarticle',component:AddArticle},
              {path:'articlelist',component:ArticleList},
              {path:'categorylist',component:CategoryList},
              {path:'userlist',component:UserList}
    ]
  }
]

const router = new VueRouter({
  routes
})

// 没有token时进入admin强制跳转到login
router.beforeEach((to,from,next)=>{
  const token=window.sessionStorage.getItem('token')
  if(to.path=='/login'){
    return next()
  }
  if(!token&&to.path=='/admin'){
    next('/login')
  }
  else{
    next()
  }
})

export default router

import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/user_Home.vue'
import Content from '../views/user_Content.vue'
import Publish from '../views/user_Publish.vue'
import Login from '../views/user_Login.vue'
import SignUp from '../views/user_SignUp.vue'
import Registration from '../views/user_Registration.vue'
import DisplayDataPage from '../views/user_DisplayDataPage.vue'
import Declaration from "../views/user_Declaration.vue"
import UserInformation from '../views/user_Information.vue'
import Fu from '../views/Fu.vue'
import Fu1 from '../views/Fu1.vue'
import admlogin from '../views/adm_login.vue'
import admindex from '../views/adm_index.vue'

const originalPush = VueRouter.prototype.push;
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err);
}
Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  //添加用户个人信息页面
  {
    path: '/userInformation',
    name: 'UserInformation',
    component: UserInformation,
  },
  {
    path: '/post/:id',
    name: 'Content',
    component: Content
  },
  {
    path: '/publish',
    name: 'Publish',
    component: Publish,
    meta: { requireAuth: true }
  },
  {
    path: '/login',
    name: "Login",
    component: Login
  },
  {
    path: '/signup',
    name: "SignUp",
    component: SignUp
  },
  {
    path: '/registration',
    name: "Registration",
    component: Registration
  },
  {
    path: '/display',
    name: "DisplayDataPage",
    component: DisplayDataPage,
    props: true
  },
  {
    path: '/declaration',
    name: "Declaration",
    component: Declaration
  },
  {
    path: '/fu',
    name: "Fu",
    component: Fu,
  },
  {
    path: '/fu1',
    name: "Fu1",
    component: Fu1,
  },
  {
    path: '/admlogin',
    name: "admlogin",
    component: admlogin,
  },
  {
    path: '/admindex',
    name: "admindex",
    component: admindex,
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router

import Vue from 'vue'
import App from './App.vue'
import router from './router'
import axios from 'axios'
import './plugin/antui'
import './assets/css/style.css'

axios.defaults.baseURL = 'http://192.168.1.17:3131/interfaces/version1'
Vue.prototype.$http = axios

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
